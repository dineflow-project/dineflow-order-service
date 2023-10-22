package rabbitmq

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type Notification struct {
	RecipientID string `json:"recipient_id"`
	OrderID     string `json:"order_id"`
	IsRead      bool   `json:"is_read"`
	Type        string `json:"type"`
	Timestamp   string `json:"timestamp"`
}

func PushNotification(recipientID string, orderID string, notiType string) error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to RabbitMQ server
	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		return err
	}
	defer conn.Close()

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// Declare a queue
	q, err := ch.QueueDeclare(
		os.Getenv("NOTI_QUEUE_NAME"), // name
		false,                        // durable
		false,                        // delete when unused
		false,                        // exclusive
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		return err
	}

	// Create a notification message
	notification := Notification{
		RecipientID: recipientID,
		OrderID:     orderID,
		IsRead:      false,
		Type:        notiType,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	// Convert notification to JSON
	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	// Publish the message to the queue
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Printf("Sent message to queue: %v", notification)

	return nil
}
