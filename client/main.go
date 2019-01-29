package main

import (
	"context"
	"fmt"
	"github.com/soichisumi/protobuf-trial/pbtrial"
	"google.golang.org/grpc"
	"log"
)

const(
	port= ":8080"
	target = "localhost:8080"
)

var (
	ctx = context.Background()
)

func main(){
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer conn.Close()

	client := pbtrial.NewUserServiceClient(conn)
	req := pbtrial.AddUserRequest{
		User: &pbtrial.User{
			Name: "yoyo",
			Age: 123,
		},
	}

	res, err := client.AddUser(ctx, &req)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	fmt.Printf("res: %+v\n", res)

	req2 := pbtrial.GetUserRequest{
		Username: "yoyo",
	}

	res2, err := client.GetUser(ctx, &req2)
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	fmt.Printf("res: %+v\n", res2)
}
