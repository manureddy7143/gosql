package main
import
(
	"bufio"
	"strings"
	"fmt"
	"net"
	"os"
	"math/rand"
	"time"
)
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	
	var queue []string
	for {   c.Write([]byte("please type 'SEND' ,'GET', 'STOP' and PRESS ENTER for sending,getting,stoping conversation\n"))
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println("connction closed for ",c.RemoteAddr().String())
					break
					return

			}
            
			temp := strings.TrimSpace(string(netData))
			switch {
			case temp=="1":
		        c.Close()
			case temp=="2":
				netData, _:= bufio.NewReader(c).ReadString('\n')
				queue = append(queue, netData)
				c.Write([]byte("msg sent\n"))
				fmt.Println("msg received from",c.RemoteAddr().String())
			case temp=="3":
				if len(queue) > 0 {
					fmt.Print(queue[0]) // First element
					queue = queue[1:]
					c.Write([]byte("msg received\n"))
				    fmt.Println("msg sent to",c.RemoteAddr().String())
				}
		    default:
				c.Write([]byte("enter valid text\n"))
			}

	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
			fmt.Println("Please provide a port number!")
			return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
			fmt.Println(err)
			return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
			c, err := l.Accept()
			if err != nil {
					fmt.Println(err)
					return
			}
			go handleConnection(c)
	}
}