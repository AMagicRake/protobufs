package main

import (
	"fmt"
	pb "proto_demo/proto"

	"google.golang.org/protobuf/proto"
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

func do_oneof(message any) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("message has unexpeted type: %v", x)
	}
}

func do_map() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	message := &pb.Simple{}
	readFromFile(path, message)
	fmt.Println(message)
}

func main() {
	fmt.Println(do_simple())
	fmt.Println(do_complex())
	fmt.Println(do_enum())

	fmt.Println("This should be an Id:")
	do_oneof(&pb.Result_Id{Id: 42})

	fmt.Println("This should be an Message:")
	do_oneof(&pb.Result_Message{Message: "a message is here"})

	fmt.Println(do_map())
	doFile(do_simple())
}
