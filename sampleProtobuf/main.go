package main

import (
	"demo/pb"
	"fmt"

	"google.golang.org/protobuf/proto"
)

func main() {
	user := &pb.User{
		Name: "Saiful",
		Age:  25,
		Details: &pb.Details{
			Address: "lakshmipur",
			Phone:   "+8801626692265",
		},
	}

	//using auto generated get function to fetch data
	fmt.Println(user.GetName())
	fmt.Println(user.GetAge())
	fmt.Println(user.Details.GetAddress())
	fmt.Println(user.Details.GetPhone())

	//making the data in serialization. Basically this format will use for server communication
	serializedData, err := proto.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serialized data:", serializedData)

	// Now let's unmarshal the data back to a new User message
	// Make sure to create a new empty User message to store the unmarshaled data
	newUser := &pb.User{}
	err = proto.Unmarshal(serializedData, newUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Unmarshaled user:", newUser)
}
