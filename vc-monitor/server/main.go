package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
	"os"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8000");

	if err != nil {
		log.Fatal("error")
	}

	go localMessage()

	go sendMessage()

	for {
		conn, err := listen.Accept()

		if err != nil{
			log.Print("connt error")
			continue
		}

		go handleRequest(conn)
	}

}

type client chan string

type clientMap map[string] client

type clientUnit struct {
	name string
	client client
}

var (
	enter = make(chan clientUnit)
	leave = make(chan clientUnit)
	message = make(chan string)
)

func handleRequest(conn net.Conn)  {
	ch := make(client)
	go clientWriter(conn, ch)

	who :=  conn.RemoteAddr().String()

	ch <- "You are " + who + "\n"

	fmt.Println(who + " is online")

	message <- who + " is entering"
	cu := clientUnit{name:who, client:ch}

	enter <- cu

	input := bufio.NewScanner(conn)
	for input.Scan() {
		message <- who + ": " + input.Text() + "\n"
	}

	leave <- cu
	message <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch client)  {
	for msg := range ch {
		fmt.Fprintf(conn, msg);
	}
}

func localMessage()  {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		message <-  input.Text()
	}
}

func remoteMessage()  {

}

func sendMessage()  {
	clients := make(clientMap)

	for {
		select {
		case s := <- message:
			command(s, clients)
		case cli := <- enter:
			clients[cli.name] = cli.client

		case cli := <- leave:
			delete(clients,cli.name);
			close(cli.client)
		}

	}
}

func command(c string, cm clientMap)  {
	if strings.HasPrefix(c, "ls") {
		for k, _ := range cm {
			fmt.Println(k)
		}
	}else if strings.HasPrefix(c, "cl") {
		n := strings.TrimPrefix(c, "cl")
		fmt.Println(n)
		ch := cm[n]
		if ch == nil{
			fmt.Println("no this client")
		}else {
			ch <- "close\n"
		}
	}else if strings.HasPrefix(c, "op") {
		n := strings.TrimPrefix(c, "op")
		fmt.Println(n)
		ch := cm[n]
		if ch == nil{
			fmt.Println("client not found")
		}else {
			ch <- "open\n"
		}
	}

}