import { FormEvent, useMemo, useState } from 'react'
import { useMutation } from '@tanstack/react-query'
import './App.css'
import { RecipeService } from './api/services/RecipeService'
import type { model_IngredientMatch } from './api/models/model_IngredientMatch'
import { ApiError } from './api/core/ApiError'

const demoRecipe = `2 chicken breasts
1 tbsp olive oil
3 cloves garlic, minced
1 lemon (juice + zest)
200g cherry tomatoes
1 handful fresh basil
sea salt and pepper`

const statusCopy: Record<
  'idle' | 'loading' | 'success' | 'error',
  { title: string; detail: string }
> = {
  idle: {
    title: 'Paste a recipe',
    detail: 'Gemini will normalize the ingredients, then Kosik search spins up.',
  },
  loading: {
    title: 'Scouting Kosik…',
    detail: 'Gemini already parsed the text. We are scraping for matches.',
  },
  success: {
    title: 'Matches ready',
    detail: 'Review the cards, open the Kosik link, tweak quantities if needed.',
  },
  error: {
    title: 'Something snapped',
    detail: 'Usually an empty prompt or missing API base URL.',
  },
}

function App() {
  const [recipeText, setRecipeText] = useState('')
  const mutation = useMutation({
    mutationFn: (text: string) =>
      RecipeService.postApiV1Recipe({
        request: { text },
      }),
  })

  const matches = mutation.data?.matches ?? []
  const canSubmit = recipeText.trim().length > 0 && !mutation.isPending
  const errorMessage = useMemo(() => {
    if (!mutation.error) {
      return ''
    }
    if (mutation.error instanceof ApiError) {
      return mutation.error.body?.error ?? mutation.error.message
    }
    if (mutation.error instanceof Error) {
      return mutation.error.message
    }
    return 'Unexpected error'
  }, [mutation.error])

  const statusKey: keyof typeof statusCopy = mutation.isError
    ? 'error'
    : mutation.isSuccess
      ? 'success'
      : mutation.isPending
        ? 'loading'
        : 'idle'

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault()
    const trimmed = recipeText.trim()
    if (!trimmed) {
      return
    }
    mutation.mutate(trimmed)
  }

  const handleFillDemo = () => {
    setRecipeText(demoRecipe)
    mutation.reset()
  }

  const handleReset = () => {
    setRecipeText('')
    mutation.reset()
  }

  return (
    <div className="app-shell">
      <header className="hero">
        <p className="eyebrow">Gemini 1.5 Flash ➜ Kosik scraper</p>
        <h1>Recipe → Kosik cart</h1>
        <p className="lede">
          Drop any free-form recipe text. We extract structured ingredients with
          Gemini JSON mode, fire concurrent Kosik searches, and keep the first
          legit hit per ingredient.
        </p>
      </header>

      <form className="recipe-form" onSubmit={handleSubmit}>
        <label htmlFor="recipe-text" className="label">
          Recipe text
        </label>
        <textarea
          id="recipe-text"
          value={recipeText}
          onChange={(event) => setRecipeText(event.target.value)}
          placeholder="Paste an ingredient list, a blog recipe, or even a voice transcription."
          rows={9}
          className="recipe-textarea"
        />
        <div className="form-actions">
          <div className="button-row">
            <button
              type="submit"
              className="cta"
              disabled={!canSubmit}
              aria-live="polite"
            >
              {mutation.isPending ? 'Syncing…' : 'Send to Kosik'}
            </button>
            <button
              type="button"
              className="ghost"
              onClick={handleFillDemo}
              disabled={mutation.isPending}
            >
              Use demo recipe
            </button>
            <button
              type="button"
              className="ghost"
              onClick={handleReset}
              disabled={mutation.isPending && !recipeText}
            >
              Reset
            </button>
          </div>
          <span className="hint">
            Tip: add quantities (e.g. “250g butter”) for sharper Kosik hits.
          </span>
        </div>
      </form>

      <section className="result-panel">
        <div className={`status-pill status-${statusKey}`}>
          <h2>{statusCopy[statusKey].title}</h2>
          <p>{mutation.isError ? errorMessage : statusCopy[statusKey].detail}</p>
        </div>
        {mutation.isError && (
          <div className="error-card">
            <p>{errorMessage}</p>
          </div>
        )}
        <MatchGrid
          matches={matches}
          isLoading={mutation.isPending}
          hasSubmitted={mutation.isSuccess || mutation.isPending}
        />
      </section>
    </div>
  )
}

type MatchGridProps = {
  matches: model_IngredientMatch[]
  isLoading: boolean
  hasSubmitted: boolean
}

const MatchGrid = ({ matches, isLoading, hasSubmitted }: MatchGridProps) => {
  if (isLoading) {
    return <div className="grid-placeholder pulse">Looking up products…</div>
  }

  if (!matches.length) {
    return (
      <div className="grid-placeholder">
        {hasSubmitted
          ? 'Gemini returned no ingredients. Try pasting a longer recipe.'
          : 'Your matched products will land here.'}
      </div>
    )
  }

  return (
    <div className="product-grid">
      {matches.map((match, index) => (
        <ProductCard key={`${match.ingredient?.name}-${index}`} match={match} />
      ))}
    </div>
  )
}

const ProductCard = ({ match }: { match: model_IngredientMatch }) => {
  const { ingredient, product, error } = match
  const hasProduct = Boolean(product)

  return (
    <article className={`product-card ${hasProduct ? '' : 'product-card--miss'}`}>
      <header>
        <div>
          <p className="card-label">Ingredient</p>
          <h3>{ingredient?.name ?? 'Unknown ingredient'}</h3>
          {ingredient?.notes && <p className="notes">{ingredient.notes}</p>}
        </div>
        <span className={`pill ${hasProduct ? 'pill-match' : 'pill-miss'}`}>
          {hasProduct ? 'Matched' : 'No match'}
        </span>
      </header>

      {hasProduct ? (
        <>
          {product?.image_url && (
            <img
              src={product.image_url}
              alt={product.name ?? 'Product image'}
              className="product-image"
              loading="lazy"
            />
          )}
          <div className="product-details">
            <p className="product-name">{product?.name}</p>
            <p className="product-price">
              {typeof product?.price === 'number'
                ? `${product.price.toFixed(2)} Kč`
                : 'N/A'}
              {product?.unit && <span className="unit">{product.unit}</span>}
            </p>
            {typeof product?.price_per_unit?.value === 'number' && (
              <p className="ppu">
                {product.price_per_unit.value.toFixed(2)} Kč /
                {product.price_per_unit.unit ?? ''}
              </p>
            )}
            <div className="links-row">
              {product?.link && (
                <a
                  href={product.link}
                  className="ghost"
                  target="_blank"
                  rel="noreferrer"
                >
                  View on Kosik
                </a>
              )}
              {product?.is_sold_out && (
                <span className="soldout-pill">Sold out</span>
              )}
            </div>
          </div>
        </>
      ) : (
        <p className="no-hit">
          {error ?? 'Kosik returned zero hits. Try tweaking the ingredient label.'}
        </p>
      )}
    </article>
  )
}

export default App
