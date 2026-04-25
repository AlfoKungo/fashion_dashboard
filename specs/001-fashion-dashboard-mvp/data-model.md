# Data Model: Personal Men's Fashion Dashboard MVP

## Article

Represents a fashion editorial or news item displayed in Top Articles and the Articles view.

**Fields**:

- `id`: Stable internal identifier.
- `source`: Publisher/source label, required for display.
- `title`: Article headline, required.
- `url`: Canonical article URL, required and unique.
- `image_url`: Original image URL, optional but retained when known.
- `image_bytes`: Stored image content, optional.
- `image_content_type`: Media type for stored image content, required when `image_bytes` exists.
- `author`: Author name, optional.
- `published_at`: Source publication timestamp, optional.
- `summary`: Short summary for cards, required for display.
- `read_time`: Read-time label, required for display and may be estimated.
- `tags`: Zero or more display tags.
- `fetched_at`: Timestamp when collected, required.
- `content_hash`: Duplicate identifier derived from canonical content details.

**Validation Rules**:

- `url` must be unique.
- `title`, `source`, `summary`, `read_time`, and `fetched_at` are required for a complete display card.
- `image_content_type` must be present when stored image bytes are present.
- Records older than 7 days by `fetched_at` are eligible for cleanup.

**Relationships**:

- Used by Today view and Articles view.
- Image route serves `image_bytes` when available, otherwise redirects or falls back to `image_url`.

## Look

Represents a daily outfit inspiration card.

**Fields**:

- `id`: Stable internal identifier.
- `source`: Source label, required.
- `title`: Look title, required.
- `image_url`: Original image URL, optional but retained when known.
- `image_bytes`: Stored image content, optional.
- `image_content_type`: Media type for stored image content, required when `image_bytes` exists.
- `source_url`: Canonical source URL, required and unique.
- `tags`: Zero or more display tags.
- `season`: Season label, optional.
- `fetched_at`: Timestamp when collected, required.
- `display_date`: Date selected for display, optional until chosen.
- `selected_for_day`: Whether this look is part of the daily set.

**Validation Rules**:

- `source_url` must be unique.
- `title`, `source`, and `fetched_at` are required.
- Exactly 4 looks should have `selected_for_day = true` for a `display_date` when at least 4 valid looks are available.
- Records older than 7 days by `fetched_at` are eligible for cleanup.

**State Transitions**:

- `collected` -> `selected_for_day` when chosen for a date.
- `selected_for_day` -> `expired` when older than retention and not needed for current display.

## Item

Represents a product or wardrobe item shown in Daily Item Focus and the Items view.

**Fields**:

- `id`: Stable internal identifier.
- `source`: Source label, required.
- `brand`: Brand name, required for display.
- `name`: Product or item name, required.
- `category`: Men's fashion category, required.
- `price`: Display price, required for display when available from source.
- `currency`: Currency label, optional when price already includes it.
- `image_url`: Original image URL, optional but retained when known.
- `image_bytes`: Stored image content, optional.
- `image_content_type`: Media type for stored image content, required when `image_bytes` exists.
- `product_url`: Canonical product URL, required and unique.
- `tags`: Zero or more display tags.
- `fetched_at`: Timestamp when collected, required.
- `display_date`: Date selected for display, optional until chosen.
- `selected_for_day`: Whether this item is part of the daily focus set.

**Validation Rules**:

- `product_url` must be unique.
- `brand`, `name`, `category`, and `fetched_at` are required.
- Exactly 6 items should have `selected_for_day = true` for a `display_date` when at least 6 valid items are available for the daily category.
- Records older than 7 days by `fetched_at` are eligible for cleanup.

**State Transitions**:

- `collected` -> `selected_for_day` when chosen for the daily category.
- `selected_for_day` -> `expired` when older than retention and not needed for current display.

## Trend Summary

Represents a dated summary of fashion trends generated during the daily workflow.

**Fields**:

- `id`: Stable internal identifier.
- `date`: Display date for the summary, required and unique per day.
- `summary`: Summary text, required.
- `created_at`: Timestamp when generated, required.

**Validation Rules**:

- `date` should be unique.
- `summary` must not be empty.
- Records older than 7 days by `created_at` are eligible for cleanup.

## Daily Category

Represents the active item focus category for a date.

**Fields**:

- `date`: Display date.
- `category`: One value from the configured category rotation.
- `item_count`: Number of selected items for the date.

**Validation Rules**:

- Category must be one of: loafers, sneakers, linen shirts, sweaters, chinos, jackets, coats, shorts.
- Category selection must be deterministic for the same date.

## Query Patterns and Indexes

- Articles: unique by `url`; sorted by `published_at` and/or `fetched_at`.
- Looks: unique by `source_url`; filtered by `display_date` and `selected_for_day`.
- Items: unique by `product_url`; filtered by `display_date`, `selected_for_day`, and `category`.
- Trend summaries: unique by `date`; cleanup by `created_at`.
