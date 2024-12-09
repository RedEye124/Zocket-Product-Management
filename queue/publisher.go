package queue

import (
	"github.com/streadway/amqp"
)

func PublishImageURL(url string) {
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")
	defer conn.Close()

	ch, _ := conn.Channel()
	defer ch.Close()

	q, _ := ch.QueueDeclare("image_queue", false, false, false, false, nil)
	ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(url),
	})
}
