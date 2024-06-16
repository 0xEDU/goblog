package pb_client

import (
	"log"

	pb "github.com/0xEDU/goblog/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetPbClient() {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("localhost:9090", creds)
	if err != nil {
		log.Println("could not connect to server: %v", err)
		return
	}
}
