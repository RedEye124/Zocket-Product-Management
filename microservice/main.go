package main

import (
	"fmt"
	"imageprocessor"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"image_processing_queue", // Queue name
		true,                     // Durable
		false,                    // Auto-delete
		false,                    // Exclusive
		false,                    // No-wait
		nil,                      // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // Queue name
		"",     // Consumer name
		true,   // Auto-acknowledge
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// Process messages
	for msg := range msgs {
		// Assume msg.Body contains the URL of the image to process
		imageUrl := string(msg.Body)
		log.Printf("Received image URL: %s", imageUrl)

		// Download, compress, and upload the image
		filename, err := imageprocessor.DownloadImage(imageUrl)
		if err != nil {
			log.Printf("Failed to download image: %v", err)
			continue
		}

		compressedImagePath, err := imageprocessor.CompressImage(filename)
		if err != nil {
			log.Printf("Failed to compress image: %v", err)
			continue
		}

		compressedImageURL, err := imageprocessor.UploadToS3(compressedImagePath)
		if err != nil {
			log.Printf("Failed to upload to S3: %v", err)
			continue
		}

		// Update the database with the compressed image URL
		// You can update the product record in PostgreSQL with the compressed image URL.
		fmt.Println("Image processed and uploaded to:", compressedImageURL)
	}
}
