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
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	p := pb.NewProjectManagementClient(conn)

	var id uint32 = 1

	printProject(p, &pb.ProjectId{Id: id})

	printProjects(p, &pb.ProjectId{Id: id})
}

// printProject gets the project for the given projectId.
func printProject(client pb.ProjectManagementClient, id *pb.ProjectId) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	project, err := client.GetProject(ctx, id)
	if err != nil {
		log.Fatalf("%v.GetProject(_) = _, %v: ", client, err)
	}
	log.Println(project)
}

// printProjects lists all the projects within the given projectId.
func printProjects(client pb.ProjectManagementClient, id *pb.ProjectId) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client.GetProjects(ctx, id)
	stream, err := client.GetProjects(ctx, id)
	if err != nil {
		log.Fatalf("%v.GetProjects(_) = _, %v", client, err)
	}
	for {
		project, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetProjects(_) = _, %v", client, err)
		}
		log.Printf(`Project Details:
		id: %d
		name: %s
		deployId: %s
		deployAt: %s
		noticeTotalCount: %d
		rejectionCount: %d
		fileCount: %d
		deployCount: %d
		groupResolvedCount: %d
		groupUnresolvedCount: %d`,
			project.GetId(), project.GetName(), project.GetDeployId(), project.GetDeployAt(), project.GetNoticeTotalCount(), project.GetRejectionCount(), project.GetFileCount(), project.GetDeployCount(), project.GetGroupResolvedCount(), project.GetGroupUnresolvedCount())
	}
}
