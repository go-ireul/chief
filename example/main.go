package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"ireul.com/chief/types"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()
	client := types.NewChiefClient(conn)
	for i := 0; i < 1000; i++ {
		id, err := client.NewID(context.Background(), &types.NewIDRequest{Name: "test"})
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("ID:", id)
	}
}
