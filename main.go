package main

import (
	"bytes"
	"first-proto/models"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
)

func main() {

	var user1 = &models.User{
		Id:       "u001",
		Name:     "Muklis",
		Password: "welcome1",
		Gender:   models.UserGender_MALE,
	}

	var userList = &models.UserList{
		List: []*models.User{
			user1,
		},
	}

	var garage1 = &models.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		GarageCoordiante: &models.GarageCoordinate{
			Latitude:  23.0003030,
			Longitude: 53.0090999,
		},
	}

	var garageList = &models.GarageList{
		List: []*models.Garage{
			garage1,
		},
	}

	fmt.Printf("#==========	Original\n  %#v \n", user1)
	fmt.Printf("#==========	As String\n  %#v \n", user1.String())
	fmt.Printf("#==========	Original\n  %#v \n", userList)
	fmt.Printf("#==========	Original\n  %#v \n", garage1)
	fmt.Printf("#==========	As List\n  %#v \n", garage1.String())
	fmt.Printf("#==========	Original\n  %#v \n", garageList)

	var buf bytes.Buffer

	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)

	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()
	fmt.Printf("#========== As Json String\n   %v \n", jsonString)

	buf2 := strings.NewReader(jsonString)

	protoObject := new(models.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}
	fmt.Printf("#================= As Proto Object\n %v \n", protoObject.String())

}
