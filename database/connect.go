package database

import (
	"context"
	"dispenser-backend/models"
	"dispenser-backend/utils"
	"encoding/json"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var NavCollection = new(mongo.Collection)
var err error

func ConnectMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(utils.Config.DB.Host))
	if err != nil {
		log.Fatal("database connection failed", err)
	}
	NavCollection = client.Database("dispenser").Collection("navigation")
}

func NavigationInsert(n *models.Navigation) error {
	var bNav interface{}
	bn, err := json.Marshal(n)
	if err != nil {
		log.Print("database.NavigationInsert: ", err)
		return err
	}
	err = bson.UnmarshalExtJSON(bn, true, &bNav)
	if err != nil {
		log.Print("database.NavigationInsert: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := NavCollection.InsertOne(ctx, bNav)
	log.Print(res)
	if err != nil {
		log.Print("database.NavigationInsert: ", err)
		return err
	}
	return nil
}

func NavigationGetByName(name string) (*models.Navigation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := new(models.Navigation)
	err := NavCollection.FindOne(ctx, bson.M{"name": name}).Decode(&result)
	if err != nil {
		log.Print("database.NavigationGet")
		return nil, err
	}
	return result, err
}
