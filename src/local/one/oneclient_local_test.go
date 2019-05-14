package main

import (
	"testing"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
	"log"
)

func BenchmarkMqttLocalOne(b *testing.B){
	uri := "localhost:1883"
	topic := "testTimeTopic"
	go listen(uri,topic)
	publisher := connect("pub",uri)
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		publisher.Publish(topic, 0, false, "hello")
	}
}

func listen(uri string, topic string){
	consumer := connect("sub",uri)
	consumer.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message){
//		log.Print("message:", string(msg.Payload()))
	})
}

func connect(clientId string, uri string) mqtt.Client{
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetClientID(clientId)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	
	for !token.WaitTimeout(1 * time.Second){
	
	}
	
	if err := token.Error();err != nil{
		log.Fatal(err)
	}
	return client
}