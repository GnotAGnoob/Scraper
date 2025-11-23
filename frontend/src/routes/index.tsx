import { useCallback, useMemo, useState } from 'react'

import { useMutation } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

import { InputArea } from '@/components/InputArea'
import { ProductGrid } from '@/components/ProductGrid'
import { SubmitButton } from '@/components/SubmitButton'
import { postRecipe } from '@/lib/api'
import type { MatchedProduct } from '@/types/recipe'

const Home = () => {
  const [text, setText] = useState('')

  const mutation = useMutation({
    mutationKey: ['recipe', text],
    mutationFn: postRecipe,
  })

  const handleSubmit = useCallback(
    (event: React.FormEvent) => {
      event.preventDefault()

      const trimmed = text.trim()
      if (trimmed.length === 0 || mutation.isPending) {
        return
      }

      mutation.mutate({ text: trimmed })
    },
    [mutation, text],
  )

  const products: MatchedProduct[] = useMemo(
    () => mutation.data ?? [],
    [mutation.data],
  )

  return (
    <main className="mx-auto flex min-h-[calc(100vh-4rem)] max-w-6xl flex-col gap-8 px-4 py-8 md:flex-row">
      <section className="w-full md:w-1/2 space-y-4">
        <form onSubmit={handleSubmit} className="space-y-3">
          <InputArea
            value={text}
            onChange={setText}
            disabled={mutation.isPending}
          />
          <SubmitButton isLoading={mutation.isPending} />
        </form>
        <p className="text-xs text-neutral-700 dark:text-neutral-200">
          We send your raw recipe text to the backend, extract ingredients with
          Gemini, then search Kosik for matching products. No magic beyond
          HTTP.
        </p>
      </section>

      <section className="w-full md:w-1/2 space-y-3">
        <h2 className="text-xl font-heading">Matched products</h2>
        <ProductGrid
          products={products}
          isLoading={mutation.isPending}
          isError={mutation.isError}
        />
      </section>
    </main>
  )
}


export const Route = createFileRoute('/')({
  component: Home,
})

