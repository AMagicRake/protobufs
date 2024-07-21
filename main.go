package main

import (
	"fmt"
	pb "proto_demo/proto"
)

func do_simple() *pb.Simple {
	return &pb.Simple{
		Id:         42,
		IsSimple:   true,
		Name:       "Niel",
		SampleList: []int32{1, 2, 3, 4, 5, 6},
	}
}

func do_complex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   1,
			Name: "A Name",
		},
		MultiDummy: []*pb.Dummy{
			{Id: 2, Name: "Name 2"},
			{Id: 3, Name: "Name 3"},
			{Id: 4, Name: "Name 4"},
		},
	}
}

func do_enum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLUE,
	}
}

func main() {
	fmt.Println(do_simple())
	fmt.Println(do_complex())
	fmt.Println(do_enum())
}
