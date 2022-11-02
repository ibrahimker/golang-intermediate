package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/ibrahimker/golang-intermediate/session-3/introduction/model"
)

func main() {
	// import dari model proto
	user1 := &model.User{
		Id:       "u001",
		Name:     "Sylvana Windrunner",
		Password: "for the horde",
		Gender:   model.UserGender_FEMALE,
	}

	log.Printf("user1 %#v\n", user1)
	log.Printf("user1.String() %#v\n", user1.String())

	// tes jsonpb marshal (marshalling proto to json)
	var (
		buf bytes.Buffer
	)
	_ = (&jsonpb.Marshaler{}).Marshal(&buf, user1)
	log.Printf("user1.jsonString %#v\n", buf.String())

	// tes jsonpb unmarshal (json string to proto)
	buf2 := strings.NewReader(buf.String())
	var user2 model.User
	_ = jsonpb.Unmarshal(buf2, &user2)
	log.Printf("user2 %#v\n", user2)
	// create user list and link to garage
	userList := &model.UserList{
		List: []*model.User{user1},
	}

	log.Println("userList", userList)

	garage1 := &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  -6.10,
			Longitude: 107.08,
		},
	}

	garageList := &model.GarageList{
		List: []*model.Garage{garage1},
	}

	garageListByUser := &model.GarageListByUser{
		List: map[string]*model.GarageList{
			user1.Id: garageList,
		},
	}

	log.Println("garageListByUser", garageListByUser)
}
