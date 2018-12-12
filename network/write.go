package network

import (
	"../untils"
	"bytes"
	"fmt"
	"net"
)

func WriteData(conn *net.TCPConn, buf bytes.Buffer) error {

	n, err := conn.Write(buf.Bytes())

	if err != nil {
		return err
	}

	if n == 0 {
		return fmt.Errorf("nul bytes was write")
	}

	untils.WriteMsgLog(fmt.Sprintf("server send %d bytes client (%s): %v", n, conn.RemoteAddr().String(), buf.String()))

	return nil
}
