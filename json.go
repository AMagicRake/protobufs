package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJson(message proto.Message) string {
	out, err := protojson.Marshal(message)
	if err != nil {
		log.Fatalln("Failed to marshal", err)
	}

	return string(out)
}

func fromJson(in string, message proto.Message) {
	if err := protojson.Unmarshal([]byte(in), message); err != nil {
		log.Fatalln("Couldn't unmarshall string", err)
	}
}
