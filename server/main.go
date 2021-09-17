package main
 
import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	credit "github.com/danielfoehrkn/resource-reservations-grpc/pkg/proto/gen/resource-reservations"
)
 
type server struct {
	credit.UnimplementedResourceReservationsServer
}
 
func main() {
	log.Println("Server running ...")
 
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
 
	srv := grpc.NewServer()
	credit.RegisterResourceReservationsServer(srv, &server{})
 
	log.Fatalln(srv.Serve(lis))
}
 
func (s *server) UpdateResourceReservations(ctx context.Context, request *credit.UpdateResourceReservationsRequest) (*credit.UpdateResourceReservationsResponse, error) {
	log.Println(fmt.Sprintf("Request kube-reserved: %v", request.KubeReserved))
	log.Println(fmt.Sprintf("Request system-reserved: %v", request.SystemReserved))

	// TODO: possibly add error message if validation failed
	return &credit.UpdateResourceReservationsResponse{}, nil
}
