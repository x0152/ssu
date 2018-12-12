package network

import (
	"../consts"
	"../untils"
	"bytes"
	"fmt"
	"net"
)

var Cash map[*net.TCPConn][]byte

func init() {
	Cash = make(map[*net.TCPConn][]byte)
}

func BrackdownPackage(buf []byte) (newPackage []byte, surplus []byte) {
	var validBacket int

	var countChars int
	var countBacket int
	for _, b := range buf {
		switch b {
		case '{':
			validBacket++
			countBacket++
		case '}':
			validBacket--
			countBacket++
		}

		countChars++
		if countBacket > 0 && validBacket == 0 {
			return buf[:countChars], buf[countChars:]
		}
	}

	return nil, buf
}

func ReadData(conn *net.TCPConn, buf *bytes.Buffer) error {

	for {
		b := make([]byte, consts.SIZE_BUFFER_RAW_DATA)
		n, err := conn.Read(b)
		b = b[:n]

		if err != nil {
			Cash[conn] = Cash[conn][:0]
			return err
		}

		if n == 0 {
			Cash[conn] = Cash[conn][:0]
			return fmt.Errorf("null bytes was read client (%s)", conn.RemoteAddr().String())
		}

		b = append(Cash[conn], b...)

		var newPackage []byte
		newPackage, Cash[conn] = BrackdownPackage(b)

		if len(newPackage) != 0 {
			untils.WriteMsgLog(fmt.Sprintf("client (%s) send %d bytes: %s", conn.RemoteAddr().String(), len(newPackage), string(newPackage)))

			buf.Write(newPackage)
			return nil
		}

	}
}
