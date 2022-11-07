package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ibrahimker/golang-intermediate/session-3/todo/common/config"
	"github.com/ibrahimker/golang-intermediate/session-3/todo/common/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TodoServer struct {
	model.UnimplementedTodoServiceServer
}

var Todos = make(map[string]*model.Todo)

func main() {
	srv := grpc.NewServer()
	userSrv := new(TodoServer)
	model.RegisterTodoServiceServer(srv, userSrv)

	log.Println("Starting Todo Server at ", config.SERVICE_TODO_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_TODO_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}

	// setup http proxy
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
		grpcServerEndpoint := flag.String("grpc-server-endpoint", "localhost"+config.SERVICE_TODO_PORT, "gRPC server endpoint")
		_ = model.RegisterTodoServiceHandlerFromEndpoint(context.Background(), r, *grpcServerEndpoint, opts)
		log.Println("Starting Todo Server HTTP at 9001 ")
		http.ListenAndServe(":9001", mux)
	}()

	log.Fatalln(srv.Serve(listener))
}

func (t *TodoServer) GetAll(ctx context.Context, req *emptypb.Empty) (*model.GetAllResponse, error) {
	var todos []*model.Todo
	for _, v := range Todos {
		todos = append(todos, &model.Todo{
			Id:   v.GetId(),
			Name: v.GetName(),
		})
	}
	return &model.GetAllResponse{Data: todos}, nil
}
func (t *TodoServer) GetByID(ctx context.Context, req *model.GetByIDRequest) (*model.GetByIDResponse, error) {
	todo, ok := Todos[req.GetId()]
	if !ok {
		return nil, errors.New("not found")
	}
	return &model.GetByIDResponse{Data: todo}, nil
}
func (t *TodoServer) Create(ctx context.Context, req *model.Todo) (*model.MutationResponse, error) {
	Todos[req.GetId()] = &model.Todo{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	msg := req.GetId() + "successfully appended"
	return &model.MutationResponse{Success: msg}, nil
}
func (t *TodoServer) Update(ctx context.Context, req *model.UpdateRequest) (*model.MutationResponse, error) {
	Todos[req.GetId()] = &model.Todo{
		Id:   req.GetId(),
		Name: req.GetName(),
	}
	msg := req.GetId() + "successfully appended"
	return &model.MutationResponse{Success: msg}, nil
}
func (t *TodoServer) Delete(ctx context.Context, req *model.DeleteRequest) (*model.MutationResponse, error) {
	delete(Todos, req.GetId())
	msg := req.GetId() + "successfully deleted"
	return &model.MutationResponse{Success: msg}, nil
}
