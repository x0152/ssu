package commands

import (
	"../network"
	"bytes"
	"encoding/json"
	"net"
)

func ReadCommand(conn *net.TCPConn) (*ResultCommandClient, error) {

	var buffer bytes.Buffer
	var result ResultCommandClient

	err := network.ReadData(conn, &buffer)

	if err != nil {
		return &result, err
	}

	err = json.Unmarshal(buffer.Bytes(), &result)

	if err != nil {
		return &result, err
	}

	return &result, nil
}
