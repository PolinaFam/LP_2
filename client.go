package main

import "net"
import "fmt"
import "bufio"
import "os"
import "strings"

func main() {

  // connect to this socket
  conn, _ := net.Dial("tcp", "127.0.0.1:8081")
  for { 
    var hash_string,skey_initial string
    protectorcl := new(Protector) 
    hash_string = protectorcl.get_hash_str()
    skey_initial = protectorcl.get_session_key()
    protectorcl.set_key(skey_initial)
    protectorcl.set_hash(hash_string)
    // read in input from stdin
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Ваше сообщение: ")
    text, _ := reader.ReadString('\n')
    // send to socket
    fmt.Fprintf(conn, hash_string + "\n")
    fmt.Fprintf(conn, skey_initial + "\n")
    // listen for reply
    key_mess, _ := bufio.NewReader(conn).ReadString('\n')
    serv_key := strings.Trim(key_mess,"\n")
    new_key:=protectorcl.next_session_key()
    fmt.Println("Текущий ключ: " + new_key) 
    if new_key == serv_key {
        fmt.Println("Ключи совпали ",text)
    }
  }
}