// Database
package database

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client   *mongo.Client
	database *mongo.Database
	context  *context.Context
}

// Single instance of the Mongo database
var instance Mongo
var once sync.Once

func debugLog(message string) {
	fmt.Printf("[ file: mongo.go ] %s\n", message)
}

func LoadMongo() *Mongo {
	return &instance
}

// Initialize single instance
func init() {
	debugLog("init method called")
	once.Do(func() {
		instance = Mongo{}

		var ctx context.Context = context.Background()

		// Mounting url to mongo database
		mongoURI := createMongoURI()

		var client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		if err != nil {
			log.Fatal(err)
		}
		debugLog("connected with Mondo DB")

		database := getDatabaseName()

		// Pass objects to instance
		instance.client = client
		instance.database = client.Database(database)
		instance.context = &ctx
		debugLog(fmt.Sprintf("%s: %s", "database accessed", database))

		// Get all existing collection in database
		var currentCollections, erro = instance.database.ListCollectionNames(*instance.context, bson.D{})
		if erro != nil {
			log.Fatal(erro)
		}
		if len(currentCollections) > 0 {
			debugLog(fmt.Sprintf("%s: %q", "database collection list", currentCollections))
		}
		for i := 0; i < len(schemaList); i++ {
			// Check if the collection already exists in the specific database, if it does not exist, it is to
			// create a database with validation forgetting regarding the collection.
			var collectionAlreadyExists bool = false
			for _, currentCollection := range currentCollections {
				if currentCollection == schemaList[i] {
					collectionAlreadyExists = true
					break
				}
			}

			if collectionAlreadyExists {
				continue
			}

			fmt.Println("[ file: mongo.go ] Setting validation schema")
			var options = options.CreateCollection().SetValidator(getSchema(schemaList[i]))
			fmt.Printf("[ file: mongo.go ] Creating collection: %s\n", schemaList[i])
			var err = instance.database.CreateCollection(*instance.context, schemaList[i], options)
			if err != nil {
				log.Fatal(err)
			}
		}
	})
}

func createMongoURI() string {
	// Getting environment variables
	username := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASS")
	hostname := os.Getenv("MONGO_HOST")
	if hostname == "" {
		hostname = "localhost"
	}

	// Mounting url to mongo database
	var mongoURI string
	if username == "" || password == "" {
		mongoURI = fmt.Sprintf("mongodb://%s", hostname)
	} else {
		mongoURI = fmt.Sprintf("mongodb+srv://%s:%s@%s", username, url.QueryEscape(password), hostname)
	}

	return mongoURI
}

func getDatabaseName() string {
	database := os.Getenv("MONGO_DB")
	if database == "" {
		database = "bytebank"
	}
	return database
}

// Disconnect from Mongo database
func (db *Mongo) Disconnect() {
	var err error = (*instance.client).Disconnect(*instance.context)
	if err == nil {
		log.Fatal(err)
	}
	(*instance.context).Done()
}
