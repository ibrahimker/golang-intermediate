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

var localStorage *model.GarageListByUser

func init() {
	localStorage = new(model.GarageListByUser)
	localStorage.List = make(map[string]*model.GarageList, 0)
}

type GaragesServer struct {
	model.UnimplementedGaragesServer
}

func (g GaragesServer) List(ctx context.Context, req *model.GarageUserId) (*model.GarageList, error) {
	return localStorage.List[req.UserId], nil

}
func (g GaragesServer) Add(ctx context.Context, req *model.GarageAndUserId) (*empty.Empty, error) {
	uid := req.GetUserId()
	if _, ok := localStorage.List[uid]; !ok {
		localStorage.List[uid] = new(model.GarageList)
		localStorage.List[uid].List = make([]*model.Garage, 0)
	}
	localStorage.List[uid].List = append(localStorage.List[uid].List, req.GetGarage())
	log.Println("Adding garage", req.GetGarage(), "for user", uid)
	return new(empty.Empty), nil
}

func main() {
	srv := grpc.NewServer()
	garageSrv := GaragesServer{}
	model.RegisterGaragesServer(srv, garageSrv)

	log.Println("Starting Garages Server at ", config.SERVICE_GARAGE_PORT)

	listener, err := net.Listen("tcp", config.SERVICE_GARAGE_PORT)
	if err != nil {
		log.Fatalf("could not listen. Err: %+v\n", err)
	}
	log.Fatalln(srv.Serve(listener))
}
