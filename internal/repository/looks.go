package repository

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *MongoStore) ListLooks(ctx context.Context, amount int) ([]models.Look, error) {
	return s.findLooks(ctx, bson.D{}, amount)
}

func (s *MongoStore) ListDailyLooks(ctx context.Context, date string, amount int) ([]models.Look, error) {
	return s.findLooks(ctx, bson.D{{Key: "display_date", Value: date}, {Key: "selected_for_day", Value: true}}, amount)
}

func (s *MongoStore) findLooks(ctx context.Context, filter bson.D, amount int) ([]models.Look, error) {
	cursor, err := s.collection("looks").Find(ctx, filter, options.Find().SetLimit(int64(amount)).SetSort(bson.D{{Key: "fetched_at", Value: -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var looks []models.Look
	if err := cursor.All(ctx, &looks); err != nil {
		return nil, err
	}
	for i := range looks {
		looks[i].ImageSrc = "/images/looks/" + looks[i].ID
	}
	return looks, nil
}

func (s *MongoStore) UpsertLooks(ctx context.Context, looks []models.Look) error {
	for _, look := range looks {
		_, err := s.collection("looks").UpdateOne(ctx, bson.D{{Key: "source_url", Value: look.SourceURL}}, bson.D{{Key: "$set", Value: look}}, options.UpdateOne().SetUpsert(true))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MongoStore) MarkDailyLooks(ctx context.Context, date string, ids []string) error {
	_, err := s.collection("looks").UpdateMany(ctx, bson.D{}, bson.D{{Key: "$set", Value: bson.D{{Key: "selected_for_day", Value: false}}}})
	if err != nil {
		return err
	}
	_, err = s.collection("looks").UpdateMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}, bson.D{{Key: "$set", Value: bson.D{{Key: "selected_for_day", Value: true}, {Key: "display_date", Value: date}}}})
	return err
}

func (s *MongoStore) DeleteLooksOlderThan(ctx context.Context, cutoff time.Time) error {
	_, err := s.collection("looks").DeleteMany(ctx, bson.D{{Key: "fetched_at", Value: bson.D{{Key: "$lt", Value: cutoff}}}})
	return err
}
