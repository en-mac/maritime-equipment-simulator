package simulator

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    "github.com/streadway/amqp"
)

func Start() {
    rabbitMQURL := os.Getenv("RABBITMQ_URL")
    if rabbitMQURL == "" {
        panic("RABBITMQ_URL is not set in environment")
    }

    simulatorID := os.Getenv("SIMULATOR_ID")
    if simulatorID == "" {
        panic("SIMULATOR_ID is not set in environment")
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
        value := 30 + rand.Float64()*10  // Generates a value between 30 and 40
        alertLevel := "normal"
        if value < 33 {
            alertLevel = "low"
        } else if value > 37 {
            alertLevel = "high"
        }

        body := fmt.Sprintf("%f:%s:%s", value, alertLevel, simulatorID)
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

        fmt.Printf("Generated salinity data: %f, Alert Level: %s, Simulator ID: %s\n", value, alertLevel, simulatorID)
        time.Sleep(5 * time.Second)
    }
}
