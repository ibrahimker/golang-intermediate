package main

import (
	"context"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ibrahimker/golang-intermediate/session-3/clientserver/common/config"
	"github.com/ibrahimker/golang-intermediate/session-3/clientserver/common/model"
	"google.golang.org/grpc"
)

var localStorage *model.UserList

func init() {
	localStorage = new(model.UserList)
	localStorage.List = make([]*model.User, 0)
}

type UsersServer struct {
	model.UnimplementedUsersServer
}

func (u UsersServer) Register(ctx context.Context, req *model.User) (*empty.Empty, error) {
	log.Printf("Register user request %+v\n", req)
	localStorage.List = append(localStorage.List, req)
	log.Println("Registering user", req.String())

	return new(empty.Empty), nil
}

func (u UsersServer) List(context.Context, *empty.Empty) (*model.UserList, error) {
	log.Printf("List user request\n")
	return localStorage, nil
}

func main() {
	srv := grpc.NewServer()
	userSrv := UsersServer{}
	model.RegisterUsersServer(srv, userSrv)

	log.Println("Starting User Server at ", config.SERVICE_USER_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_USER_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}
	log.Fatalln(srv.Serve(listener))
}
