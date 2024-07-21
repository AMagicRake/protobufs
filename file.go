package main

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fname string, pb proto.Message) {
	outData, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
	}

	if err = os.WriteFile(fname, outData, 0644); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Data written")
}

func readFromFile(fname string, pb proto.Message) {
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Couldn't Unmarshall file, ", err)
	}

	fmt.Println(pb)
}
