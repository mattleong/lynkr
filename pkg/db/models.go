package db

import (
	"context"

	l "github.com/mattleong/lynkr/pkg/lynk"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseStore interface {
	SaveLynk(*l.Lynk) (*l.Lynk)
	FindLynkById(context.Context, string) (*l.Lynk, error)
	Disconnect(context.Context)
	Ping(context.Context)
}

type Database struct {
	client *mongo.Client
	collection *mongo.Collection
	UplynkChan chan *l.Lynk
	DownlynkChan chan *l.Lynk
}

