package znet

import (
	"Zmin/engine/zlog"
	"Zmin/engine/zutil"
	"net"
	"time"
)

const (
	_RESTART_TCP_SERVER_INTERVAL = 3 * time.Second
	_RESTART_UDP_SERVER_INTERVAL = 3 * time.Second
)

// ServeTCPForever serves on specified address as TCP server, for ever ...
func ServeTCPForever(listenAddr string, service IService) {
	for {
		err := serveTCPForeverOnce(listenAddr, service)
		zlog.Errorf("server@%s failed with error: %v, will restart after %s", listenAddr, err, _RESTART_TCP_SERVER_INTERVAL)
		time.Sleep(_RESTART_TCP_SERVER_INTERVAL)
	}
}

func serveTCPForeverOnce(listenAddr string, service IService) error {
	defer func() {
		if err := recover(); err != nil {
			zlog.Error("serveTCPImpl: paniced with error %s", err)
		}
	}()

	return ServeTCP(listenAddr, service)

}

// ServeTCP serves on specified address as TCP server
func ServeTCP(listenAddr string, service IService) error {
	ln, err := net.Listen("tcp", listenAddr)
	zlog.Infof("Listening on TCP: %s ...", listenAddr)

	if err != nil {
		return err
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			if zutil.IsTimeoutError(err) {
				continue
			} else {
				return err
			}
		}
		zlog.Infof("Connection from: %s", conn.RemoteAddr())
		go service.NewTcpConnection(conn)
	}
}