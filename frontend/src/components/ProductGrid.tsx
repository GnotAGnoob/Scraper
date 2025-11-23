import type { MatchedProduct } from '@/types/recipe'

import { ProductCard } from './ProductCard'

type ProductGridProps = {
  products: MatchedProduct[]
  isLoading: boolean
  isError: boolean
}

export const ProductGrid = ({
  products,
  isLoading,
  isError,
}: ProductGridProps) => {
  if (isLoading) {
    return (
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        {Array.from({ length: 6 }).map((_, index) => (
          <div
            key={index}
            className="h-40 animate-pulse rounded-[var(--radius-base)] border-2 border-dashed border-border bg-secondary-background/70 shadow-shadow"
          />
        ))}
      </div>
    )
  }

  if (isError) {
    return (
      <div className="rounded-[var(--radius-base)] border-2 border-border bg-secondary-background px-4 py-3 text-sm font-base text-red-700 shadow-shadow dark:text-red-300">
        Something went wrong while talking to the server. Try again in a bit.
      </div>
    )
  }

  if (products.length === 0) {
    return (
      <div className="rounded-[var(--radius-base)] border-2 border-dashed border-border bg-secondary-background/60 px-4 py-6 text-sm font-base text-neutral-700 shadow-shadow dark:text-neutral-200">
        Matched products will appear here. Submit a recipe to see Kosik
        suggestions.
      </div>
    )
  }

  return (
    <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
      {products.map((product) => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  )
}



