package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	subscriberBroker   = "tcp://localhost:1883"
	subscriberTopic    = "test/topic"
	subscriberClientID = "subscriber"
)

// Callback for message reception
var subscriberMessageHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

// Callback for successful connection
var subscriberConnectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
}

// Callback for connection lost
var subscriberConnectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v\n", err)
}

func main() {
	// Configure MQTT client options
	opts := mqtt.NewClientOptions()
	opts.AddBroker(subscriberBroker)
	opts.SetClientID(subscriberClientID)
	opts.SetDefaultPublishHandler(subscriberMessageHandler)
	opts.OnConnect = subscriberConnectHandler
	opts.OnConnectionLost = subscriberConnectLostHandler

	// Create MQTT client
	client := mqtt.NewClient(opts)

	// Connect to broker
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to connect to broker: %v", token.Error())
	}

	// Subscribe to topic
	token := client.Subscribe(subscriberTopic, 1, subscriberMessageHandler)
	if token.Wait() && token.Error() != nil {
		log.Fatalf("Failed to subscribe to topic: %v", token.Error())
	}

	fmt.Printf("Subscribed to topic: %s\n", subscriberTopic)
	fmt.Println("Waiting for messages. Press Ctrl+C to exit.")

	// Wait for signal (exit with Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Unsubscribe and disconnect
	client.Unsubscribe(subscriberTopic)
	client.Disconnect(250)
	fmt.Println("Subscriber disconnected")
}
