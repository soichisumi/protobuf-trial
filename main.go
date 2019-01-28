package main

import (
	"fmt"
	"github.com/soichisumi/protobuf-trial/pbtrial"
	"log"
	"net/http"
)

var (
	userDB map[string]pbtrial.User
)

func init(){
	userDB = make(map[string]pbtrial.User)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`this is root of test api.`)))
	})

	// api without proto service
	http.HandleFunc("/norm/adduser", normAddUser)
	http.HandleFunc("/norm/getuser", normGetUser)

	fmt.Println("service is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Listen and serve:", err)
	}
}