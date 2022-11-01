package main

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ibrahimker/golang-intermediate/session-3/clientserver/common/config"
	"github.com/ibrahimker/golang-intermediate/session-3/clientserver/common/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serviceGarage() model.GaragesClient {
	port := config.SERVICE_GARAGE_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return model.NewGaragesClient(conn)
}

func serviceUser() model.UsersClient {
	port := config.SERVICE_USER_PORT
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	return model.NewUsersClient(conn)
}

func main() {
	userSvc := serviceUser()
	garageSvc := serviceGarage()
	ctx := context.Background()

	user1 := &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "for the horde",
		Gender:   model.UserGender_FEMALE,
	}
	_, _ = userSvc.Register(ctx, user1)

	user2 := &model.User{
		Id:       "u002",
		Name:     "John Doe",
		Password: "for the horde",
		Gender:   model.UserGender_MALE,
	}
	_, _ = userSvc.Register(ctx, user2)

	users, _ := userSvc.List(ctx, new(empty.Empty))
	log.Printf("List Users %+v\n", users.GetList())

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  -6.10,
			Longitude: 107.08,
		},
	}
	_, _ = garageSvc.Add(ctx, &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: garage1,
	})
	_, _ = garageSvc.Add(ctx, &model.GarageAndUserId{
		UserId: user1.Id,
		Garage: garage1,
	})
	garages, _ := garageSvc.List(ctx, &model.GarageUserId{UserId: user1.Id})
	log.Printf("List garages %+v\n", garages.GetList())

}
