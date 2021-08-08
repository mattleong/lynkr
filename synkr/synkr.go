package synkr

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mattleong/lynkr/lynkr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


type SynkrClient struct {
	db *mongo.Client
}

func (s *SynkrClient) Ping() {
	ctx, cancel := CreateContext()
	defer cancel()
	if err := s.db.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Database is alive!")
}

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func (s *SynkrClient) FindOne(id string) *lynkr.Lynk {
	collection := s.db.Database("testing").Collection("lynks")
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

func (s *SynkrClient) Save(requestLynk *RequestLynk) (*lynkr.Lynk, error) {
	fmt.Printf("Saving new lynk: %s\n", requestLynk)
	collection := s.db.Database("testing").Collection("lynks")

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

func NewSynkrClient() *SynkrClient {
	db := GetClient()
	return &SynkrClient{ db: db }
}

func GetClient() *mongo.Client {
	// @TODO Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb://localhost:27017"
	ctx, cancel := CreateContext()
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
//	defer func() {
//		if err = client.Disconnect(ctx); err != nil {
//			panic(err)
//		}
//	}()
	return client
}

