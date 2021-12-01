package mgo_test

import (
	"context"
	"github.com/ybalcin/ecommerce-study/pkg/mgo"
	"testing"
)

const (
	uri      = ""
	database = "database"
)

func TestNew(t *testing.T) {
	ctx := context.Background()

	store := mgo.NewStore(ctx, uri, database)
	if store == nil {
		t.Fail()
	}

	if !store.IsConnected() {
		t.Fail()
	}
}

func TestMgo_Collection(t *testing.T) {
	ctx := context.Background()

	store := mgo.NewStore(ctx, uri, database)

	collection := store.Collection("collection")
	if collection == nil {
		t.Fail()
	}
}
