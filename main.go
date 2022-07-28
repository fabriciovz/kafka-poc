package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

const (
	Topic   = "my-topic"
	Brokers = "my-cluster-kafka-brokers:9093"
	Start   = 0
	End     = 1
)

func main() {
	item_updated()
}

func item_updated() {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{})

	w := &kafka.Writer{
		Addr:        kafka.TCP(Brokers),
		Topic:       Topic,
		Compression: kafka.Lz4,
		Transport: &kafka.Transport{
			TLS: TlsConfig(),
		},
		//Balancer: &kafka.Hash{},
		Logger: logger,
	}
	for i := Start; i < End; i++ {
		item := &ItemFileUpdated{
			UUID:       "11111",
			Timestamp:  "222222222",
			Issuer:     "333333333",
			StatusCode: "ok",
			ReplyTo:    "",
		}

		infItem, _ := json.Marshal(item)
		//fmt.Println(string(infItem))
		key := fmt.Sprintf("Key-%d", i)
		kafkaMessage := &kafka.Message{
			Key:   []byte(key),
			Value: infItem,
		}
		err := w.WriteMessages(context.Background(), *kafkaMessage)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
	}
	w.Close()
}

func TlsConfig() *tls.Config {
	caCert, err := ioutil.ReadFile(fmt.Sprintf("./%s", "ca.crt"))
	if err != nil {
		fmt.Println("Cagó el CA")
		return nil
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	// Keystore
	//key, err := os.ReadFile(fmt.Sprintf("/etc/secrets/%s", constant.CatalogKeyFile))
	key, err := os.ReadFile(fmt.Sprintf("./%s", "user.key"))
	if err != nil {
		log.Fatal("error reading kafka-catalog-client-key.pem file ", err.Error())
	}
	/*
		block, _ := pem.Decode(key)
		keyPassword := ""
		dkey, err := x509.DecryptPEMBlock(block, []byte(keyPassword)) //nolint
		if err != nil {
			log.Fatal("error decrypting key: ", err.Error())
		}
		pkey := pem.EncodeToMemory(&pem.Block{
			Type:  block.Type,
			Bytes: dkey,
		})*/
	//certPEM, err := os.ReadFile(fmt.Sprintf("/etc/secrets/%s", constant.CatalogCertFile))
	certPEM, err := os.ReadFile(fmt.Sprintf("./%s", "user.crt"))
	if err != nil {
		log.Fatal("error reading cert kafka-catalog-client-cert.pem ", err.Error())
	}
	cert, err := tls.X509KeyPair(certPEM, key)
	if err != nil {
		log.Fatal(err.Error())
	}
	config := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	return config
}
