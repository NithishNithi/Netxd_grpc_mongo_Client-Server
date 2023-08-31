package main

import (
	"context"
	"fmt"
	"log"

	pro "github.com/NithishNithi/netxd_customerproto/netxd_grpc_mongo_proto/Customer_Protobuff"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5050", grpc.WithInsecure())
	if err != nil {
		log.Fatal("Failed to Connect", err)
	}
	defer conn.Close()
	client := pro.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pro.CustomerDetails{CustomerId: 102, FirstName: "Vikram", LastName: "s", BankId: "1664Bank2", Balance: 5000.00})
	if err != nil {
		log.Fatalf("Failed to call CreateCustomer: %v", err)
	}

	fmt.Println("Response:", response.CustomerId, response.CreatedAt)
}
