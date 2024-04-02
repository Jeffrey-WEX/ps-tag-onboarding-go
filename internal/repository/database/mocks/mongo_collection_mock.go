package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoCollectionMock struct {
	mock.Mock
}

func (m *IMongoCollectionMock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.Cursor), args.Error(1)
}

func (m *IMongoCollectionMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *IMongoCollectionMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}
