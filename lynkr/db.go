package lynkr

import (
	"context"
	"flag"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	client *mongo.Client
	collection *mongo.Collection
}

func getDBURI() *string {
	dbURI := flag.String("dbHost", "mongodb://localhost:27017", "URI to db host")
	return dbURI
}

func newDBClient(uri *string) *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		client: client,
		collection: client.Database("lynks-test").Collection("lynks"),
	}
}

func (db *Database) Disconnect(ctx context.Context) {
	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *Database) Ping(ctx context.Context) {
	if err := db.client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Database is alive!")
}

func (db *Database) SaveLynk(ctx context.Context, requestLynk *RequestLynk) (*Lynk, error) {
	url := "/z/" + requestLynk.Id
	lynk := CreateLynk(requestLynk.Id, url, requestLynk.Url)

	_, err := db.collection.InsertOne(ctx, bson.D{
		{ "lynkId", lynk.Id },
		{ "url", lynk.Url },
		{ "goUrl", lynk.GoUrl },
	})

	return lynk, err
}

func (db *Database) FindLynkById(ctx context.Context, id string) (*Lynk, error) {
	var result Lynk
	filter := bson.D{{"lynkId", id}}

	err := db.collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Println("record does not exist")
	}

	return &result, err
}
