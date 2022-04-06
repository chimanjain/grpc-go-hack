package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/chimanjain/grpc-go-hack/project"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type projectManagementServer struct {
	pb.UnimplementedProjectManagementServer
}

func (s *projectManagementServer) GetProject(ctx context.Context, in *pb.ProjectId) (*pb.Project, error) {
	return &pb.Project{
		Id:                   in.GetId(),
		Name:                 "Airbrake project name",
		DeployId:             "1",
		DeployAt:             "2022-04-26T17:37:33.638348Z",
		NoticeTotalCount:     1,
		RejectionCount:       1,
		FileCount:            1,
		DeployCount:          1,
		GroupResolvedCount:   1,
		GroupUnresolvedCount: 1,
	}, nil
}

func (s *projectManagementServer) GetProjects(projectId *pb.ProjectId, stream pb.ProjectManagement_GetProjectsServer) error {
	var projects []*pb.Project
	projectFilePath := filepath.Join("test_data", "project_db.json")
	// Open our jsonFile
	projectFile, err := os.Open(projectFilePath)
	if err != nil {
		log.Fatalf("Failed to load projects: %v", err)
	}
	byteValue, _ := io.ReadAll(projectFile)
	// defer the closing of our projectFile so that we can parse it later on
	defer projectFile.Close()
	if err := json.Unmarshal(byteValue, &projects); err != nil {
		log.Fatalf("Failed to unmarshal projects: %v", err)
	}
	for _, project := range projects {
		if err := stream.Send(project); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProjectManagementServer(s, &projectManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
