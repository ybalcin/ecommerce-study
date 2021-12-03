package mgo

import (
	"context"
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
		return nil, ThrowNilCollectionError()
	}

	return c.collection.InsertOne(ctx, document)
}

// FindOne finds document from collection
func (c *Collection) FindOne(ctx context.Context, filter interface{}, decode interface{}) error {
	if decode == nil {
		return ThrowNilCollectionError()
	}

	if err := c.collection.FindOne(ctx, filter).Decode(decode); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}

		return err
	}

	return nil
}

// Find finds one or many document
func (c *Collection) Find(ctx context.Context, filter interface{}, decode interface{}) error {
	if decode == nil {
		return ThrowDecodeModelIsNilError()
	}

	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return err
	}

	defer cursorClose(cursor, ctx)

	err = cursor.All(ctx, decode)
	if err != nil {
		return err
	}

	return nil
}

// UpdateOne updates one
func (c *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}) error {
	_, err := c.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func cursorClose(cursor *mongo.Cursor, ctx context.Context) {
	err := cursor.Close(ctx)
	if err != nil {
		panic(err)
	}
}
