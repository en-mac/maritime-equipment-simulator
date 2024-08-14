package simulator

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    "github.com/streadway/amqp"
)

func Start() {
    // Get the RabbitMQ connection string from environment variables
    rabbitMQURL := os.Getenv("RABBITMQ_URL")
    if rabbitMQURL == "" {
        panic("RABBITMQ_URL is not set in environment")
    }

    conn, err := amqp.Dial(rabbitMQURL)
    if err != nil {
        panic(fmt.Sprintf("Failed to connect to RabbitMQ: %s", err))
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        panic(fmt.Sprintf("Failed to open a channel: %s", err))
    }
    defer ch.Close()

    q, err := ch.QueueDeclare(
        "sensor_data_queue",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        panic(fmt.Sprintf("Failed to declare a queue: %s", err))
    }

    rand.Seed(time.Now().UnixNano())

    for {
        value := rand.Float64() * 100
        alertLevel := "none"
        if value >= 90 {
            alertLevel = "high"
        }

        body := fmt.Sprintf("%f:%s", value, alertLevel)
        err = ch.Publish(
            "",
            q.Name,
            false,
            false,
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(body),
            },
        )
        if err != nil {
            panic(fmt.Sprintf("Failed to publish a message: %s", err))
        }

        fmt.Printf("Generated data: %f, Alert Level: %s\n", value, alertLevel)
        time.Sleep(5 * time.Second)
    }
}
