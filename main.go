package main

import (
	"fmt"
	"context"
	"github.com/atotto/clipboard"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
func dbConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("connection error: ", err)
	}
	fmt.Println("Mongodb Successfully Connected")
	return client
}
func dbDisconnect(client *mongo.Client) {
	err := client.Disconnect(context.Background())
	if err != nil {
		fmt.Println("disconnect error: ", err)
	}
	fmt.Println("Mongodb Successfully Disconnected")
}
func main() {
	
	client := dbConnection()

	content , err := clipboard.ReadAll()
	fmt.Println("client: ", client)
	file, err := os.OpenFile("clipboard.txt" , os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if file != nil {
		file, err = os.Create("clipboard.txt")
	} 
	file.WriteString(content + "\n")
	file.Close()
	dbDisconnect(client)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf(content)
	}
}