package main

import "net"
import "fmt"
import "bufio"
import "strings" // only needed below for sample processing

func main() {

  fmt.Println("Launching server...")

  // listen on all interfaces
  ln, _ := net.Listen("tcp", ":8081")

  // accept connection on port
  conn, _ := ln.Accept()

  // run loop forever (or until ctrl-c)
  for {
	//var hash_string, skey_initial string
	protectorserv := new(Protector)
	// will listen for message to process ending in newline (\n)
	hash_string, _ := bufio.NewReader(conn).ReadString('\n')
    skey_initial, _ := bufio.NewReader(conn).ReadString('\n')	
	hash_string = strings.Trim(hash_string,"\n")
	skey_initial = strings.Trim(skey_initial,"\n")
	protectorserv.set_key(skey_initial)
    protectorserv.set_hash(hash_string)  
	// output message received
	fmt.Println("Hash Received:", protectorserv.get_hash())
	fmt.Println("key Received:", protectorserv.get_key())
	new_key := protectorserv.next_session_key()
	fmt.Println("Текущий ключ: " + new_key+"\n")
	// send new string back to client
	conn.Write([]byte(new_key + "\n"))
  }
}