package helpers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/EmeraldLS/Gin_And_MongoDB/database"
	"github.com/EmeraldLS/Gin_And_MongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddoneStudentToDB(student models.Student) {
	result, err := database.Collection.InsertOne(context.Background(), student)
	if err != nil {

		fmt.Printf("\nAn error occured while inserting %v into the database\n. Error detail is listed below.\n %v", student, err)
	} else {
		IndentedJSON, jsonErr := json.MarshalIndent(student, "", "  ")
		if jsonErr != nil {
			panic(jsonErr)
		}
		fmt.Printf("\nData \n%v\n has been inserted into the database successfully with insert id: %v", string(IndentedJSON), result.InsertedID)
	}
}

func RemoveOneStudentFromDB(studentID string) {
	ID, IdErr := primitive.ObjectIDFromHex(studentID)
	if IdErr != nil {
		fmt.Printf("\nAn error occured while converting ID to mongoDB ID\n %v\n", IdErr)
	}
	filter := bson.M{"_id": ID}
	result, err := database.Collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("\nAn error occured while deleting with Object ID of %v\n", result.DeletedCount)
	}
}

func RemoveAllStudentFromDB() {
	result, err := database.Collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("\nAn error occured while removing all the students\n %v\n", err)
	}
	fmt.Printf("\nAll student has been successfully deleted : %v\n", result.DeletedCount)
}

func GetAllStudentFromDB() []models.Student {
	cursor, err := database.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Printf("\nAn error occured while trying to get all student.\n%v\n", err)
	}
	var students []models.Student
	for cursor.Next(context.TODO()) {
		var student models.Student
		if err := cursor.Decode(&student); err != nil {
			fmt.Printf("\nAn error occured while decoding the student\n%v\n", err)
		}
		students = append(students, student)
	}
	defer cursor.Close(context.Background())
	return students
}

func GetSingleStudentFromDB(studentId string) models.Student {
	ID, IdErr := primitive.ObjectIDFromHex(studentId)
	if IdErr != nil {
		fmt.Printf("\nAn error occured while converting student Id to HEXFormat\n%verr\n", IdErr)
	}
	filter := bson.M{"_id": ID}
	var student models.Student
	err := database.Collection.FindOne(context.TODO(), filter).Decode(&student)
	if err != nil {
		fmt.Printf("\nAn error occured while finding student with ID: %v\n. Error Detail below\n %v\n", filter, err)
	}
	fmt.Println(student)
	return student
}
