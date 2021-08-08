package synkr

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/mattleong/lynkr/lynkr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type SynkrClient struct {
	db *mongo.Client
	router *mux.Router
}

func CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func NewSynkrClient() *SynkrClient {
	db := NewDBClient()
	r := NewRouter()
	return &SynkrClient{ db: db, router: r }
}

func (s *SynkrClient) Disconnect() {
	ctx, cancel := CreateContext()
	defer cancel()
	err := s.db.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SynkrClient) Ping() {
	ctx, cancel := CreateContext()
	defer cancel()
	if err := s.db.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database is alive!")
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

