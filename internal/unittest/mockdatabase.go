package testing

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mockUserCollection struct {
	insertOneFn func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	findOneFn   func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	findFn      func(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
}

func (m *mockUserCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return m.insertOneFn(ctx, document, opts...)
}

func (m *mockUserCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return m.findOneFn(ctx, filter, opts...)
}

func (m *mockUserCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return m.findFn(ctx, filter, opts...)
}
