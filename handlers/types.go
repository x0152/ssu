package handlers

import (
	"../client"
	"../untils"
	"sync"
)

var mutex = &sync.Mutex{}

type Connections struct {
	/*key - id, value -> ptrStruct*/
	Clients map[untils.Id]*client.Client
}
