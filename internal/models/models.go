package models

import "time"

type Article struct {
	ID               string    `json:"id" bson:"_id,omitempty"`
	Source           string    `json:"source" bson:"source"`
	Title            string    `json:"title" bson:"title"`
	URL              string    `json:"url" bson:"url"`
	ImageURL         string    `json:"-" bson:"image_url,omitempty"`
	ImageBytes       []byte    `json:"-" bson:"image_bytes,omitempty"`
	ImageContentType string    `json:"-" bson:"image_content_type,omitempty"`
	ImageSrc         string    `json:"image_src" bson:"-"`
	Author           string    `json:"author,omitempty" bson:"author,omitempty"`
	PublishedAt      time.Time `json:"published_at,omitempty" bson:"published_at,omitempty"`
	Summary          string    `json:"summary" bson:"summary"`
	ReadTime         string    `json:"read_time" bson:"read_time"`
	Tags             []string  `json:"tags" bson:"tags"`
	FetchedAt        time.Time `json:"fetched_at,omitempty" bson:"fetched_at"`
	ContentHash      string    `json:"-" bson:"content_hash"`
}

type Look struct {
	ID               string    `json:"id" bson:"_id,omitempty"`
	Source           string    `json:"source" bson:"source"`
	Title            string    `json:"title" bson:"title"`
	ImageURL         string    `json:"-" bson:"image_url,omitempty"`
	ImageBytes       []byte    `json:"-" bson:"image_bytes,omitempty"`
	ImageContentType string    `json:"-" bson:"image_content_type,omitempty"`
	ImageSrc         string    `json:"image_src" bson:"-"`
	SourceURL        string    `json:"source_url" bson:"source_url"`
	Tags             []string  `json:"tags" bson:"tags"`
	Season           string    `json:"season,omitempty" bson:"season,omitempty"`
	FetchedAt        time.Time `json:"fetched_at,omitempty" bson:"fetched_at"`
	DisplayDate      string    `json:"display_date,omitempty" bson:"display_date,omitempty"`
	SelectedForDay   bool      `json:"selected_for_day,omitempty" bson:"selected_for_day"`
}

type Item struct {
	ID               string    `json:"id" bson:"_id,omitempty"`
	Source           string    `json:"source" bson:"source"`
	Brand            string    `json:"brand" bson:"brand"`
	Name             string    `json:"name" bson:"name"`
	Category         string    `json:"category" bson:"category"`
	Price            string    `json:"price" bson:"price"`
	Currency         string    `json:"currency,omitempty" bson:"currency,omitempty"`
	ImageURL         string    `json:"-" bson:"image_url,omitempty"`
	ImageBytes       []byte    `json:"-" bson:"image_bytes,omitempty"`
	ImageContentType string    `json:"-" bson:"image_content_type,omitempty"`
	ImageSrc         string    `json:"image_src" bson:"-"`
	ProductURL       string    `json:"product_url" bson:"product_url"`
	Tags             []string  `json:"tags" bson:"tags"`
	FetchedAt        time.Time `json:"fetched_at,omitempty" bson:"fetched_at"`
	DisplayDate      string    `json:"display_date,omitempty" bson:"display_date,omitempty"`
	SelectedForDay   bool      `json:"selected_for_day,omitempty" bson:"selected_for_day"`
}

type TrendSummary struct {
	ID        string    `json:"id" bson:"_id,omitempty"`
	Date      string    `json:"date" bson:"date"`
	Summary   string    `json:"summary" bson:"summary"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type DailyCategory struct {
	Date      string `json:"date"`
	Category  string `json:"category"`
	ItemCount int    `json:"item_count"`
}

type Image struct {
	ID          string
	URL         string
	Bytes       []byte
	ContentType string
}
