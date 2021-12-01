package mgo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Collection is mongo collection
type Collection struct {
	collection *mongo.Collection
}

func ToObjectID(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

// Collection returns mongo collection
func (m *Store) Collection(collection string) *Collection {
	checkAgainstNil(m)

	return &Collection{m.db.Collection(collection)}
}

// InsertOne inserts new document to collection
func (c *Collection) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	if c == nil || c.collection == nil {
		panic("collection is nil")
	}

	return c.collection.InsertOne(ctx, document)
}

// FindOne finds document from collection
func (c *Collection) FindOne(ctx context.Context, filter interface{}, decode interface{}) error {
	if decode == nil {
		return fmt.Errorf("mgo: decode value is nil")
	}

	if err := c.collection.FindOne(ctx, filter).Decode(decode); err != nil {
		return err
	}

	return nil
}
