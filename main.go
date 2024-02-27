package main

import (
	"context"
	"log"

	"recipes/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client //The mongo.Client type is provided by the MongoDB Go driver and is used for interacting with a MongoDB database.

var recipesHandler *handlers.RecipesHandler

func init() {
	// Connect to MongoDB
	ctx = context.TODO()                                                                        // create ctx Using context.TODO() for initialization                                                                      // Using context.TODO() for initialization                                                                      // Using context.TODO() for initialization
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/db")) //connect to database named db

	// Ping MongoDB to verify the connection
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")

	collection := client.Database("db").Collection("recipes")
	recipesHandler = handlers.NewRecipesHandler(ctx, collection)

}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("recipes/:id", recipesHandler.DeleteRecipeHandler)
	router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	router.Run()
}
