package main

import (
	"flag"
	"log"

	"github.com/knry0329/gomque/que"
)

func main() {
	var host string
	var vhost string
	var user string
	var password string
	var port string
	var connection string
	var qName string
	var isDeq bool
	flag.StringVar(&host, "h", "", "connect host")
	flag.StringVar(&vhost, "v", "", "connect vhost")
	flag.StringVar(&user, "u", "", "username")
	flag.StringVar(&password, "p", "", "connect password")
	flag.StringVar(&port, "P", "", "connect port")
	flag.StringVar(&connection, "c", "", "connect identifier")
	flag.StringVar(&qName, "q", "", "queue name")
	flag.BoolVar(&isDeq, "deq", false, "dequeue flg")
	flag.Parse()
	msg := flag.Arg(0)
	if msg == "" {
		log.Fatalf("invalid message.")
	}
	if connection == "" {
		connection = "amqp://" + user + ":" + password + "@" + host + ":" + port + "/" + vhost
	}
	if isDeq {
		enqExec(connection, qName, msg)
	} else {
		// deqExec(connection, qName, msg)
	}
}

func enqExec(connection, qName, msg string) {
	enq, err := que.NewEnq(connection, qName)
	if err != nil {
		log.Fatalf("connect error. %v", err)
	}
	defer enq.Close()
	enq.Enqueue([]byte(msg))

}

// func deqExec(connection, qName, msg string) {
// 	deq, err := que.NewDeq(connection, qName)
// 	if err != nil {
// 		log.Fatalf("connect error. %v", err)
// 	}
// 	defer deq.Close()
// 	deq.Dequeue()

// }
