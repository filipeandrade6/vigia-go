package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"
	"github.com/filipeandrade6/vigia-go/internal/core/camera"
	"github.com/filipeandrade6/vigia-go/internal/core/processo"
	"github.com/filipeandrade6/vigia-go/internal/core/registro"
	"github.com/filipeandrade6/vigia-go/internal/core/veiculo"
	"github.com/filipeandrade6/vigia-go/internal/gravacao/service/processador"
	"github.com/filipeandrade6/vigia-go/internal/sys/database"
	"github.com/filipeandrade6/vigia-go/internal/sys/operrors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OperationError struct {
	ProcessoID string
	RegistroID string
}

type GravacaoService struct {
	servidorGravacaoID string

	pb.UnimplementedGravacaoServer

	log *zap.SugaredLogger

	gerencia *GerenciaClient

	cameraCore   camera.Core
	processoCore processo.Core
	registroCore registro.Core
	veiculoCore  veiculo.Core

	processador *processador.Processador
	errChan     chan operrors.OpError
	matchChan   chan string

	errBuffChan   chan operrors.OpError
	matchBuffChan chan string
}

func NewGravacaoService(log *zap.SugaredLogger) *GravacaoService {
	return &GravacaoService{
		log:           log,
		errChan:       make(chan operrors.OpError),
		matchChan:     make(chan string),
		errBuffChan:   make(chan operrors.OpError, 1000), // TODO otimizar aqui
		matchBuffChan: make(chan string, 100),            // TODO otimizar aqui
	}
}

func (g *GravacaoService) Registrar(ctx context.Context, req *pb.RegistrarReq) (*pb.RegistrarRes, error) {
	if g.gerencia != nil {
		e := "already registered gerencia service"
		g.log.Errorw("Could not register gerencia", "error", e)
		return &pb.RegistrarRes{}, status.Error(codes.AlreadyExists, e)
	}

	err := os.MkdirAll(req.GetArmazenamento(), os.ModePerm)
	if err != nil {
		e := fmt.Sprintf("could not create directory: %s", err)
		g.log.Errorw("Could not register gerencia", "error", e)
		return &pb.RegistrarRes{}, status.Error(codes.InvalidArgument, e)
	}

	db, err := database.Connect(database.Config{
		User:         req.GetDbUser(),
		Password:     req.GetDbPassword(),
		Host:         req.GetDbHost(),
		Name:         req.GetDbName(),
		MaxIDLEConns: int(req.GetDbMaxIdleConns()),
		MaxOpenConns: int(req.GetDbMaxOpenConns()),
		DisableTLS:   req.GetDbDisableTls(),
	})
	if err != nil {
		e := fmt.Sprintf("could not connect to database: %s", err)
		g.log.Errorw("Could not register gerencia", "error", e)
		return &pb.RegistrarRes{}, status.Error(codes.Internal, e)
	}

	gerenciaClient, err := NewClientGerencia(req.GetEnderecoIp(), int(req.GetPorta()))
	if err != nil {
		e := fmt.Sprintf("could not connect to gerencia gRPC server: %s", err)
		g.log.Errorw("Could not register gerencia", "error", e)
		return &pb.RegistrarRes{}, status.Error(codes.Internal, e)
	}
	// if err := gerenciaClient.Check(req.ServidorGravacaoId); err != nil {
	// 	e := fmt.Sprintf("could not connect to gerencia gRPC server: %s", err)
	// 	g.log.Errorw("Could not register gerencia", "error", e)
	// 	return &pb.RegistrarRes{}, status.Error(codes.Internal, e)
	// }
	// ! habilitar acima

	g.gerencia = gerenciaClient

	g.cameraCore = camera.NewCore(g.log, db)
	g.processoCore = processo.NewCore(g.log, db)
	g.registroCore = registro.NewCore(g.log, db)
	g.veiculoCore = veiculo.NewCore(g.log, db)

	g.processador = processador.New(req.GetArmazenamento(), int(req.GetHorasRetencao()), g.registroCore, g.errChan, g.matchChan)

	g.UpdateVeiculos(ctx, &pb.UpdateVeiculosReq{})

	go g.start()
	go g.processador.Start()

	g.servidorGravacaoID = req.GetServidorGravacaoId()

	req.DbPassword = ""
	g.log.Infow("Gerencia registered", "params", req)

	return &pb.RegistrarRes{}, nil
}

func (g *GravacaoService) start() {
	t := time.NewTicker(10 * time.Second)

	for {
		select {
		case e := <-g.errChan:
			e.ServidorGravacaoID = g.servidorGravacaoID
			g.log.Errorw("Error received", "error", e.Error())
			if err := g.gerencia.ErrorReport(e); err != nil {
				e2 := fmt.Sprintf("could not call gerencia gRPC server: %s", err)
				g.log.Errorw("Could not call ErrorReport", "error", e2) // ! Em loop com Errorreceived acima
				g.errBuffChan <- e
			}

		case m := <-g.matchChan:
			g.log.Infow("Vehicle found", "registro_id", m)
			if err := g.gerencia.Match(m); err != nil {
				e := fmt.Sprintf("could not call gerencia gRPC server: %s", err)
				g.log.Errorw("Could not call Match", "error", e)
				g.errBuffChan <- operrors.OpError{ServidorGravacaoID: g.servidorGravacaoID, RegistroID: m, Err: err}
				g.matchBuffChan <- m
			}

		case <-t.C:
			if g.gerencia == nil {
				break
			}

			go g.errBuffFlush()
			go g.matchBuffFlush()

		}
	}
}

// TODO testar batch flush de dados

func (g *GravacaoService) errBuffFlush() {
	for {
		select {
		case e := <-g.errBuffChan:
			if err := g.gerencia.ErrorReport(e); err != nil {
				g.errBuffChan <- e
				return
			}

		default:
			return
		}
	}
}

func (g *GravacaoService) matchBuffFlush() {
	for {
		select {
		case m := <-g.matchBuffChan:
			if err := g.gerencia.Match(m); err != nil {
				e := fmt.Errorf("could not call gerencia gRPC server: %w", err)
				g.errBuffChan <- operrors.OpError{ServidorGravacaoID: g.servidorGravacaoID, RegistroID: m, Err: e}
				g.matchBuffChan <- m
				return
			}

		default:
			return
		}
	}
}

func (g *GravacaoService) RemoverRegistro(ctx context.Context, req *pb.RemoverRegistroReq) (*pb.RemoverRegistroRes, error) {
	if g.gerencia == nil {
		e := "there is not a gerencia service registered"
		g.log.Errorw("Could not remove registered gerencia", "error", e)
		return &pb.RemoverRegistroRes{}, status.Error(codes.NotFound, e)
	}
	g.gerencia = nil

	if err := g.processador.Stop(); err != nil {
		e := fmt.Sprintf("could not stop processador: %s", err)
		g.log.Errorw("Could not remove registered gerencia", "error", e)
		return &pb.RemoverRegistroRes{}, status.Error(codes.Internal, e)
	}

	// TODO interromper os cores, processador, e db?
	// TODO colocar mutex nessas paradas antes de verificar ID de servidor, se tem gerencia registrado, etc

	g.servidorGravacaoID = ""

	g.log.Infow("Gerencia service unregistered")
	return &pb.RemoverRegistroRes{}, nil
}

func (g *GravacaoService) StartProcessos(ctx context.Context, req *pb.StartProcessosReq) (*pb.StartProcessosRes, error) {
	prcs := req.GetProcessos()

	var processos []processador.Processo
	for _, prc := range prcs {
		p, err := g.processoCore.QueryByID(ctx, prc)
		if err != nil {
			e := fmt.Sprintf("query processos database: %s", err)
			g.log.Errorw("Could not start processos", "error", e)
			return &pb.StartProcessosRes{}, status.Error(codes.Internal, e)
		}

		c, err := g.cameraCore.QueryByID(ctx, p.CameraID)
		if err != nil {
			e := fmt.Sprintf("query cameras database: %s", err)
			g.log.Errorw("Could not start processos", "error", e)
			return &pb.StartProcessosRes{}, status.Error(codes.Internal, e)
		}

		processos = append(processos, processador.Processo{
			ProcessoID:  prc,
			EnderecoIP:  c.EnderecoIP,
			Porta:       c.Porta,
			Canal:       c.Canal,
			Usuario:     c.Usuario,
			Senha:       c.Senha,
			Processador: p.Processador,
		})
	}

	g.processador.StartProcessos(processos)

	g.log.Infow("Started processos", "processos", prcs)
	return &pb.StartProcessosRes{}, nil
}

func (g *GravacaoService) StopProcessos(ctx context.Context, req *pb.StopProcessosReq) (*pb.StopProcessosRes, error) {
	prcs := req.GetProcessos()

	if nonStoppedPrc := g.processador.StopProcessos(prcs); nonStoppedPrc != nil {
		e := fmt.Sprintf("could not stop processos %s from %s", nonStoppedPrc, prcs)
		g.log.Errorw("Could not stop all processos", "error", e, "processos", nonStoppedPrc)
		return &pb.StopProcessosRes{NonStoppedProcessos: nonStoppedPrc}, status.Error(codes.Internal, e)
	}

	g.log.Infow("Stopped processos", "processos", prcs)
	return &pb.StopProcessosRes{}, nil
}

func (g *GravacaoService) ListProcessos(ctx context.Context, req *pb.ListProcessosReq) (*pb.ListProcessosRes, error) {
	processos, retrying := g.processador.ListProcessos()

	g.log.Infow("Listed processos", "running", processos, "retrying", retrying)
	return &pb.ListProcessosRes{ProcessosEmExecucao: processos, ProcessosEmTentativa: retrying}, nil
}

func (g *GravacaoService) UpdateVeiculos(ctx context.Context, req *pb.UpdateVeiculosReq) (*pb.UpdateVeiculosRes, error) {
	veiculos, err := g.veiculoCore.QueryAll(ctx)
	if err != nil {
		e := fmt.Sprintf("query veiculos database: %s", err)
		g.log.Errorw("Could not update veiculos", "error", e)
		return &pb.UpdateVeiculosRes{}, status.Error(codes.Internal, e)
	}

	var matchlist []string
	for _, v := range veiculos {
		matchlist = append(matchlist, v.Placa)
	}

	g.processador.UpdateMatchlist(matchlist)

	g.log.Infow("Updated veiculos", "veiculos", matchlist)
	return &pb.UpdateVeiculosRes{}, nil
}

func (g *GravacaoService) UpdateArmazenamento(ctx context.Context, req *pb.UpdateArmazenamentoReq) (*pb.UpdateArmazenamentoRes, error) {
	armazenamento := req.GetArmazenamento()
	horasRetencao := int(req.GetHorasRetencao())

	err := g.processador.UpdateArmazenamento(armazenamento, horasRetencao)
	if err != nil {
		g.log.Errorw("Could not update armazenamento", "error", err)
		return &pb.UpdateArmazenamentoRes{}, status.Error(codes.Internal, err.Error())
	}

	g.log.Infow("Updated armazenamento", "armazenamento", armazenamento, "horas_retencao", horasRetencao)
	return &pb.UpdateArmazenamentoRes{}, nil
}
