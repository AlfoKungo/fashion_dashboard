# Personal Men’s Fashion Dashboard (MVP) — UI-Aligned Technical Specification

## Goal

Build a lightweight personal men’s fashion dashboard that matches the provided UI design.

The system fetches, processes, stores, and displays:
1. Fashion articles (Top Articles section)
2. Daily inspiration looks (Daily Inspiration section)
3. Daily item focus (Daily Item Focus section)

The UI is the source of truth for structure and behavior.

---

## Tech Stack

Backend:
- Language: Go
- HTTP: net/http (or Chi)
- Templates: Go HTML templates
- Database: MongoDB Atlas
- Scheduler: internal Go cron-like loop
- Config: environment variables

Frontend:
- Server-rendered HTML (Go templates)
- CSS (custom, matching design)
- No frontend framework

---

## Core Architecture

Single Go service responsible for:
- HTTP server
- MongoDB connection
- Scheduled workflows
- Fetching external data
- Data normalization
- Storage
- Rendering HTML UI

IMPORTANT:
- No client-side fetching from external APIs
- UI reads only from backend endpoints / server rendering

---

## UI Structure (Based on Design)

### Top Navigation Bar

Elements:
- Logo: "MEN’S FASHION DASHBOARD"
- Tabs:
  - TODAY (default active)
  - LOOKS
  - ITEMS
  - ARTICLES
- Current date (top-right)

Routing:
- / → TODAY view
- /looks
- /items
- /articles

---

## Main Dashboard Layout (TODAY)

### 1. Top Articles Section

Displays horizontally:
- 4 article cards (default)

Each card includes:
- Background image
- Source label (e.g., GQ, HIGHSNOBIETY)
- Title
- Short summary
- Read time (mock)
- Tag (e.g., Footwear, Trend Analysis)

Data source:
GET /api/articles?amount=4

---

### 2. Daily Inspiration Section

Header:
- "DAILY INSPIRATION"
- "4 LOOKS" indicator

Displays:
- 4 look cards

Each card includes:
- Image
- Look title (e.g., "Relaxed Neutrals")
- Tags (e.g., Casual, Summer)

CTA:
- "VIEW ALL LOOKS" → /looks

Data source:
GET /api/looks?amount=4

---

### 3. Daily Item Focus Section

Header:
- "DAILY ITEM FOCUS"
- Category label (e.g., "LOAFERS")
- "6 ITEMS" indicator

Displays:
- 6 item cards horizontally

Each item card includes:
- Product image
- Brand
- Name
- Price

CTA:
- "VIEW ALL {CATEGORY}" → /items

Data source:
GET /api/items?amount=6

---

### Footer

- Quote text
- Weather (mock/static)

---

## API Endpoints

GET /api/articles
- Query param: amount (default 4)

GET /api/looks
- Query param: amount (default 4)

GET /api/items
- Query param: amount (default 6)

Validation:
- amount must be integer >= 1
- max: 50

---

## MongoDB Collections

### articles

{
  _id: ObjectId,
  source: string,
  title: string,
  url: string,
  image_url: string,
  image_bytes: binary,
  image_content_type: string,
  author: string,
  published_at: datetime,
  summary: string,
  read_time: string,
  tags: string[],
  fetched_at: datetime,
  content_hash: string
}

Indexes:
- unique(url)
- published_at

---

### looks

{
  _id: ObjectId,
  source: string,
  title: string,
  image_url: string,
  image_bytes: binary,
  image_content_type: string,
  source_url: string,
  tags: string[],
  season: string,
  fetched_at: datetime,
  display_date: string,
  selected_for_day: bool
}

Indexes:
- unique(source_url)
- display_date

---

### items

{
  _id: ObjectId,
  source: string,
  brand: string,
  name: string,
  category: string,
  price: string,
  currency: string,
  image_url: string,
  image_bytes: binary,
  image_content_type: string,
  product_url: string,
  tags: string[],
  fetched_at: datetime,
  display_date: string,
  selected_for_day: bool
}

Indexes:
- unique(product_url)
- display_date

---

### trend_summaries

{
  _id: ObjectId,
  date: string,
  summary: string,
  created_at: datetime
}

---

## Image Handling

For ALL entities (articles, looks, items):

During fetch:
- Extract image_url
- Download image
- Store:
  - image_url
  - image_bytes
  - image_content_type

Failure handling:
- If download fails → keep URL only

---

## Image Serving

Routes:

GET /images/articles/{id}  
GET /images/looks/{id}  
GET /images/items/{id}

Behavior:
- If image_bytes exists:
  - Return bytes with Content-Type
- Else:
  - Redirect to image_url

HTML usage:

<img src="/images/articles/{{.ID}}">
<img src="/images/looks/{{.ID}}">
<img src="/images/items/{{.ID}}">

---

## Scheduled Daily Workflow

Runs once per day.

### Steps:

1. Fetch Articles
2. Fetch Looks
3. Fetch Items
4. Generate Trend Summary
5. Cleanup Old Data

---

## Sub-Workflow: Articles

- Fetch from sources
- Normalize
- Generate:
  - summary
  - tags
  - read_time (mock: "2 min read")
- Deduplicate by URL
- Store

---

## Sub-Workflow: Looks

- Detect season
- Fetch looks
- Normalize
- Store
- Select EXACTLY 4 looks:
  - selected_for_day = true
  - display_date = today

---

## Sub-Workflow: Items

- Determine daily category
- Fetch items
- Normalize
- Store
- Select EXACTLY 6 items:
  - selected_for_day = true
  - display_date = today

---

## Category Logic

Example:
index := hash(today) % len(categories)

categories:
- loafers
- sneakers
- linen shirts
- sweaters
- chinos
- jackets
- coats
- shorts

---

## Cleanup Policy

Run daily.

Delete records older than 7 days:

- articles → fetched_at
- looks → fetched_at
- items → fetched_at
- trend_summaries → created_at

DO NOT delete:
- sources

---

## Dashboard Rendering Logic

Server-side rendering using Go templates.

TODAY page:
- Load:
  - articles (limit 4)
  - looks (today, limit 4)
  - items (today, limit 6)
- Inject into template

---

## Styling Requirements (Based on Design)

- Dark theme background
- Card-based layout
- Rounded corners
- Image overlays for articles
- Horizontal scroll for sections
- Consistent spacing and typography
- Tag badges (rounded pills)
- CTA buttons

---

## Project Structure

fashion-dashboard/
  cmd/server/main.go
  internal/
    config/
    db/
    models/
    repository/
    fetchers/articles/
    fetchers/looks/
    fetchers/items/
    scheduler/
    processing/
    web/
      templates/
      static/

---

## Environment Variables

PORT=8080
MONGODB_URI=
MONGODB_DATABASE=fashion_dashboard
APP_ENV=development
DAILY_WORKFLOW_HOUR=7

---

## Acceptance Criteria

- UI matches provided design structure
- Articles section shows 4 cards
- Looks section shows 4 cards
- Items section shows 6 cards
- API supports dynamic amount
- Images render correctly (bytes or fallback)
- Daily workflow populates all sections
- Old data is deleted after 7 days
- No duplicate records
- Dashboard loads fully from MongoDB

---

## MVP Constraints

- No authentication
- No personalization
- No advanced scraping
- No React
- No real-time updates

---

## Deliverables

- Full Go backend
- MongoDB integration
- Scheduled workflows
- Image storage + serving
- UI matching design
- API endpoints
- README
- .env.example
