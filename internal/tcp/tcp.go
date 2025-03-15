package tcp

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"scraper-url/internal/crawler/spider"
)

func ListenAndServe(log *slog.Logger, host string, service *spider.Service) error {
	l, err := net.Listen("tcp4", host)
	if err != nil {

		return fmt.Errorf("tcp serve can't start")
	}

	log.Info("tcp server is listening", "addres", host)

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err //  fix IT

		}

		go handle(log, conn)
	}
}

func handle(log *slog.Logger, c net.Conn) {
	defer c.Close()

	log.Debug("started connect with client")

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return
		}

		c.Write([]byte(userInput))
	}
}
