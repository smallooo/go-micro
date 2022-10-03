package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log-service/data"
	"time"
)

type MessageServer struct{}

type MessagePayLoad struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func (r *MessageServer) InsertMessage(message MessagePayLoad, resp *string) error {
	collection := client.Database("comments").Collection("comments")
	_, err := collection.InsertOne(context.TODO(), data.MessageEntry{
		ID:       message.ID,
		Message:  message.Message,
		CreateAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*resp = "Processed message via RPC:" + message.Message
	return nil
}

func (r *MessageServer) GetMessage(message MessagePayLoad, resp *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("comments").Collection("comments")

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})
	//
	//cursor, err := collection.Find(context.TODO(), bson.D{}, opts)

	cursor, err := collection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	defer cursor.Close(ctx)
	var messages []*data.MessageEntry

	for cursor.Next(ctx) {
		var item data.MessageEntry

		err := cursor.Decode(&item)
		if err != nil {
			log.Print("Error decoding message into slice:", err)
			return err
		} else {
			messages = append(messages, &item)
		}
	}

	*resp = "Processed message via RPC:" + message.Message + string(rune(len(messages))) + string("i") + "|||||" + string(messages[0].Message) + string(messages[0].ID) + string(messages[0].CreateAt.String())
	return nil
}
