package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoStore struct {
	db *mongo.Database
}

func NewMongoStore(client *mongo.Client, database string) *MongoStore {
	if client == nil {
		return nil
	}
	return &MongoStore{db: client.Database(database)}
}

func (s *MongoStore) EnsureIndexes(ctx context.Context) error {
	if s == nil {
		return nil
	}
	indexes := map[string][]mongo.IndexModel{
		"articles": {
			{Keys: bson.D{{Key: "url", Value: 1}}, Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "published_at", Value: -1}}},
		},
		"looks": {
			{Keys: bson.D{{Key: "source_url", Value: 1}}, Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "display_date", Value: -1}, {Key: "selected_for_day", Value: 1}}},
		},
		"items": {
			{Keys: bson.D{{Key: "product_url", Value: 1}}, Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "display_date", Value: -1}, {Key: "selected_for_day", Value: 1}, {Key: "category", Value: 1}}},
		},
		"trend_summaries": {
			{Keys: bson.D{{Key: "date", Value: 1}}, Options: options.Index().SetUnique(true)},
			{Keys: bson.D{{Key: "created_at", Value: -1}}},
		},
	}
	for collection, models := range indexes {
		if _, err := s.db.Collection(collection).Indexes().CreateMany(ctx, models); err != nil {
			return err
		}
	}
	return nil
}

func (s *MongoStore) collection(name string) *mongo.Collection {
	return s.db.Collection(name)
}
