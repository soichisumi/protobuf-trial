package main
// 実際に実装するときにはservice名のパッケージ内にServer structを作ってそこでglobalなオブジェクトを管理する

import (
	"encoding/json"
	"fmt"
	"github.com/GincoInc/go-global/utils"
	"github.com/golang/protobuf/proto"
	"github.com/soichisumi/protobuf-trial/pbtrial"
	"net/http"
)

func normAddUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("... normAddUser")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var jsonMap map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&jsonMap); err != nil {
		fmt.Printf("err: %+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("jsonmap: %+v\n", jsonMap)

	var user pbtrial.User
	err := utils.Parse(jsonMap, &user)
	if err != nil {
		fmt.Printf("parse error: %+v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Age == 0 {
		fmt.Printf("given user data is invalid: %+v", user)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, exists := userDB[user.Name]
	if exists {
		fmt.Printf("user is already registered. user: %+v", user)
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	userDB[user.Name] = user

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("the user is registered."))
}

func normGetUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("... normAddUser")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	keys, ok := r.URL.Query()["name"]
	if !ok || len(keys) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("key is missing or invalid. keys: %+v\n", keys)
		return
	}

	user, ok := userDB[keys[0]]
	if !ok {
		fmt.Printf("given user is missing. username: %s\n", keys[0])
		w.WriteHeader(http.StatusNotFound)
		return
	}
	bytes, err := proto.Marshal(&user)
	if err != nil {
		fmt.Printf("error in marshalling result. err: %+v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}