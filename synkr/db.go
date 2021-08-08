package synkr

import (
	"log"
	"fmt"

	"github.com/mattleong/lynkr/lynkr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SynkrDB struct {
	client *mongo.Client
}

func NewDBClient() *SynkrDB {
	// @TODO Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb://localhost:27017"
	ctx, cancel := CreateContext()
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return &SynkrDB{client: client}
}

func (db *SynkrDB) disconnect() {
	ctx, cancel := CreateContext()
	defer cancel()
	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *SynkrDB) ping() {
	ctx, cancel := CreateContext()
	defer cancel()
	if err := db.client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database is alive!")
}

func (db *SynkrDB) Save(requestLynk *RequestLynk) (*lynkr.Lynk, error) {
	fmt.Printf("Saving new lynk: %s\n", requestLynk)
	collection := db.client.Database("testing").Collection("lynks")

	url := "/z/" + requestLynk.Id
	lynk := lynkr.CreateLynk(requestLynk.Id, url, requestLynk.Url)

	ctx, cancel := CreateContext()
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.D{
		{ "id", lynk.Id },
		{ "url", lynk.Url },
		{ "goUrl", lynk.GoUrl },
	})

	return lynk, err
}

func (db *SynkrDB) FindOne(id string) *lynkr.Lynk {
	collection := db.client.Database("testing").Collection("lynks")
	filter := bson.D{{"id", id}}
	ctx, cancel := CreateContext()
	defer cancel()
	result := lynkr.Lynk{}
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}

	return &result
}
