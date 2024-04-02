package mock

import (
	"adsb-api/internal/global"
	"adsb-api/internal/utility/logger"
	"net"
	"os"
	"time"
)

type TcpStub interface {
	StartServer() (string, error)
	Close() error
	SetResponse(response []byte)
}

type StubImpl struct {
	server   net.Listener
	response []byte
	addr     string
}

func InitStub(addr string, response []byte) *StubImpl {
	return &StubImpl{addr: addr, response: response}
}

func (stub *StubImpl) StartServer() (err error) {
	stub.server, err = net.Listen("tcp", stub.addr)
	if err != nil {
		return err
	}

	logger.Info.Println("starting mock TCP server")

	go func() {
		defer func(ln net.Listener) {
			err := ln.Close()
			if err != nil {
				return
			}
		}(stub.server)

		conn, err := stub.server.Accept()
		if err != nil {
			return
		}

		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				return
			}
		}(conn)

		_, err = conn.Write(stub.response)
		if err != nil {
			return
		}
	}()

	return nil
}

func (stub *StubImpl) Close() error {
	return stub.server.Close()
}

func (stub *StubImpl) SimulateConnectionDrop(delay time.Duration) {
	go func() {
		time.Sleep(delay)
		if stub.server != nil {
			err := stub.server.Close()
			if err != nil {
				logger.Error.Printf("error closing connection: %v", err)
			}
		}
	}()
}

func StartStubServer() {
	mockData, err := os.ReadFile("./resources/mock/mockSbsDataLen5.txt")
	if err != nil {
		logger.Error.Printf("error reading file: %q", err)
	}

	stub := InitStub(global.SbsSource, mockData)
	err = stub.StartServer()
	if err != nil {
		logger.Error.Fatalf("error starting stub server: %q", err)
		return
	}

}
