package repository

import (
	"sort"

	"fashion_dashboard/internal/models"
)

func orderArticlesByFreshnessAndSource(articles []models.Article) []models.Article {
	sort.SliceStable(articles, func(i, j int) bool {
		left := articles[i].PublishedAt
		right := articles[j].PublishedAt
		if left.IsZero() {
			left = articles[i].FetchedAt
		}
		if right.IsZero() {
			right = articles[j].FetchedAt
		}
		return left.After(right)
	})
	return interleaveArticleSources(articles)
}

func interleaveArticleSources(articles []models.Article) []models.Article {
	queues := map[string][]models.Article{}
	var sources []string
	for _, article := range articles {
		if _, ok := queues[article.Source]; !ok {
			sources = append(sources, article.Source)
		}
		queues[article.Source] = append(queues[article.Source], article)
	}

	out := make([]models.Article, 0, len(articles))
	lastSource := ""
	for len(out) < len(articles) {
		best := -1
		for i, source := range sources {
			if len(queues[source]) == 0 || source == lastSource {
				continue
			}
			if best == -1 || fresher(queues[source][0], queues[sources[best]][0]) {
				best = i
			}
		}
		if best == -1 {
			for i, source := range sources {
				if len(queues[source]) > 0 {
					best = i
					break
				}
			}
		}
		source := sources[best]
		out = append(out, queues[source][0])
		queues[source] = queues[source][1:]
		lastSource = source
	}
	return out
}

func fresher(left, right models.Article) bool {
	leftTime := left.PublishedAt
	rightTime := right.PublishedAt
	if leftTime.IsZero() {
		leftTime = left.FetchedAt
	}
	if rightTime.IsZero() {
		rightTime = right.FetchedAt
	}
	return leftTime.After(rightTime)
}
