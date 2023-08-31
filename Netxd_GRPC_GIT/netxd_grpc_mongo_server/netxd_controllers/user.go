package netxdcontrollers


import(
	pro "github.com/NithishNithi/netxd_customerproto/netxd_grpc_mongo_proto/Customer_Protobuff"
	"github.com/NithishNithi/netxd_grpc_mongo_dal/netxd_grpc_mongo_dal/interfaces"
)
type RPCServer struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomers
)