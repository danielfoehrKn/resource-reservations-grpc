package main
 
import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	proto "github.com/danielfoehrkn/resource-reservations-grpc/pkg/proto/gen/resource-reservations"
)
 
type server struct {
	proto.UnimplementedResourceReservationsServer
}
 
func main() {
	log.Println("Server running ...")
 
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
 
	srv := grpc.NewServer()
	proto.RegisterResourceReservationsServer(srv, &server{})
 
	log.Fatalln(srv.Serve(lis))
}
 
func (s *server) UpdateResourceReservations(ctx context.Context, request *proto.UpdateResourceReservationsRequest) (*proto.UpdateResourceReservationsResponse, error) {
	log.Println(fmt.Sprintf("Request kube-reserved: %v", request.KubeReserved))
	log.Println(fmt.Sprintf("Request system-reserved: %v", request.SystemReserved))

	// TODO: possibly add error message if validation failed
	return &proto.UpdateResourceReservationsResponse{}, nil
}

func (s *server) GetResourceReservations(ctx context.Context, request *proto.GetResourceReservationsRequest) (*proto.GetResourceReservationsResponse, error) {
	systemReserved := map[string]string{
		"memory": "2Gi",
		"cpu": "1",
	}

	kubeReserved := map[string]string{
		"memory": "3Gi",
		"cpu": "2",
	}

	return &proto.GetResourceReservationsResponse{
		SystemReserved: systemReserved,
		KubeReserved:   kubeReserved,
	}, nil
}