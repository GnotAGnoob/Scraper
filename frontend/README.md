# Recipe to Kosik Cart - Frontend

Neo-brutalist React + TypeScript + Vite frontend.

## Design System

### Colors
- Primary Yellow: `#ffd700`
- Pink: `#ff6b9d`
- Cyan: `#00ffff`
- Green: `#9aff9a`
- Black borders: `#000`
- White background: `#fff`

### Typography
- Font: `'Courier New', monospace`
- Bold weights (700-900)
- Uppercase text for emphasis

### Borders & Shadows
- Borders: `3-5px solid #000`
- Shadows: `4-8px hard drop shadows`
- No rounded corners
- No gradients

### Interactions
- Tactile button presses (shadow moves with transform)
- High contrast hover states
- No subtle transitions

## Components

### ProductCard
Displays a single product with:
- Product image
- Ingredient tag (pink)
- Product name and price
- Link to Kosik

### ProductGrid
Grid layout for product cards using CSS Grid.

### InputArea (App.tsx)
Large textarea for recipe input with:
- Bold borders
- Hard shadow
- Yellow background on focus

### SubmitButton (App.tsx)
Big cyan button with:
- Uppercase text
- Bold font
- Tactile press effect

## Development

```bash
npm install
npm run dev
```

Visit http://localhost:5173

## Build

```bash
npm run build
npm run preview
```

## API Integration

The frontend connects to the Go backend at `http://localhost:8080`.

API proxy is configured in `vite.config.ts` for development.

