package main
 
import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	resources "github.com/danielfoehrKn/resource-reservations-grpc/pkg/proto/gen/resource-reservations"
)
 
func main() {
	log.Println("Client running ...")
 
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
 
	client := resources.NewResourceReservationsClient(conn)

	kubeReserved := map[string]string{
		"cpu": "1",
		"memory": "2Gi",
		"pid": "10k",
	}

	systemReserved := map[string]string{
		"memory": "200Mi",
		"pid": "1k",
	}

	request := &resources.UpdateResourceReservationsRequest{KubeReserved: kubeReserved, SystemReserved: systemReserved}
 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
 
	response, err := client.UpdateResourceReservations(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	// will be empty
	log.Println("Response:", response)
}
