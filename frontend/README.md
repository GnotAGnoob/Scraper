## Kosik Cart Pilot Frontend

Neobrutalist React + Vite surface that turns messy recipe text into a curated Kosik cart. Gemini 2.5 Flash extracts ingredients, we validate them, then scrape Kosik concurrently to display shoppable matches.

### Quick start

```bash
cd frontend
echo "VITE_API_URL=http://localhost:8080" > .env.local  # adjust if backend differs
bun install
bun run dev
```

The dev server automatically regenerates the API client before each run (`bun run gen:client`). Production builds will do the same.

### Scripts

- `bun run dev` — regenerate the swagger client and launch Vite on port 3000.
- `bun run build` — type-check, regenerate the client, and emit a production build.
- `bun run test` — Vitest suite (jsdom).
- `bun run lint|format|check` — Biome-powered consistency checks.
- `bun run gen:client` — manual swagger-typescript-api generation from `../docs/swagger.json`.

### Stack

- React 19 + TanStack Router, Query, Devtools.
- Tailwind CSS v4 with neobrutalism.dev palette + shadcn primitives (button, card, textarea, badge, skeleton).
- Axios API client with environment-driven base URL (`VITE_API_URL`, default `http://localhost:8080`).
- Biome for lint/format, Bun runtime.

### Data flow

1. User pastes recipe text into `InputArea`.
2. `submitRecipeText` mutation posts to `/api/v1/text/products`.
3. `ProductGrid` renders `IngredientProducts` groups with rich `ProductCard`s, including price per unit, stock state, and outbound links back to Kosik.
4. `StatusPanel` summarizes parsed ingredient counts vs surfaced products so the operator can see pipeline health at a glance.
