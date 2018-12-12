package commands

import (
	"../network"
	"bytes"
	"encoding/json"
	"net"
)

func WriteCommand(conn *net.TCPConn, command *CommandServer) error {
	b, err := json.Marshal(command)

	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	buffer.Write(b)

	err = network.WriteData(conn, buffer)

	if err != nil {
		return err
	}

	return nil
}
