package main

import (
	"encoding/json"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	VERBOSE   = kingpin.Flag("verbose", "Make the application real chatty.").Short('v').Bool()
	CONN_TYPE = kingpin.Flag("connection", "Make the application real chatty.").Short('c').Default("tcp").String()
	CONN_HOST = kingpin.Flag("host", "Make the application real chatty.").Short('H').Default("localhost").String()
	CONN_PORT = kingpin.Flag("port", "Make the application real chatty.").Short('p').Default("91337").Int64()
)

func main() {
	kingpin.Parse()
	if *VERBOSE {
		log.SetLevel(log.DebugLevel)
	}

	l, err := net.Listen(*CONN_TYPE, fmt.Sprintf("%s:%d", *CONN_HOST, *CONN_PORT))
	if err != nil {
		log.Fatalf("error listening: %s", err)
	}
	defer l.Close()
	log.Debugf("Listening on %s:%d", *CONN_HOST, *CONN_PORT)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalf("error accepting: %s", err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Errorf("error reading: %s", err)
	}

	var a Action
	err = json.Unmarshal(buf[:n], &a)
	if err != nil {
		log.Errorf("difficulity understanding message: %s", err)
	}

	res, err := a.Distribute()
	var data []byte
	if err != nil {
		log.Errorf("action distribution: %s", err)
		data, err = json.Marshal(NewErrorResponse(err))
	} else {
		data, err = json.Marshal(res)
	}
	if err != nil {
		log.Errorf("difficulity to express feelings: %s", err)
	}

	conn.Write([]byte(data))
	conn.Close()
}
