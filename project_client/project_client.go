package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/chimanjain/grpc-go-hack/project"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(address, opts...)
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()
	p := pb.NewProjectManagementClient(conn)

	var id uint32 = 1

	printProject(p, &pb.ProjectID{Id: id})

	printProjects(p, &pb.ProjectID{Id: id})

	printfetchProjects(p)

	printStreamProjects(p)
}

// printProject gets the project for the given projectId.
func printProject(client pb.ProjectManagementClient, id *pb.ProjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	project, err := client.GetProject(ctx, id)
	if err != nil {
		log.Printf("%v.GetProject(_) = _, %v: ", client, err)
	}
	log.Println(project)
}

// printProjects lists all the projects within the given projectId.
func printProjects(client pb.ProjectManagementClient, id *pb.ProjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.GetProjects(ctx, id)
	if err != nil {
		log.Printf("%v.GetProjects(_) = _, %v", client, err)
	}
	for {
		project, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("%v.GetProjects(_) = _, %v", client, err)
		}
		log.Println(project)
	}
}

// printfetchProjects sends a sequence of ProjectID to server and expects to get a Projects from server.
func printfetchProjects(client pb.ProjectManagementClient) {
	// Create a random number of random points
	IDs := []uint32{1, 2, 3, 4}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.FetchProjects(ctx)
	if err != nil {
		log.Printf("client.RecordRoute failed: %v", err)
	}
	for _, ID := range IDs {
		if err := stream.Send(&pb.ProjectID{Id: ID}); err != nil {
			log.Printf("client.FetchProjects: stream.Send(%v) failed: %v", ID, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("client.FetchProjects failed: %v", err)
	}
	log.Printf("fetch Projects summary: %v", reply)
}

// printStreamProjects receives a sequence of projects, while sending projectID.
func printStreamProjects(client pb.ProjectManagementClient) {
	projectIDs := []*pb.ProjectID{
		{Id: 1},
		{Id: 2},
		{Id: 3},
		{Id: 4},
		{Id: 5},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.StreamProjects(ctx)
	if err != nil {
		log.Printf("client.RouteChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Printf("client.RouteChat failed: %v", err)
			}
			log.Println(in)
		}
	}()
	for _, projectID := range projectIDs {
		if err := stream.Send(projectID); err != nil {
			log.Printf("client.RouteChat: stream.Send(%v) failed: %v", projectID, err)
		}
	}
	_ = stream.CloseSend()
	<-waitc
}
