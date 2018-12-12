package commands

import (
	"../untils"
	"net"
)

func ExecuteFunction(conn *net.TCPConn, function int) (*ResultCommandClient, error) {
	var result *ResultCommandClient
	var commandServer CommandServer

	commandServer.Act = Action(function)

	err := WriteCommand(conn, &commandServer)

	if err != nil {
		untils.WriteMsgLogError(err)
		return result, err
	}

	result, err = ReadCommand(conn)

	if err != nil {
		untils.WriteMsgLogError(err)
		return result, err
	}

	if result.Act != Action(function) {
		untils.WriteMsgLog("No correct command")
		return result, err
	}

	return result, nil
}

func GetFunctions(conn *net.TCPConn) (*ResultCommandClient, error) {
	var result *ResultCommandClient
	var commandServer CommandServer

	commandServer.Act = CMD_GET_FUNCTIONS

	err := WriteCommand(conn, &commandServer)

	if err != nil {
		untils.WriteMsgLogError(err)
		return result, err
	}

	result, err = ReadCommand(conn)

	if err != nil {
		untils.WriteMsgLogError(err)
		return result, err
	}

	if result.Act != CMD_GET_FUNCTIONS {
		untils.WriteMsgLog("No correct command")
		return result, err
	}

	return result, nil
}

func IsLiveClient(conn *net.TCPConn) bool {
	var commandServer CommandServer
	var result *ResultCommandClient

	commandServer.Act = CMD_IS_LIVE

	err := WriteCommand(conn, &commandServer)

	if err != nil {
		untils.WriteMsgLogError(err)
		return false
	}

	result, err = ReadCommand(conn)

	if err != nil {
		untils.WriteMsgLogError(err)
		return false
	}

	if result.Act == CMD_IS_LIVE {
		return true
	}

	return false
}
