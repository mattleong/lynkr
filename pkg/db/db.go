package db

import (
	"context"
	"flag"
	"log"
	"time"

	l "github.com/mattleong/lynkr/pkg/lynk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func GetDBURI() *string {
	dbURI := flag.String("dbHost", "mongodb://localhost:27017", "URI to db host")
	return dbURI
}

func NewDBClient() DatabaseStore {
	uri := GetDBURI()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(*uri))
	if err != nil {
		log.Fatal(err)
	}

	db := &Database{
		client: client,
		collection: client.Database("lynks-test").Collection("lynks"),
		UplynkChan: make(chan *l.Lynk),
	}

	go db.startUplynk()

	var store DatabaseStore = db

	return store
}

func (db *Database) startUplynk() {
	for lynk := range db.UplynkChan {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := db.collection.InsertOne(ctx, bson.D{
			{ "lynkId", lynk.Id },
			{ "url", lynk.Url },
			{ "goUrl", lynk.GoUrl },
		})

		// @todo send to error channel?
		if (err != nil) {
			log.Println("Bad lynk -> ", err.Error())
			return
		}

		log.Println("Saved lynk -> ", lynk)
	}
}

func (db *Database) SaveLynk(url string) *l.Lynk {
	lynk := l.CreateLynk(url)
	log.Println("Saving lynk -> ", lynk)
	db.UplynkChan <- lynk

	return lynk
}

func (db *Database) FindLynkById(ctx context.Context, id string) (*l.Lynk, error) {
	var lynk l.Lynk
	filter := bson.D{{"lynkId", id}}

	err := db.collection.FindOne(ctx, filter).Decode(&lynk)
	if err == mongo.ErrNoDocuments {
		log.Println("Lynk does not exist")
	}

	return &lynk, err
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
