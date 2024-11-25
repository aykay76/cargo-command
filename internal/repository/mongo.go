package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewMongoRepository(uri, dbName, collectionName string) (*MongoRepository, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := client.Database(dbName)
	collection := database.Collection(collectionName)

	return &MongoRepository{
		client:     client,
		database:   database,
		collection: collection,
	}, nil
}

func (r *MongoRepository) Create(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return r.collection.InsertOne(ctx, document)
}

func (r *MongoRepository) Read(ctx context.Context, filter interface{}) (*mongo.SingleResult, error) {
	return r.collection.FindOne(ctx, filter), nil
}

func (r *MongoRepository) Update(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	return r.collection.UpdateOne(ctx, filter, bson.M{"$set": update})
}

func (r *MongoRepository) Delete(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return r.collection.DeleteOne(ctx, filter)
}

func (r *MongoRepository) Close() error {
	return r.client.Disconnect(context.TODO())
}
