package processor

import (
    "database/sql"
    "log"
    "strings"
    "os"

    "github.com/streadway/amqp"
    _ "github.com/lib/pq"
)

func Start() {
    connStr := os.Getenv("POSTGRES_CONN_STR")
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
    defer db.Close()

    createTableQuery := `
    CREATE TABLE IF NOT EXISTS sensor_data (
        id SERIAL PRIMARY KEY,
        value FLOAT NOT NULL,
        alert_level VARCHAR(10) NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        simulator_id VARCHAR(10) NOT NULL
    );`
    _, err = db.Exec(createTableQuery)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }

    rabbitMQConn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ: %v", err)
    }
    defer rabbitMQConn.Close()

    ch, err := rabbitMQConn.Channel()
    if err != nil {
        log.Fatalf("Failed to open a channel: %v", err)
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
        log.Fatalf("Failed to declare a queue: %v", err)
    }

    msgs, err := ch.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Fatalf("Failed to register a consumer: %v", err)
    }

    forever := make(chan bool)

    go func() {
        for d := range msgs {
            parts := strings.Split(string(d.Body), ":")
            value := parts[0]
            alertLevel := parts[1]
            simulatorID := parts[2]

            insertQuery := "INSERT INTO sensor_data (value, alert_level, simulator_id) VALUES ($1, $2, $3)"
            _, err = db.Exec(insertQuery, value, alertLevel, simulatorID)
            if err != nil {
                log.Printf("Failed to insert data: %v", err)
            } else {
                log.Printf("Received data: %s, Processed with alert level: %s, Simulator ID: %s", value, alertLevel, simulatorID)
            }
        }
    }()

    log.Printf("Waiting for messages. To exit press CTRL+C")
    <-forever
}
