package main

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	ln := startApp()
	defer ln.Close()

	u := &url.URL{
		Scheme: "http",
		Host:   ln.Addr().String(),
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	server := http.Server{Handler: proxy}

	sl, _ := net.Listen("tcp", "127.0.0.1:0")
	defer sl.Close()
	go server.Serve(sl)

	clientConn, _ := net.Dial("tcp", sl.Addr().String())
	clientReq, _ := http.NewRequest("GET", "/", nil)
	writer := bufio.NewWriter(clientConn)
	clientReq.Write(writer)
	writer.Flush()
	reader := bufio.NewReader(clientConn)
	rsp, _ := http.ReadResponse(reader, clientReq)
	println(rsp.Status)
	clientConn.Close()
}

func startApp() net.Listener {
	ln, _ := net.Listen("tcp", "127.0.0.1:0")
	go func() {
		var tempDelay time.Duration // how long to sleep on accept failure
		for {
			conn, err := ln.Accept()
			if err != nil {
				if ne, ok := err.(net.Error); ok && ne.Temporary() {
					if tempDelay == 0 {
						tempDelay = 5 * time.Millisecond
					} else {
						tempDelay *= 2
					}
					if max := 1 * time.Second; tempDelay > max {
						tempDelay = max
					}
					println("http: Accept error: %v; retrying in %v", err, tempDelay)
					time.Sleep(tempDelay)
					continue
				}
				break
			}
			go func() {
				// reader := bufio.NewReader(conn)
				writer := bufio.NewWriter(conn)
				// req, _ := http.ReadRequest(reader)
				// println(req.URL.Path)
				rsp := &http.Response{
					StatusCode: http.StatusOK,
					ProtoMajor: 1,
					ProtoMinor: 1,
				}
				rsp.Write(writer)
				writer.Flush()
				conn.Close()
			}()
		}
	}()

	return ln
}
