package tcp

import (
	"bufio"
	"fmt"
	"log/slog"
	"net"
	"scraper-url/internal/crawler/spider"
)

type Server struct {
	log     *slog.Logger
	Host    string
	service *spider.Service
}

func New(log *slog.Logger, host string, service *spider.Service) *Server {
	return &Server{
		log:     log,
		Host:    host,
		service: service,
	}
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp4", s.Host)
	if err != nil {

		return fmt.Errorf("tcp serve can't start")
	}

	s.log.Info("tcp server is listening", "addres", s.Host)

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err //  fix IT

		}

		go s.handle(conn)
	}
}

func (s *Server) handle(c net.Conn) {
	defer c.Close()

	s.log.Debug("started connect with client")

	for {
		userInput, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			return
		}

		//userInput = strings.TrimLeft(userInput, "\n")

		res := s.service.Index.GetUrls(userInput)
		if len(res) == 0 {
			c.Write([]byte("no urls found\n"))
			continue
		}

		for _, url := range res {
			c.Write([]byte(url + "\n"))
		}

	}
}
