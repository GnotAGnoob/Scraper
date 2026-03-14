# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

ScrapeMonkey is a recipe-to-products mapper: it accepts a recipe or ingredient list, uses an LLM (Google Gemini) to extract ingredients, scrapes Kosik.cz for matching products, and returns structured product data with prices and nutrition info.

## Commands

### Go Backend

```bash
make con              # Run interactive console scraper
make con-debug        # Run console with debug logging
make server           # Run REST API server (port 8080)
make server-debug     # Run server with debug logging
make swagger          # Regenerate Swagger docs (requires swag CLI)
make bench            # Run benchmarks (6 iterations)
make bench-stat       # Analyze benchmark results
```

### Frontend (Bun-based, in `frontend/`)

```bash
bun run dev           # Generate API client + start Vite dev server
bun run build         # Generate API client + production build + type check
bun run lint          # Biome linter
bun run format        # Biome formatter
bun run check         # Biome lint + format combined
bun run test          # Vitest tests
bun run gen:client    # Regenerate TypeScript API client from Swagger JSON
```

### Running a single Go test

```bash
go test ./internal/... -run TestName
go test ./internal/scraper/kosik/... -v
```

## Architecture

### Request Flow

```
POST /api/v1/text/products
    → Orchestrator.ProcessTextToProducts()
    → LLM extracts ingredient names from text
    → Concurrent scraping (max 10 workers) per ingredient against Kosik.cz
    → Product list fetch → per-product detail page fetch (nutrition data)
    → RecipeResult response
```

### Key Layers

- **`cmd/server/`** — REST API server (chi router). Single endpoint: `POST /api/v1/text/products`. Swagger docs at `/swagger/`.
- **`cmd/console/`** — Interactive CLI using table/progress bar output.
- **`internal/orchestrator/`** — Business logic coordinating LLM + scraper. Semaphore-limited concurrency.
- **`internal/scraper/kosik/`** — Kosik.cz HTTP scraper. Handles search URL construction (`urlParams/`), HTML parsing, product schema extraction, and retry logic.
- **`internal/llm/`** — LLM provider abstraction. Current implementation: Google Gemini (`google/`).
- **`frontend/`** — React 19 + TypeScript, TanStack Router (file-based routes in `src/routes/`), TanStack Query, Tailwind CSS v4, Shadcn/Radix UI. The TypeScript API client is auto-generated from Swagger — run `gen:client` after changing API types.

### Environment Variables

Required — copy `.env.example` to `.env`:

```
LLM_API_KEY=          # Google Gemini API key
LLM_MODEL=gemini-2.5-flash
LLM_PROVIDER=google
SERVER_PORT=8080
SERVER_ALLOWED_ORIGINS=http://localhost:3100
LOG_LEVEL=info
APP_ENV=development
```

## Keeping Docs Updated

When making changes that affect architecture, commands, environment variables, or key conventions, update this file (`CLAUDE.md`) and `README.md` to reflect those changes.

## Key Conventions

- **Scraper errors are partial** — the `Product` struct tracks per-field errors; a failed nutrition fetch doesn't abort the product.
- **Frontend API client is generated** — never hand-edit files under `src/generated/`. Run `bun run gen:client` after backend API changes.
- **Logging**: zerolog throughout the backend; use `logger.Logger` from `internal/logger/`.
- **Swagger**: annotations live in handler files. Regenerate with `make swagger` before committing API changes.
- **Biome** (not ESLint/Prettier) is the frontend linter and formatter.
