package main

import (
	"fmt"
	"github.com/soichisumi/protobuf-trial/pbtrial"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

//var (
//	userDB map[string]pbtrial.User // normal implementationで使用
//)

const (
	port = "8080"
)

//func init(){
//	userDB = make(map[string]pbtrial.User)
//}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %+v\n", err)
	}
	server := NewServer()
	s := grpc.NewServer()

	pbtrial.RegisterUserServiceServer()

}

// normal implementation

//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//	w.Write([]byte(fmt.Sprintf(`this is root of test api.`)))
//})
//
//// api without proto service
//http.HandleFunc("/norm/adduser", normAddUser)
//http.HandleFunc("/norm/getuser", normGetUser)
//
//fmt.Println("service is running on port 8080...")
//if err := http.ListenAndServe(":8080", nil); err != nil {
//	log.Fatal("Listen and serve:", err)
//}