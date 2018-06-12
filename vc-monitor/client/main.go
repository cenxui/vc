package main

import (
	"log"
	"net"
	"os"
	"fmt"
	"bufio"
)

var (
	in = make(chan string)
	ou = make(chan string)
)

//!+
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	/**
	remote command
	 */
	go func() {

		input := bufio.NewScanner(conn)

		for input.Scan() {
			in <- "remote :" + input.Text()
		}

		//io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("remote done")
		done <- struct{}{} // signal the main goroutine
	}()


	/**
	local command
	 */

	 go func() {
	 	input := bufio.NewScanner(os.Stdin)

	 	for input.Scan() {
	 		in <- "local :" + input.Text()
		}
		 log.Println("local done")
		 done <- struct{}{} // signal the main goroutine
	 }()

	/**
	receive
	 */
	go func() {
		for {
			fmt.Println(<-in)
		}
	}()

	/**
	send
	 */
	func() {
		for msg := range ou {
			fmt.Fprintf(conn, msg);
		}
	}()

	conn.Close()
	<-done // wait for background goroutine to finish remote
	<-done // wait for background goroutine to finish local

}
