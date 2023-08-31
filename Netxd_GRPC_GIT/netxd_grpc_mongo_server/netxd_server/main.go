package main

import (
	netxdconfig "Netxd_GRPC_GIT/netxd_grpc_mongo_server/netxd_config"
	netxdconstants "Netxd_GRPC_GIT/netxd_grpc_mongo_server/netxd_constants"
	netxdcontrollers "Netxd_GRPC_GIT/netxd_grpc_mongo_server/netxd_controllers"

	pro "github.com/NithishNithi/netxd_customerproto/netxd_grpc_mongo_proto/Customer_Protobuff"
	"google.golang.org/grpc"

	"context"
	"fmt"
	"net"

	"github.com/NithishNithi/netxd_grpc_mongo_dal/netxd_grpc_mongo_dal/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDatabase(client *mongo.Client) {
	profileCollection1 := netxdconfig.GetCollection(client, "Netxd_Bank", "Customers")
	profileCollection2 := netxdconfig.GetCollection(client, "Netxd_Bank", "Transactions")
	netxdcontrollers.CustomerService = services.InitCustomerService(profileCollection1,profileCollection2, context.Background())
}

func main() {
	mongoclient, err := netxdconfig.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", netxdconstants.Port)

	if err != nil {
		fmt.Println("Failed TO Listen", err)
	}
	s := grpc.NewServer()

	pro.RegisterCustomerServiceServer(s, &netxdcontrollers.RPCServer{})

	fmt.Println("Server listening on", netxdconstants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
