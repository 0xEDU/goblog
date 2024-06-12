package main

import (
	"fmt"
	"io"
	"net"
	"strings"

	"os"

	"cloud.google.com/go/storage"
	pb "github.com/0xEDU/goblog/pkg/proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/appengine/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	Port       = "9090"
	ProjectID  = "goblog-426103"
	BucketName = "article-storage-2207"
)

type articleServiceServer struct {
	pb.UnimplementedArticleServiceServer
}

func NewArticeServiceServer() *articleServiceServer {
	return &articleServiceServer{}
}

func NewArticle(id string, author string, markdown []byte) *pb.Article {
	return &pb.Article{
		Id:       id,
		Author:   author,
		Markdown: markdown,
	}
}

func getBucket(ctx context.Context) (*storage.BucketHandle, error) {
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Errorf(ctx, "failed to create client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(BucketName)
	return bucket, nil
}

func (svc *articleServiceServer) GetArticleList(ctx context.Context, request *pb.ArticleListRequest) (*pb.ArticleListResponse, error) {
	bucket, err := getBucket(ctx)
	if err != nil {
		return nil, err
	}
	articlesObjs := bucket.Objects(ctx, nil)
	articlesList := []*pb.Article{}
	for {
		articleAttrs, err := articlesObjs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorf(ctx, "failed to get article: %v", err)
		}

		rc, err := bucket.Object(articleAttrs.Name).NewReader(ctx)
		if err != nil {
			// log.Errorf(ctx, "failed to create reader for %s: %v", articleAttrs.Name, err)
			fmt.Fprintln(os.Stdout, []any{"failed to create reader for %s: %v", articleAttrs.Name, err}...)
		}

		data, err := io.ReadAll(rc)
		if err != nil {
			log.Errorf(ctx, "failed to read data for %s: %v", articleAttrs.Name, err)
		}
		articlesList = append(articlesList, NewArticle(articleAttrs.Name, "edu", data))
	}
	return &pb.ArticleListResponse{Articles: articlesList}, nil
}

func (svc *articleServiceServer) CreateArticle(ctx context.Context, request *pb.CreateArticleRequest) (*pb.Article, error) {
	bucket, err := getBucket(ctx)
	if err != nil {
		return nil, err
	}
	id := uuid.New().String()
	w := bucket.Object("articles/" + id).NewWriter(ctx)
	defer w.Close()

	if _, err := io.Copy(w, strings.NewReader(string(request.Markdown))); err != nil {
		return nil, err
	}
	article := &pb.Article{
		Id:       id,
		Author:   request.Author,
		Markdown: request.Markdown,
	}
	return article, nil
}

func main() {
	ctx := context.Background()
	l, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Errorf(ctx, "failed to listen: %v", err)
	}
	defer l.Close()

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterArticleServiceServer(grpcServer, NewArticeServiceServer())
	reflection.Register(grpcServer)

	grpcServer.Serve(l)
}
