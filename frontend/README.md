# Recipe to Kosik Cart - Frontend

A modern React frontend for converting recipes into Kosik shopping carts with a neobrutalist design aesthetic.

## Features

- 🎨 **Neobrutalist Design** - Bold borders, hard shadows, and high contrast
- 🌗 **Dark/Light Mode** - Theme toggle with localStorage persistence
- 🔍 **Recipe Parsing** - Extract ingredients from recipes using AI
- 🛒 **Product Matching** - Find matching products on Kosik.cz
- ⚡ **Modern Stack** - React 19, Vite, TanStack Query & Router
- 📱 **Responsive** - Works great on all devices

## Tech Stack

- **React 19** - Latest React with concurrent features
- **Vite** - Lightning-fast development and building
- **TanStack Router** - Type-safe routing
- **TanStack Query** - Powerful async state management
- **Tailwind CSS 4** - Utility-first styling with CSS variables
- **Bun** - Fast package manager and runtime
- **TypeScript** - Type safety throughout
- **Biome** - Fast linter and formatter

## Getting Started

### Prerequisites

- [Bun](https://bun.sh/) (v1.0.0 or higher)
- Backend server running on port 8080 (or configure VITE_API_URL)

### Installation

```bash
# From the project root
make frontend-install

# Or from the frontend directory
bun install
```

### Development

```bash
# From the project root
make frontend

# Or from the frontend directory
bun run dev
```

The app will be available at http://localhost:3000

### Environment Variables

Create a `.env` file in the frontend directory:

```env
VITE_API_URL=http://localhost:8080
```

### Build for Production

```bash
# From the project root
make frontend-build

# Or from the frontend directory
bun run build
```

## Project Structure

```
frontend/
├── src/
│   ├── components/
│   │   ├── ui/           # Reusable UI components (shadcn-style)
│   │   ├── Header.tsx    # App header with theme toggle
│   │   ├── ProductCard.tsx
│   │   └── ProductGrid.tsx
│   ├── lib/
│   │   ├── api.ts        # API client
│   │   └── utils.ts      # Utility functions
│   ├── routes/           # TanStack Router routes
│   │   ├── __root.tsx
│   │   └── index.tsx
│   ├── styles.css        # Global styles with neobrutalism theme
│   └── main.tsx          # App entry point
├── public/               # Static assets
└── vite.config.ts        # Vite configuration
```

## Design System

### Colors

The app uses a neobrutalist color scheme with OKLCH colors for better perceptual uniformity:

**Light Mode:**
- Background: Warm cream
- Foreground: Black
- Main: Soft yellow
- Border: Black with 4px shadows

**Dark Mode:**
- Background: Dark gray
- Foreground: Off-white
- Main: Muted yellow
- Border: Black with 4px shadows

### Typography

- **Font Family**: Space Grotesk (Google Fonts)
- **Base Weight**: 600 (semi-bold)
- **Heading Weight**: 800 (extra-bold)

### Components

All UI components follow the neobrutalism aesthetic:
- 4px black borders
- Box shadows (4px offset)
- Rounded corners (5px)
- Bold, chunky interactions
- High contrast colors

## Scripts

- `bun run dev` - Start development server
- `bun run build` - Build for production
- `bun run serve` - Preview production build
- `bun run gen:client` - Generate API client from Swagger
- `bun run lint` - Run Biome linter
- `bun run format` - Format code with Biome
- `bun run check` - Check code with Biome

## Adding Components

Use shadcn-style components:

```bash
bunx --bun shadcn@latest add <component> --yes
```

Then customize them to match the neobrutalism aesthetic.

## API Integration

The frontend communicates with the Go backend via REST API:

- `POST /api/v1/recipe` - Parse recipe and get matching products

See `src/lib/api.ts` for the API client implementation.

## Contributing

When contributing, please follow the existing code style:

- Use arrow functions: `const Component = () => {}`
- Prefer `useCallback` and `useMemo` where appropriate
- Avoid unnecessary `useEffect`
- Use semantic HTML
- Follow the neobrutalism design system

## License

Part of the Recipe to Kosik Cart project.
