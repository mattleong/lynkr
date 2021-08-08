package synkr

import (
	"time"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
//    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/mattleong/lynkr/lynkr"
)


type SynkrClient struct {
	db *mongo.Client
}

func (s *SynkrClient) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.db.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Database is alive!")
}

func (s *SynkrClient) Save(requestLynk *lynkr.RequestLynk) (*lynkr.Lynk, error) {
//	collection := s.db.Database("testing").Collection("lynks")

	s.Ping()
	url := "/z/" + requestLynk.Id
	lynk := lynkr.Lynk{ Url: url }

//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	res, err := collection.InsertOne(ctx, bson.D{{"url", url}})
//	id := res.InsertedID

	fmt.Printf("Saving new lynk: %s", requestLynk)
//	fmt.Printf("inserted id: %s", id)

	return &lynk, nil
}

func NewSynkrClient(db *mongo.Client) *SynkrClient {
	return &SynkrClient{ db: db }
}

func GetClient() *mongo.Client {
	// @TODO Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb://localhost:27017"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")

	return client
}

