package handlers

import (
	"../client"
	"../consts"
	"../untils"
	"fmt"
	"net"
	"time"
)

func (cons *Connections) HandleTcpSocket() {
	servAddr := fmt.Sprintf(":%d", consts.TCP_PORT)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)

	CheckError(err)

	untils.WriteMsgLog(fmt.Sprintf("запуск сервера на порту %d...", consts.TCP_PORT))
	ln, err := net.ListenTCP("tcp", tcpAddr)

	//CheckError create tcp socket
	CheckError(err)

	for {

		conn, err := ln.AcceptTCP()

		mutex.Lock()
		//CheckError connect
		CheckError(err)

		var cl client.Client

		cl.IsConnect = true
		cl.Ip = conn.RemoteAddr().String()
		cl.Connection = conn

		cl.DateTimeLastConnect = time.Now()
		cl.DateTimeLastAnswer = time.Now()

		untils.WriteMsgLog(fmt.Sprintf("connected client (%s)", cl.Ip))

		id := untils.GenerationId()
		cl.Id = id
		cons.Clients[id] = &cl

		mutex.Unlock()

		//go client.HandlerClient(conn)
	}
}
