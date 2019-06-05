package main

import (
	"testing"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
	"log"
)

func BenchmarkMqttVpsOne(b *testing.B){
	uri := "142.93.161.16:1883"

	topic := "testTimeTopic"
	listen(uri,topic)
	publisher := connect("pub",uri)
	
	for i:=0;i<b.N;i++{
		token := publisher.Publish(topic, 2, false, "hello")
		token.Wait()
	}
}

func listen(uri string, topic string){
	consumer := connect("sub",uri)
	consumer.Subscribe(topic, 2, func(client mqtt.Client, msg mqtt.Message){
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