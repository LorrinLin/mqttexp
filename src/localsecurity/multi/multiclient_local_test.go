package main

import (
	"testing"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"time"
	"log"
	"strconv"
	"io/ioutil"
	tls "crypto/tls"
	x509 "crypto/x509"
)

func BenchmarkLocalMultiClients(b *testing.B){
	uri := "ssl://localhost:8883"
	topic := "testTimeTopic"
	publisher := connect("pub",uri)
	for i:=0;i<b.N;i++{
		listen(uri,topic,i)
		publisher.Publish(topic, 2, false, "hello")
	}
}

func listen(uri string, topic string, i int){
	consumer := connect("sub"+strconv.Itoa(i), uri)
	consumer.Subscribe(topic, 2, func(client mqtt.Client, msg mqtt.Message){
//		log.Print("message:", string(msg.Payload()))
	})
}

func connect(clientId string, uri string) mqtt.Client{
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetClientID(clientId)
	tlsconf := createTlsConf()
	opts.SetTLSConfig(tlsconf)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	
	for !token.WaitTimeout(1 * time.Second){
	
	}
	
	if err := token.Error();err != nil{
		log.Fatal(err)
	}
	return client
}

func createTlsConf() *tls.Config{
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile("ca.pem")
	//log.Println(pemCerts)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}
	
	cert,err := tls.LoadX509KeyPair("client-crt.pem", "client-key.pem")
	if err != nil{
		log.Println("err in load crt..",err)
	}
	//log.Println(cert)
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil{
		panic(err)
	}
	
	return &tls.Config{
		RootCAs:	certpool,
		ClientAuth:	tls.NoClientCert,
		ClientCAs:	nil,
		InsecureSkipVerify:	true,
		Certificates:	[]tls.Certificate{cert},
	}
	
}

