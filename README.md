# Recipe → Kosik Cart

LLM-assisted ingredient extraction plus the existing Kosik scraper. Paste any recipe, Gemini 1.5 Flash returns normalized ingredients, then we concurrently fetch the first valid Kosik hit per ingredient and surface it through a React frontend.

## Architecture

- **Go backend**
  - `internal/llm`: Google Gemini 1.5 Flash JSON-mode client (API key via env).
  - `internal/service`: recipe pipeline + concurrent Kosik fetcher (errgroup + semaphore).
  - `internal/api`: `POST /api/v1/recipe` handler behind chi, CORS, swagger docs served from `/swagger/index.html`.
- **React frontend (`frontend/`)**
  - Vite + TS + neo-brutalist UI.
  - OpenAPI client generated via `openapi-typescript-codegen` + Axios.
  - `@tanstack/react-query` to orchestrate POSTs and render product cards.

## Backend setup

```bash
export GOOGLE_GEMINI_API_KEY=your-key
export PORT=8080            # optional
export KOSIK_WORKERS=4      # optional cap for concurrent scrapes

make server                 # go run ./cmd/server
```

Swagger stays in `docs/`; regenerate after handler changes:

```bash
make swagger
```

## Frontend setup

```bash
cd frontend
cp .env.example .env
echo "VITE_API_BASE_URL=http://localhost:8080" >> .env # adjust per deploy
npm install
npm run generate   # re-build the OpenAPI client when swagger changes
npm run dev        # launches Vite dev server with neo-brutalist UI
```

The React app reads `VITE_API_BASE_URL`, so dev proxies straight to the Go server while prod can point at any hosted instance.

## Make targets worth remembering

```bash
make con           # console-based scraping runner
make con-debug     # verbose logging
make bench         # go test -bench with tee’d output
make bench-clean   # drop ./benchmark artifacts
make frontend-dev  # cd frontend && npm run dev
make frontend-init # cd frontend && npm install
```

## Notes

- Kosik scraping remains the bottleneck; tweak `KOSIK_WORKERS` if Gemini responses queue faster than HTTP allows.
- For Gemini costs, cache responses per recipe hash (easy future enhancement).
- Generated API client lives in `frontend/src/api`; don’t hand edit—use `npm run generate`.
