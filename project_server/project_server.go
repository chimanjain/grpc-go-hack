package main

import (
	"context"
	"encoding/json"
	"errors"
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

func (s *projectManagementServer) GetProject(_ context.Context, projectID *pb.ProjectID) (*pb.Project, error) {
	projects := getTestProjects()
	for _, project := range projects {
		if project.GetId() == projectID.GetId() {
			return project, nil
		}
	}
	return nil, errors.New("project not found")
}

func (s *projectManagementServer) GetProjects(projectID *pb.ProjectID, stream pb.ProjectManagement_GetProjectsServer) error {
	projects := getTestProjects()
	for _, project := range projects {
		if project.GetId() == projectID.GetId() {
			if err := stream.Send(project); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *projectManagementServer) FetchProjects(stream pb.ProjectManagement_FetchProjectsServer) error {
	testProjects := getTestProjects()
	var IDs []uint32

	for {
		projectID, err := stream.Recv()
		IDs = append(IDs, projectID.GetId())
		if err == io.EOF {
			return stream.SendAndClose(&pb.Projects{
				Projects: filterProjects(testProjects, IDs),
			})
		}
	}
}

func filterProjects(testProjects []*pb.Project, projectIDs []uint32) []*pb.Project {
	var projects []*pb.Project
	for _, testProject := range testProjects {
		if contains(projectIDs, testProject.Id) {
			projects = append(projects, testProject)
		}
	}
	return projects
}

func (s *projectManagementServer) StreamProjects(stream pb.ProjectManagement_StreamProjectsServer) error {
	testProjects := getTestProjects()

	for {
		in, err := stream.Recv()

		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		projectID := in.GetId()
		for _, testProject := range testProjects {
			if testProject.GetId() == projectID {
				if err := stream.Send(testProject); err != nil {
					return err
				}
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProjectManagementServer(s, &projectManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func getTestProjects() []*pb.Project {
	var projects []*pb.Project
	projectFilePath := filepath.Join("test_data", "project_db.json")
	// Open our jsonFile
	projectFile, err := os.Open(projectFilePath)
	if err != nil {
		log.Printf("Failed to load projects: %v", err)
	}
	byteValue, _ := io.ReadAll(projectFile)
	// defer the closing of our projectFile so that we can parse it later on
	defer projectFile.Close()
	if err := json.Unmarshal(byteValue, &projects); err != nil {
		log.Printf("Failed to unmarshal projects: %v", err)
	}
	return projects
}

func contains(s []uint32, x uint32) bool {
	for i := range s {
		if uint32(i) == x {
			return true
		}
	}
	return false
}
