package main

import (
	"context"
	"log"
	"net"

	pb "github.com/0xEDU/goblog/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const Port = "9090"

type articleServiceServer struct {
	pb.UnimplementedArticleServiceServer
}

func NewArticeServiceServer() *articleServiceServer {
	return &articleServiceServer{}
}

func (svc *articleServiceServer) GetArticleList(ctx context.Context, request *pb.ArticleListRequest) (*pb.ArticleListResponse, error) {
	articles := []*pb.Article{
		{Id: "1", Author: "edu", Markdown: []byte("## Hello")},
	}
	return &pb.ArticleListResponse{Articles: articles}, nil
}

func main() {
	l, err := net.Listen("tcp", ":" + Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArticleServiceServer(grpcServer, NewArticeServiceServer())
	reflection.Register(grpcServer)

	log.Printf("Serving gRPC server on port %s", Port)
	grpcServer.Serve(l)
}
