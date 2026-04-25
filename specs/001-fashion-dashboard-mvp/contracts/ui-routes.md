# UI Route Contract: Personal Men's Fashion Dashboard MVP

## Shared Layout

All page routes render a dark-themed dashboard shell with:

- Logo text: `MEN'S FASHION DASHBOARD`
- Navigation links: Today, Looks, Items, Articles
- Current date in the top-right area
- Active navigation state for the current route

## Routes

### `GET /`

Renders the Today view.

Required sections:

- Top Articles with 4 cards by default
- Daily Inspiration with `4 LOOKS` indicator, 4 look cards, and View All Looks control
- Daily Item Focus with category label, `6 ITEMS` indicator, 6 item cards, and View All category control
- Footer with quote text and weather information

### `GET /looks`

Renders the Looks view.

Required behavior:

- Looks navigation item is active.
- The page displays look cards sourced from stored look records.
- Cards use the same core display fields as the Today view: image, title, and tags.

### `GET /items`

Renders the Items view.

Required behavior:

- Items navigation item is active.
- The page displays item cards sourced from stored item records.
- Cards use the same core display fields as the Today view: image, brand, name, and price.

### `GET /articles`

Renders the Articles view.

Required behavior:

- Articles navigation item is active.
- The page displays article cards sourced from stored article records.
- Cards use the same core display fields as the Today view: image, source, title, summary, read-time label, and tags.

## Empty and Partial States

- If fewer than the target number of records exists, render only complete cards.
- Do not render broken image elements or empty placeholder cards.
- Page shell and navigation must remain visible when a content section is empty.
