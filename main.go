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

func main() {
	fmt.Println(do_simple())
}
