package main
import(
	"log"
)
func main() {
    queues := []string{"q1", "q2", "q3", "q4"}

    forever := make(chan bool)
    for _, queue := range queues {
        go processQueue(queue)
    }
    <-forever
}

func processQueue(name string) {
    conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
    defer conn.Close()

    ch, _ := conn.Channel()
    defer ch.Close()

    msgs, _ := ch.Consume(name, "test-dev", false, false, false, false, nil)
    go func() {
        for d := range msgs {
            log.Printf("[%s] %s", name, d.RoutingKey)
            d.Ack(true)
        }
    }()
}