package data

import (
	"context"
	"example/go_auth/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// // type UserAndTask interface {
// // 	AddDocumentToCollection()
// // 	UpdateDocumentInCollection()
// // 	DeleteDocumentInCollection()
// // 	ReadDocumentInCollection()
// // 	ReadAllDocumentInCollection()
// // }

func AddDocumentToCollection(collectionName string, document models.Document, client *mongo.Client) (bool, string) {
	collection := client.Database("JWT_Database").Collection(collectionName)
	_, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		return false, err.Error()
	}
	return true, "Sucessfully added the task"
}

func UpdateDocumentInCollection(collectionName string, id int, document models.Document, client *mongo.Client) (bool, string) {
	collection := client.Database("JWT_Database").Collection(collectionName)

	filter := bson.M{
		"id": id,
	}
	_, err := collection.UpdateOne(context.TODO(), filter, document)
	if err != nil {
		return false, err.Error()
	}
	return true, "update is sucessful"

}

// func DeleteDocumentInCollection(collectionName string, id int, client *mongo.Client) (bool, string) {
// 	collection := client.Database("JWT_Database").Collection(collectionName)

// 	filter := bson.M{
// 		"id": id,
// 	}
// 	_, err := collection.DeleteOne(context.TODO(), filter)
// 	if err != nil {
// 		return false, err.Error()
// 	}
// 	return true, "Delete is sucessful"

// }

// func ReadAllDocumentInCollection(collectionName string, client *mongo.Client) ([]interface{}, error) {
// 	collection := client.Database("JWT_Database").Collection(collectionName)
// 	findOption := options.Find()
// 	findOption.SetLimit(100)
// 	var result []interface{}
// 	cursor, err := collection.Find(context.TODO(), bson.D{}, findOption)
// 	if err != nil {
// 		return result, err
// 	}
// 	for cursor.Next(context.TODO()) {
// 		var doc bson.M
// 		err := cursor.Decode(&doc)
// 		if err != nil {
// 			return result, err
// 		}
// 		result = append(result, doc)
// 	}
// 	return result, nil

// }
