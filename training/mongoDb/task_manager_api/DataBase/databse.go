package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskManager interface{
	Connect()
	CheckConnnection()
	ConnectTOCollection()
	Disconnect()

}
type Manager struct{

}

func (t *Manager)  Connect(url string) (*mongo.Client, error) {
	clientOption := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (t *Manager) CheckConnnection(client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func (t *Manager) ConnectTOCollection(client *mongo.Client, DatabaseName string, CollectionName string) *mongo.Collection {
	collection := client.Database(DatabaseName).Collection(CollectionName)
	fmt.Printf("sucessfully connected to DB:%v Collection:%v\n", DatabaseName, CollectionName)
	return collection
}

func (t *Manager) Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal("error connecting to the database")
	}
	fmt.Println("successfully disconnected  the database")

}
