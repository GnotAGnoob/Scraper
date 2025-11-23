import { useMutation } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { useCallback, useMemo, useState } from 'react'

import InputArea from '@/components/InputArea'
import ProductGrid from '@/components/ProductGrid'
import StatusPanel from '@/components/StatusPanel'
import { Badge } from '@/components/ui/badge'
import type { ExampleRecipe } from '@/lib/recipes'
import { exampleRecipes } from '@/lib/recipes'
import { extractErrorMessage, submitRecipeText } from '@/lib/api'

const CHAR_LIMIT = 2000

const IndexPage = () => {
  const [recipeText, setRecipeText] = useState(exampleRecipes[0]?.text ?? '')
  const [activeExampleId, setActiveExampleId] = useState<string | null>(exampleRecipes[0]?.id ?? null)

  const mutation = useMutation({
    mutationFn: submitRecipeText,
  })

  const stats = useMemo(() => {
    const ingredients = mutation.data?.ingredientsProducts ?? []
    const products = ingredients.reduce((total, group) => total + group.products.length, 0)
    const soldOut = ingredients.reduce(
      (total, group) =>
        total + group.products.filter((product) => product.isSoldOut?.value).length,
      0,
    )

    return {
      ingredients: ingredients.length,
      products,
      soldOut,
    }
  }, [mutation.data])

  const lastUpdated = useMemo(() => {
    if (!mutation.data) {
      return null
    }
    const now = new Date()
    return `Updated ${now.toLocaleTimeString('en-GB', {
      hour: '2-digit',
      minute: '2-digit',
    })}`
  }, [mutation.data])

  const errorMessage = mutation.error ? extractErrorMessage(mutation.error) : null

  const handleSubmit = useCallback(() => {
    if (!recipeText.trim()) {
      return
    }
    mutation.mutate({ text: recipeText.trim() })
  }, [mutation, recipeText])

  const handleExampleSelect = useCallback((example: ExampleRecipe) => {
    setRecipeText(example.text)
    setActiveExampleId(example.id)
  }, [])

  return (
    <main className="mx-auto flex max-w-6xl flex-col gap-12 px-4 py-10 md:py-16">
      <section className="space-y-4">
        <Badge variant="outline" className="w-fit">
          Beta neobrutal stack
        </Badge>
        <h1 className="max-w-3xl text-4xl font-heading uppercase leading-tight md:text-5xl">
          Paste chaotic recipe text, ship a curated Kosik cart in seconds.
        </h1>
        <p className="max-w-2xl text-base font-base text-foreground/80">
          Gemini 2.5 Flash extracts ingredients + quantities. We validate them, then fan out
          concurrent Kosik searches with price context so you can go from inspiration to checkout
          instantly.
        </p>
        <div className="flex flex-wrap gap-2 text-xs font-semibold uppercase text-foreground/70">
          <Badge variant="outline">Gemini 2.5 Flash</Badge>
          <Badge variant="outline">Concurrent Kosik scraper</Badge>
          <Badge variant="outline">LLM guardrails</Badge>
        </div>
      </section>

      <section className="grid gap-6 lg:grid-cols-[1.4fr,0.8fr]">
        <InputArea
          value={recipeText}
          onChange={(value) => {
            setRecipeText(value)
            setActiveExampleId(null)
          }}
          onSubmit={handleSubmit}
          isLoading={mutation.isPending}
          disabled={mutation.isPending}
          charLimit={CHAR_LIMIT}
          examples={exampleRecipes}
          activeExampleId={activeExampleId}
          onExampleSelect={handleExampleSelect}
        />
        <StatusPanel isLoading={mutation.isPending} stats={stats} lastUpdated={lastUpdated} />
      </section>

      <section className="space-y-6">
        <div className="flex flex-wrap items-center justify-between gap-3">
          <div>
            <p className="text-xs uppercase text-foreground/70">Ingredient → Product map</p>
            <h2 className="text-3xl font-heading uppercase">Latest Kosik matches</h2>
          </div>
          <Badge variant="outline">
            {stats.products} products • {stats.ingredients} ingredients
          </Badge>
        </div>

        <ProductGrid
          groups={mutation.data?.ingredientsProducts}
          isLoading={mutation.isPending}
          error={errorMessage}
        />
      </section>
    </main>
  )
}

export const Route = createFileRoute('/')({
  component: IndexPage,
})
