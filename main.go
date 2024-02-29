package main

import (
	"context"
	"log"

	"recipes/handlers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ctx context.Context
var err error
var client *mongo.Client //The mongo.Client type is provided by the MongoDB Go driver and is used for interacting with a MongoDB database.

var recipesHandler *handlers.RecipesHandler

var authHandler *handlers.AuthHandler

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
	// authHandler = &handlers.AuthHandler{}

	collectionUsers := client.Database("db").Collection("users")
	authHandler = handlers.NewAuthHandler(ctx, collectionUsers)

}

func main() {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("recupes_api", store))
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.POST("/signin", authHandler.SignInHandler)
	router.POST("/refresh", authHandler.RefreshHandler)
	router.POST("/signout", authHandler.SignOutHandler)

	authorized := router.Group("/")
	authorized.Use(authHandler.AuthMiddleware())
	{
		authorized.POST("/recipes", recipesHandler.NewRecipeHandler)
		authorized.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
		authorized.DELETE("recipes/:id", recipesHandler.DeleteRecipeHandler)
		authorized.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	}
	router.Run()

}
