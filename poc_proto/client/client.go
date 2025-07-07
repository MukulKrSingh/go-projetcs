package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/go_bank/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9001", grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		log.Fatalf("Failed to get client : %v", err)
	}
	defer conn.Close()

	c := pb.NewShopClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})

	if err != nil {
		log.Fatalf("Failed to get steam:%v", err)
	}
	done := make(chan bool)
	var items []*pb.Item

	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil && err != io.EOF {
				log.Fatalf("Problem : %v", err)

			}
			items = resp.Items
			log.Printf("Response : %v", items)
		}
	}()
	<-done

	receipt, err := c.PlaceOrder(ctx, &pb.Order{Item: items})
	log.Printf("Receipte: %v", receipt)
	if err != nil {
		log.Printf("failed to place order : %v", err)
	}

	status, err := c.GetOrderStatus(ctx, receipt)
	log.Printf("Status : %v", status)
	if err != nil {
		log.Printf("failed to status order : %v", err)
	}

}
