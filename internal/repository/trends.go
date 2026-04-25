package repository

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *MongoStore) SaveTrendSummary(ctx context.Context, summary models.TrendSummary) error {
	_, err := s.collection("trend_summaries").UpdateOne(ctx, bson.D{{Key: "date", Value: summary.Date}}, bson.D{{Key: "$set", Value: summary}}, options.UpdateOne().SetUpsert(true))
	return err
}

func (s *MongoStore) DeleteTrendSummariesOlderThan(ctx context.Context, cutoff time.Time) error {
	_, err := s.collection("trend_summaries").DeleteMany(ctx, bson.D{{Key: "created_at", Value: bson.D{{Key: "$lt", Value: cutoff}}}})
	return err
}

func (s *MongoStore) GetArticleImage(ctx context.Context, id string) (models.Image, bool, error) {
	return s.getImage(ctx, "articles", id, "url")
}

func (s *MongoStore) GetLookImage(ctx context.Context, id string) (models.Image, bool, error) {
	return s.getImage(ctx, "looks", id, "source_url")
}

func (s *MongoStore) GetItemImage(ctx context.Context, id string) (models.Image, bool, error) {
	return s.getImage(ctx, "items", id, "product_url")
}

func (s *MongoStore) getImage(ctx context.Context, collection, id, urlField string) (models.Image, bool, error) {
	var doc struct {
		ID               string `bson:"_id"`
		ImageURL         string `bson:"image_url"`
		ImageBytes       []byte `bson:"image_bytes"`
		ImageContentType string `bson:"image_content_type"`
	}
	err := s.collection(collection).FindOne(ctx, bson.D{{Key: "_id", Value: id}}).Decode(&doc)
	if err != nil {
		return models.Image{}, false, nil
	}
	_ = urlField
	return models.Image{ID: id, URL: doc.ImageURL, Bytes: doc.ImageBytes, ContentType: doc.ImageContentType}, true, nil
}
