package processing

import (
	"hash/fnv"
	"time"

	"fashion_dashboard/internal/models"
)

var Categories = []string{"loafers", "sneakers", "linen shirts", "sweaters", "chinos", "jackets", "coats", "shorts"}

func DailyCategory(date string) string {
	h := fnv.New32a()
	_, _ = h.Write([]byte(date))
	return Categories[int(h.Sum32())%len(Categories)]
}

func SelectLooksForDay(looks []models.Look, date string, amount int) []models.Look {
	out := make([]models.Look, 0, amount)
	for _, look := range looks {
		if !completeLook(look) {
			continue
		}
		look.DisplayDate = date
		look.SelectedForDay = true
		out = append(out, look)
		if len(out) == amount {
			break
		}
	}
	return out
}

func SelectItemsForDay(items []models.Item, date, category string, amount int) []models.Item {
	out := make([]models.Item, 0, amount)
	for _, item := range items {
		if !completeItem(item) || item.Category != category {
			continue
		}
		item.DisplayDate = date
		item.SelectedForDay = true
		out = append(out, item)
		if len(out) == amount {
			break
		}
	}
	if len(out) < amount {
		for _, item := range items {
			if !completeItem(item) || item.Category == category {
				continue
			}
			item.DisplayDate = date
			item.SelectedForDay = true
			out = append(out, item)
			if len(out) == amount {
				break
			}
		}
	}
	return out
}

func CleanupCutoff(now time.Time) time.Time {
	return now.AddDate(0, 0, -7)
}
