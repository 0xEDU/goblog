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
	// request := &pb.CreateArticleRequest{
	// 	Author:  "edu",
	// 	Markdown: []byte("## Hello"),
	// }
	// response, err := client.CreateArticle(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to get article list: %v", err)
	}
	// log.Printf("Article %s", response.String())

	for _, article := range response.Articles {
		log.Printf("Article: %s", article.String())
	}
}
