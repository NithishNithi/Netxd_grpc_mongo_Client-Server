package main

import (
	"context"
	"fmt"
	"log"

	pro "github.com/NithishNithi/netxd_customerproto/netxd_grpc_mongo_proto/Customer_Protobuff"
	"google.golang.org/grpc"
)


func main(){
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to Connect", err)
	}
	defer conn.Close()
	client := pro.NewCustomerServiceClient(conn)
	fmt.Print("1")

	response, err := client.MakeTransaction(context.Background(),&pro.UpdateBalance{From_CustomerId: 101,To_CustomerId: 102,Amount: 5000.00})
	if err != nil {
		log.Fatalf("Failed to call MakeTransaction: %v", err)
	}

	fmt.Println("Response:", response.From_CustomerId, response.To_CustomerId,response.Transaction_Time)
}