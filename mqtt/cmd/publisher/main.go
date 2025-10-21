package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	broker   = "tcp://localhost:1883"
	topic    = "test/topic"
	clientID = "publisher"
)

// Callback for message publish completion
var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// Callback for successful connection
var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
}

// Callback for connection lost
var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func main() {
	// Configure MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	// Create MQTT client
	client := mqtt.NewClient(opts)

	// Connect to broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to broker: %v", token.Error())
	}

	// Publish messages
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Hello MQTT! Message #%d", i+1)

		token := client.Publish(topic, 0, false, message)
		if token.Wait() && token.Error() != nil {
			log.Printf("Failed to publish message: %v", token.Error())
		} else {
			fmt.Printf("Published: %s\n", message)
		}

		time.Sleep(2 * time.Second)
	}

	// Disconnect
	client.Disconnect(250)
	fmt.Println("Publisher disconnected")
}
