package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

func GetEsClient() *elastic.Client {
	// 启动客户端
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println(client)

	return client
}

func query() {
	
}