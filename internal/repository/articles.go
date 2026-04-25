package repository

import (
	"context"
	"time"

	"fashion_dashboard/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *MongoStore) ListArticles(ctx context.Context, amount int) ([]models.Article, error) {
	fetchLimit := int64(amount * 4)
	if fetchLimit < 24 {
		fetchLimit = 24
	}
	opts := options.Find().SetLimit(fetchLimit).SetSort(bson.D{{Key: "published_at", Value: -1}, {Key: "fetched_at", Value: -1}})
	cursor, err := s.collection("articles").Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var articles []models.Article
	if err := cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	articles = orderArticlesByFreshnessAndSource(articles)
	if len(articles) > amount {
		articles = articles[:amount]
	}
	for i := range articles {
		articles[i].ImageSrc = "/images/articles/" + articles[i].ID
	}
	return articles, nil
}

func (s *MongoStore) UpsertArticles(ctx context.Context, articles []models.Article) error {
	for _, article := range articles {
		_, err := s.collection("articles").UpdateOne(ctx, bson.D{{Key: "url", Value: article.URL}}, bson.D{{Key: "$set", Value: article}}, options.UpdateOne().SetUpsert(true))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *MongoStore) DeleteArticlesOlderThan(ctx context.Context, cutoff time.Time) error {
	_, err := s.collection("articles").DeleteMany(ctx, bson.D{{Key: "fetched_at", Value: bson.D{{Key: "$lt", Value: cutoff}}}})
	return err
}
