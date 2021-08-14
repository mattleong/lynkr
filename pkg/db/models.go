package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	l "github.com/mattleong/lynkr/pkg/lynk"
)

type DatabaseStore interface {
	SaveLynk(context.Context, string) (*l.Lynk, error)
	FindLynkById(context.Context, string) (*l.Lynk, error)
	Disconnect(context.Context)
	Ping(context.Context)
}

type Database struct {
	client *mongo.Client
	collection *mongo.Collection
}

