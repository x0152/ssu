package main

import (
	"./client"
	"./consts"
	"./handlers"
	"./untils"
	"fmt"
	"log"
	"net/http"
	"os"
)

func InitLog() {
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	log.SetOutput(f)
}

func main() {

	InitLog()

	var connections handlers.Connections

	connections.Clients = make(map[untils.Id]*client.Client)

	go connections.HandleTcpSocket()

	untils.WriteMsgLog(fmt.Sprintf("запуск http-сервера на порту %d...", consts.HTTP_PORT))
	http.HandleFunc("/", connections.HandlerClients)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) { return })
	http.HandleFunc("/function/", connections.HandlerFunctionClient)
	http.HandleFunc("/function/execute/", connections.HandlerExecuteFunction)
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("template/styles"))))

	untils.WriteMsgLogError(http.ListenAndServe(fmt.Sprintf(":%d", consts.HTTP_PORT), nil))
}
