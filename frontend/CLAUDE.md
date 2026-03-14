# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

All commands run from the `frontend/` directory using Bun:

```bash
bun run dev           # Generate API client + start Vite dev server (port 3100)
bun run build         # Generate API client + production build + tsc type check
bun run test          # Run Vitest tests
bun run check         # Biome lint + format check combined
bun run lint          # Biome linter only
bun run format        # Biome formatter only
bun run gen:client    # Regenerate TypeScript client from ../docs/swagger.json
```

## Architecture

### Tech Stack

React 19, TypeScript, Vite, TanStack Router (file-based), TanStack Query, Tailwind CSS v4, Shadcn (https://www.neobrutalism.dev/) / Radix UI, Biome (linter + formatter).

### Key Conventions

- **Generated client** — `src/generated/kosik.ts` is auto-generated from `../docs/swagger.json` via `swagger-typescript-api`. Never edit it manually. Run `bun run gen:client` after any backend API change.
- **API client** — instantiated once in `src/lib/api.ts`, uses `VITE_API_URL` env var (defaults to `http://localhost:8080`).
- **Routing** — TanStack Router with file-based routes in `src/routes/`. `routeTree.gen.ts` is auto-generated — do not edit it.
- **UI components** — Shadcn (https://www.neobrutalism.dev/) components live in `src/components/ui/`. Add new Shadcn (https://www.neobrutalism.dev/) components there; do not modify generated Radix primitives directly.
- **Linting/formatting** — Use Biome, not ESLint or Prettier. Run `bun run check` before committing.

### Component Structure

```
src/
├── components/
│   ├── ui/              # Shadcn (https://www.neobrutalism.dev/) + custom primitives (Button, Card, etc.)
│   ├── Product/         # ProductCard, ProductGrid, IngredientSection, NutritionItem
│   ├── RecipeDisplay/   # Recipe text rendering
│   ├── Header/
│   ├── RecipeScraper.tsx  # Root feature component (input + results)
│   └── InputArea.tsx
├── routes/
│   ├── __root.tsx       # Root layout
│   └── index.tsx        # Single page — renders <RecipeScraper />
├── lib/
│   ├── api.ts           # API client singleton
│   └── utils.ts         # cn() helper (clsx + tailwind-merge)
└── generated/
    └── kosik.ts         # Auto-generated — do not edit
```

## Keeping Docs Updated

When making changes that affect the stack, component structure, conventions, or commands, update this file and the root `README.md`.
