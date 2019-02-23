package project

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Repository interface {
	FindOne(id string) (GameProject, error)
	FindAll() ([]GameProject, error)
}

type ProjectRepository struct {
	client mongo.Client
}

func (ps *ProjectRepository) Create(model GameProject) (GameProject, error) {
	collection := ps.client.Database("projects").Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	res, err := collection.InsertOne(ctx, model)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	hexid := res.InsertedID.(primitive.ObjectID).Hex()
	docID, err := primitive.ObjectIDFromHex(hexid)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", docID}}

	var result GameProject
	err = collection.FindOne(ctx, filter).Decode(&result)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ps *ProjectRepository) FindOne(id string) (GameProject, error) {
	collection := ps.client.Database("projects").Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", docID}}

	var result GameProject
	err = collection.FindOne(ctx, filter).Decode(&result)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	return result, nil
}

func (ps *ProjectRepository) FindAll() ([]GameProject, error) {
	collection := ps.client.Database("projects").Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.M{})
	defer cancel()

	if err != nil {

	}

	var projects []GameProject

	for cur.Next(ctx) {
		var result GameProject
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, result)
		fmt.Println("data from a find all: ", result)
	}

	return projects, nil
}

func (ps *ProjectRepository) Update(id string, project GameProject) (GameProject, error) {
	collection := ps.client.Database("projects").Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)
	filter := bson.D{{"_id", docID}}

	if err != nil {
		log.Fatal(err)
	}

	lol := bson.D{
		{"$set", bson.D{
			{"name", project.Name},
			{"briefSynopsis", project.BriefSynopsis},
			{"fullDescription", project.FullDescription},
			{"contributors", project.Contributors},
			{"state", project.State},
		}},
	}

	_, err = collection.UpdateOne(ctx, filter, lol)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}

	var updated GameProject

	err = collection.FindOne(ctx, filter).Decode(&updated)
	defer cancel()

	if err != nil {
		log.Fatal(err)
	}
	return updated, nil
}

func (ps *ProjectRepository) Delete(id string) error {
	collection := ps.client.Database("projects").Collection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	objectIDS, err := primitive.ObjectIDFromHex(id)
	defer cancel()
	if err != nil {
		return fmt.Errorf("deleteTask: couldn't convert to-do ID from input: %v", err)
	}
	idDoc := bson.D{{"_id", objectIDS}}
	_, err = collection.DeleteOne(ctx, idDoc)
	defer cancel()
	if err != nil {
		return fmt.Errorf("deleteTask: couldn't delete: %v", err)
	}
	return nil
}

func NewRepository() *ProjectRepository {
	mclient, err := mongo.Connect(context.TODO(), "mongodb://mongodb:27017")

	if err != nil {
		fmt.Println("err mongo: ", err)
		log.Fatal(err)
	}

	err = mclient.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return &ProjectRepository{
		client: *mclient,
	}
}
