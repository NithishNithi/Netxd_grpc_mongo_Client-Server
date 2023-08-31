package netxdcontrollers

import (
	"context"
	"fmt"
	"log"

	pro "github.com/NithishNithi/netxd_customerproto/netxd_grpc_mongo_proto/Customer_Protobuff"
	"github.com/NithishNithi/netxd_grpc_mongo_dal/netxd_grpc_mongo_dal/models"
)



func (s *RPCServer) MakeTransaction(ctx context.Context, req *pro.UpdateBalance) (*pro.UpdateResponse, error) {
	fmt.Println("1")
	customer_record := &models.MakeTransactions{
		From_CustomerId: int(req.From_CustomerId),
		To_CustomerId: int(req.To_CustomerId),
		Amount: float64(req.Amount),
	}

	// Call the CreateCustomer method of the CustomerService
	result, err := CustomerService.MakeTransactions(customer_record)
	if err != nil {
		log.Fatal(err)
	}

	// Create the responseCustomer only when there's no error
	responseCustomer := &pro.UpdateResponse{
		From_CustomerId: int32(result.From_CustomerId),
		To_CustomerId: int32(result.To_CustomerId),
		Transaction_Time: result.Transactiontime,
	}

	return responseCustomer, nil
}
