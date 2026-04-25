package repository

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *MongoStore) ListItems(ctx context.Context, amount int) ([]models.Item, error) {
	return s.findItems(ctx, bson.D{}, amount)
}

func (s *MongoStore) ListDailyItems(ctx context.Context, date string, amount int) ([]models.Item, error) {
	return s.findItems(ctx, bson.D{{Key: "display_date", Value: date}, {Key: "selected_for_day", Value: true}}, amount)
}

func (s *MongoStore) findItems(ctx context.Context, filter bson.D, amount int) ([]models.Item, error) {
	cursor, err := s.collection("items").Find(ctx, filter, options.Find().SetLimit(int64(amount)).SetSort(bson.D{{Key: "fetched_at", Value: -1}}))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var items []models.Item
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}
	for i := range items {
		items[i].ImageSrc = "/images/items/" + items[i].ID
	}
	return items, nil
}

func (s *MongoStore) UpsertItems(ctx context.Context, items []models.Item) error {
	for _, item := range items {
		_, err := s.collection("items").UpdateOne(ctx, bson.D{{Key: "product_url", Value: item.ProductURL}}, bson.D{{Key: "$set", Value: item}}, options.UpdateOne().SetUpsert(true))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MongoStore) MarkDailyItems(ctx context.Context, date string, ids []string) error {
	_, err := s.collection("items").UpdateMany(ctx, bson.D{}, bson.D{{Key: "$set", Value: bson.D{{Key: "selected_for_day", Value: false}}}})
	if err != nil {
		return err
	}
	_, err = s.collection("items").UpdateMany(ctx, bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: ids}}}}, bson.D{{Key: "$set", Value: bson.D{{Key: "selected_for_day", Value: true}, {Key: "display_date", Value: date}}}})
	return err
}

func (s *MongoStore) DeleteItemsOlderThan(ctx context.Context, cutoff time.Time) error {
	_, err := s.collection("items").DeleteMany(ctx, bson.D{{Key: "fetched_at", Value: bson.D{{Key: "$lt", Value: cutoff}}}})
	return err
}
