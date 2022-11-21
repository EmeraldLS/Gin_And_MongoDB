package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://lawrence:segun@cluster0.qvy6ing.mongodb.net/?retryWrites=true&w=majority"
const dbName = "StudentCollection"
const colName = "student"

var ctxTODO = context.TODO()
var ctxBackground = context.Background()

var Collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(ctxTODO, clientOptions)
	if err != nil {
		fmt.Printf("An error occured while connecting to the databse - %v\n", err)
	}
	fmt.Printf("Database with name: %v is successfully connected\n", dbName)
	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

}
