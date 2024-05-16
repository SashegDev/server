package main
import (
	"fmt"
	"log"
	"net"
)

var (
	listener     net.Listener
	err          error
	clients_conn []net.Conn
	clients_name []string
	
	HOST   = ""
	PORT   string
	BUFSIZ = 1024
)
	
func accept_incomming_connections(a chan int) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Client %v connected\n", conn)
		go handle_client(conn)
	}
}

func handle_client(conn net.Conn) {
	name := make([]byte, BUFSIZ)
	count, err := conn.Read(name)
	if err != nil {
		log.Fatal(err)
	}
}

conn.Write([]byte("BG Messanger TEST 1.1"))
msg := fmt.Sprintf("%s - joined", name[:count])
broadcast([]byte(msg), "")
clients_conn = append(clients_conn, conn)
clients_name = append(clients_name, string(name[:count]))

for {
	nmsg := make([]byte, BUFSIZ)
	count, err := conn.Read(nmsg)
	if err != nil {
		fmt.Println("[!] Client has disconnected!")
		break
	}
		if string(nmsg[:count]) != "{quit}" || string(nmsg[:count]) != "/leave" || string(nmsg[:count]) != "/quit" {
		broadcast(nmsg, " ")
	} else {
		conn.Write([]byte("{quit}"))
		conn.Close()
		broadcast([]byte(fmt.Sprintf("%s - left", string(name[:count]))), "")
		break
	}
}

func broadcast(msg []byte, prefix string) {
	for _, sock := range clients_conn {
		sock.Write([]byte(string(prefix) + string(msg)))
	}
}
	
func main() {
	fmt.Print("Input PORT: ")
	fmt.Scanf("%s", &PORT)
	
	listener, err = net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Please wait...")
	fmt.Println("\nP2P Server Started!")
	fmt.Println("waiting connections...")
	a := make(chan int)
	go accept_incomming_connections(a)
	<-a
}