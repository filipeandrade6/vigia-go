// Server de serviço de gravação
package server

import (
	"context"
	"log"
	"net"
	"reflect"
	"testing"

	pb "github.com/filipeandrade6/vigia-go/internal/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	pb.RegisterGravacaoServer(s, &GravacaoServer{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestConfigurarProcesso(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := pb.NewGravacaoClient(conn)

	// TODO adicionar camera
	// TODO adicionar processador
	// TODO configurar processo
	// TODO info do processo
	// TODO deletar processo
	// TODO combinar com o TestCases abaixo?

	resp, err := client.ConfigurarProcesso(ctx, &pb.ConfigurarProcessoReq{
		Acao:               pb.ConfigurarProcessoReq_CONFIGURAR,
		CameraId:           0,
		ProcessadorCaminho: "",
	})
	if err != nil {
		t.Fatalf("ConfigurarProcesso failed: %v", err)
	}

	log.Printf("Response: %+v", resp)
	// Test for output here.
}

func TestGravacaoServer_InfoProcessos(t *testing.T) {
	type fields struct {
		UnimplementedGravacaoServer pb.UnimplementedGravacaoServer
	}
	type args struct {
		ctx context.Context
		req *pb.InfoProcessosReq
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.InfoProcessosResp
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &GravacaoServer{
				UnimplementedGravacaoServer: tt.fields.UnimplementedGravacaoServer,
			}
			got, err := s.InfoProcessos(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GravacaoServer.InfoProcessos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GravacaoServer.InfoProcessos() = %v, want %v", got, tt.want)
			}
		})
	}
}
