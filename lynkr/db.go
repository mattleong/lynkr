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

type LynkrDB struct {
	client *mongo.Client
}

func getDBURI() *string {
	dbURI := flag.String("dbHost", "mongodb://localhost:27017", "URI to db host")
	return dbURI
}

func newDBClient() *LynkrDB {
	uri := getDBURI()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatal(err)
	}

	return &LynkrDB{client: client}
}

func (db *LynkrDB) Disconnect(ctx context.Context) {
	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *LynkrDB) Ping(ctx context.Context) {
	if err := db.client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Database is alive!")
}

func (db *LynkrDB) getLynkCollection() *mongo.Collection {
	return db.client.Database("testing").Collection("lynks")
}

func (db *LynkrDB) SaveLynk(ctx context.Context, requestLynk *RequestLynk) (*Lynk, error) {
	collection := db.getLynkCollection()
	url := "/z/" + requestLynk.Id
	lynk := CreateLynk(requestLynk.Id, url, requestLynk.Url)

	_, err := collection.InsertOne(ctx, bson.D{
		{ "id", lynk.Id },
		{ "url", lynk.Url },
		{ "goUrl", lynk.GoUrl },
	})

	return lynk, err
}

func (db *LynkrDB) FindLynkById(ctx context.Context, id string) (*Lynk, error) {
	var result Lynk
	collection := db.getLynkCollection()
	filter := bson.D{{"id", id}}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		log.Println("record does not exist")
	}

	return &result, err
}
