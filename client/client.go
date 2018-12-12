package client

import (
	"../commands"
	"../untils"
	"net"
	"time"
)

type Client struct {
	Id                  untils.Id
	IsConnect           bool
	Ip                  string
	Connection          *net.TCPConn
	DateTimeLastConnect time.Time
	DateTimeLastAnswer  time.Time
	Functions           *commands.ResultCommandClient
}
