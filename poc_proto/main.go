package main

import (
	"context"
	"log"
	"net"

	pb "github.com/go_bank/protos"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedShopServer
}

func (s *Server) GetMenu(menuRequest *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		{
			Id:    "1",
			Name:  "Momos",
			Price: 300.0,
		},
		{
			Id:    "2",
			Name:  "Coffee",
			Price: 250,
		},
		{
			Id:    "3",
			Name:  "Tea",
			Price: 50,
		},
	}

	for i := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}
	return nil
}
func (s *Server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {

	return &pb.Receipt{
		Id:         "54",
		TotalPrice: 1250.23,
	}, nil

}
func (s *Server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {

	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "Cooking",
	}, nil

}

func main() {
	// setup a listener on 9001
	listen, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterShopServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to start grpc server: %v", err)
	}
	log.Print("Started grpc server")
}
