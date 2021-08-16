package db

import (
	"context"

	l "github.com/mattleong/lynkr/pkg/lynk"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseStore interface {
	SaveLynk(context.Context, *l.Lynk) (UplynkResult)
	FindLynkById(context.Context, string) (*l.Lynk, error)
	Disconnect(context.Context)
	Ping(context.Context)
}

type Uplynk struct {
	ctx context.Context
	lynk *l.Lynk
}

type UplynkResult struct {
	Lynk *l.Lynk
	Err error
}

type Database struct {
	client *mongo.Client
	collection *mongo.Collection
	UplynkChan chan Uplynk
	UplynkResultChan chan UplynkResult
}

