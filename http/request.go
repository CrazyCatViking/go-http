package http

import (
	"bufio"
	"net"
	"strings"
)

type Request struct {
  connection net.Conn
  Method string
  Route string
}

func (r *Request) String(content string) {
  r.connection.Write([]byte("HTTP/1.1 200 Success\r\n"))
  r.connection.Write([]byte("Content-Type: text/plain\r\n"))
  r.connection.Write([]byte("\r\n"))
  r.connection.Write([]byte(content))
}

func (r *Request) Html(content string) {
  r.connection.Write([]byte("HTTP/1.1 200 Success\r\n"))
  r.connection.Write([]byte("Content-Type: text/html\r\n"))
  r.connection.Write([]byte("\r\n"))
  r.connection.Write([]byte(content))
}

func InitRequest(connection net.Conn) (*Request, error) {
  reader := bufio.NewReader(connection)

  requestLine, _, err := reader.ReadLine()

  if (err != nil) {
    return nil, err
  } 

  parts := strings.Fields(string(requestLine))

  method := parts[0]
  route := parts[1]

  request := Request{
    connection: connection,
    Method: method,
    Route: route,
  }

  return &request, nil
}
