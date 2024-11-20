package main

import (
	pb "awesomeProject/src/go"
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"strconv"
	"time"
)

const (
// host     = "localhost"
// port     = 5432
// user     = "postgres"
// password = ""
// dbname   = "go_grpc_demo"
)

var db *sql.DB

func initDB() error {
	//var err error
	//
	//connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname)
	//
	//db, err = sql.Open("postgres", connectionString)
	//
	//if err != nil {
	//	return err
	//}
	//
	//err = db.Ping()
	//if err != nil {
	//	return err
	//}
	return nil
}

type taskServiceServer struct {
	pb.UnimplementedTaskServiceServer
}

func (s *taskServiceServer) CreateTask(ctx context.Context, request *pb.CreateTaskRequest) (*pb.Task, error) {
	var task = &pb.Task{
		Description: request.Description,
		UserId:      request.UserId,
		Deadline:    request.Deadline,
		CreatedAt:   timestamppb.Now(),
		Status:      pb.TaskStatus_TASK_STATUS_INCOMPLETE,
	}

	//var taskId int
	//
	//insertStmt := `
	//    INSERT INTO "tasks"("description", "user_id", "status", "deadline", "created_at")
	//    VALUES($1, $2, $3, $4, $5) RETURNING id;
	//`
	//err := db.QueryRow(insertStmt, task.Description, task.UserId, pb.TaskStatus_name[int32(task.Status)], task.Deadline.AsTime(), task.CreatedAt.AsTime()).Scan(&taskId)
	//if err != nil {
	//	return nil, err
	//}

	taskId := 123
	task.Id = strconv.Itoa(taskId)
	return task, nil
}

func (s *taskServiceServer) GetTask(ctx context.Context, request *pb.GetTaskRequest) (*pb.Task, error) {
	var (
		id          int
		description string
		user_id     int
		status      string
		deadline    string
		created_at  string
	)
	//err := db.QueryRow("SELECT * FROM tasks WHERE tasks.id = $1", request.TaskId).Scan(
	//	&id, &description, &user_id, &status, &deadline, &created_at)
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println(request.TaskId)
	id = 1
	description = "This is a fake task"
	user_id = 42
	status = "TASK_STATUS_INCOMPLETE"
	deadline = "2024-12-31T12:00:00Z"
	created_at = "2024-01-01T12:00:00Z"

	deadlineTime, err := time.Parse(time.RFC3339, deadline)
	if err != nil {
		log.Fatalf("Error: Invalid time for deadline: %v", err)
	}
	createdAtTime, err := time.Parse(time.RFC3339, created_at)
	if err != nil {
		log.Fatalf("Error: Invalid time for created_at: %v", err)
	}
	task := &pb.Task{
		Id:          strconv.Itoa(id),
		Description: description,
		UserId:      strconv.Itoa(user_id),
		Status:      pb.TaskStatus(pb.TaskStatus_value[status]),
		Deadline:    timestamppb.New(deadlineTime),
		CreatedAt:   timestamppb.New(createdAtTime),
	}
	return task, nil
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("Error initiating database: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9090))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, &taskServiceServer{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Error starting gRPC server: %v", err)
	}
}
