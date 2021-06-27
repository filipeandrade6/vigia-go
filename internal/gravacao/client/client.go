package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/pb"
	"google.golang.org/grpc"
)

func Main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial("localhost:10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	// TODO ver nomes -> NewGravacaoConnClient?
	client := pb.NewGravacaoConnClient(conn)

	teste(client, &pb.GravacaoConfig{Id: "mintPC"})

}

func teste(client pb.GravacaoConnClient, cfg *pb.GravacaoConfig) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	dbConfig, err := client.GetDatabaseConfig(ctx, cfg)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(dbConfig.GetDbname(), dbConfig.GetHost(), dbConfig.GetPassword(), dbConfig.GetPoolmaxconns(), dbConfig.GetPort(), dbConfig.GetUser())
}
