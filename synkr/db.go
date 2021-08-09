package synkr

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/mattleong/lynkr/lynkr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type SynkrDB struct {
	client *mongo.Client
}

func getDBURI() *string {
	dbURI := flag.String("dbHost", "mongodb://localhost:27017", "URI to db host")
	return dbURI
}

func newDBClient() *SynkrDB {
	uri := getDBURI()
	ctx, cancel := createContext()
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatal(err)
	}

	return &SynkrDB{client: client}
}

func (db *SynkrDB) Disconnect(ctx context.Context) {
	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *SynkrDB) Ping(ctx context.Context) {
	if err := db.client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database is alive!")
}

func (db *SynkrDB) getLynkCollection() *mongo.Collection {
	return db.client.Database("testing").Collection("lynks")
}

func (db *SynkrDB) saveLynk(ctx context.Context, requestLynk *RequestLynk) (*lynkr.Lynk, error) {
	collection := db.getLynkCollection()
	url := "/z/" + requestLynk.Id
	lynk := lynkr.CreateLynk(requestLynk.Id, url, requestLynk.Url)

	_, err := collection.InsertOne(ctx, bson.D{
		{ "id", lynk.Id },
		{ "url", lynk.Url },
		{ "goUrl", lynk.GoUrl },
	})

	return lynk, err
}

func (db *SynkrDB) findLynkById(ctx context.Context, id string) (*lynkr.Lynk, error) {
	var result lynkr.Lynk
	collection := db.getLynkCollection()
	filter := bson.D{{"id", id}}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Println("record does not exist")
	}

	return &result, err
}
