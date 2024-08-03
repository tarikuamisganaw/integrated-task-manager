package data

import (
	"context"
	"log"
	"task-mongo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var TaskCollection *mongo.Collection

func ConnectDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	Client = client
	TaskCollection = client.Database("taskdb").Collection("tasks")
}

// CreateTask inserts a new task into the MongoDB collection
func CreateTask(task models.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return TaskCollection.InsertOne(ctx, task)
}

// GetTasks retrieves all tasks from the MongoDB collection
func GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := TaskCollection.Find(ctx, bson.M{})
	if err != nil {
		return tasks, err
	}

	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	return tasks, cursor.Err()
}

// GetTaskByID retrieves a task by its ID from the MongoDB collection
func GetTaskByID(id string) (models.Task, error) {
	var task models.Task

	objectID, err := primitive.ObjectIDFromHex(id)
	log.Println(objectID)
	if err != nil {
		return task, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = TaskCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task)
	return task, err
}

// UpdateTask updates an existing task in the MongoDB collection
func UpdateTask(id string, task models.Task) (*mongo.UpdateResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return TaskCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.D{
		{"$set", task},
	})
}

// DeleteTask deletes a task from the MongoDB collection by its ID
func DeleteTask(id string) (*mongo.DeleteResult, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return TaskCollection.DeleteOne(ctx, bson.M{"_id": objectID})
}
