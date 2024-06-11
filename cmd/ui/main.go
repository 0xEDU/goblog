package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/0xEDU/goblog/pkg/proto"
)

const GrpcPort = "9090"

func main() {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:"+GrpcPort, creds)
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewArticleServiceClient(conn)
	response, err := client.GetArticleList(context.Background(), &pb.ArticleListRequest{})
	if err != nil {
		log.Fatalf("failed to get article list: %v", err)
	}

	for _, article := range response.Articles {
		log.Printf("Article: %s", article.String())
	}
}
