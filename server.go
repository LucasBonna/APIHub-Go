package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	
	
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
	"config"
)

var mongoClient *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic("Error creating a MongoDB client: " + err.Error())
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	err = client.Connect(ctx)
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}
	
	mongoClient = client
	
	session := mongoClient.Database("go").Session
	err = config.CreateDefaultCollections(session)
	if err != nil {
		panic("Error creating default collections: " + err.Error())
	}
}

func main() {	
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API Online!",
		})
	})
	
	public := r.Group("/api/v1")
	
	public.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Users API Online!",
		})
	})
	
	r.Run(":8000") // listen and serve on localhost:8080
	
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			fmt.Println("Error disconnecting from MongoDB:", err)
		}
	}()
}