package http

import (
	"net"
)

type RequestHandler func (r *Request)

type HttpServer struct {
  listener net.Listener
  getHandlers map[string]RequestHandler
}

func Create(address string) (*HttpServer, error) {
  listener, err := net.Listen("tcp4", address)
  
  if err != nil {
    return nil, err
  } 

  return &HttpServer{
    listener: listener,
    getHandlers: make(map[string]RequestHandler),
  }, nil
}

func (s *HttpServer) Close() {
  s.listener.Close()
}

func (s *HttpServer) Listen() {
  for {
    c, _ := s.listener.Accept()
    go s.handleConnection(c)
  }
}

func (s *HttpServer) Get(route string, handler RequestHandler) {
  s.getHandlers[route] = handler
}

func (s *HttpServer) handleConnection(c net.Conn) {
  defer c.Close()

  request, err := InitRequest(c)

  if err != nil {
    return
  }

  handler, ok := s.getHandlers[request.Route]

  if !ok {
    c.Write([]byte("HTTP/1.1 404 Not Found"))
  }

  handler(request)
}
