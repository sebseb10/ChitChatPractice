package server

import (
	proto "ChitChatPractice/grpc"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedChitChatServiceServer
	lamport_time int64
	clients      map[int32]string
}

var clientCount int32 = 0

func main() {

}

func (s *Server) join(ctx context.Context, req *proto.JoinRequest) (*proto.JoinResponse, error) {
	clientCount++
	s.clients[clientCount] = req.ClientName

	/*s.broadcast(proto.ChatEvent{
		Type:        2,
		ClientId:    clientCount,
		Name:        s.clients[clientCount],
		LamportTime: 0,
		Content:     "",
	}*/
	return &proto.JoinResponse{ClientId: clientCount}, nil
}

func (s *Server) startServer() {
	grpcServer := grpc.NewServer()
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	proto.RegisterChitChatServiceServer(grpcServer, s)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
