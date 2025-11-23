import { useMemo } from 'react'
import { ProductCard } from './ProductCard'
import { Skeleton } from '@/components/ui/skeleton'
import type { IngredientProducts } from '@/lib/api'
import { cn } from '@/lib/utils'

interface ProductGridProps {
  data: IngredientProducts[] | undefined
  isLoading?: boolean
}

export const ProductGrid = ({ data, isLoading }: ProductGridProps) => {
  const totalProducts = useMemo(() => {
    if (!data) return 0
    return data.reduce((sum, item) => sum + item.products.length, 0)
  }, [data])

  if (isLoading) {
    return (
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {Array.from({ length: 8 }).map((_, i) => (
          <div key={i} className="space-y-4">
            <Skeleton className="w-full aspect-square" />
            <Skeleton className="h-4 w-3/4" />
            <Skeleton className="h-4 w-1/2" />
          </div>
        ))}
      </div>
    )
  }

  if (!data || totalProducts === 0) {
    return (
      <div className="text-center py-12">
        <p className="text-lg font-base text-muted-foreground">
          No products found. Try entering a recipe or ingredient list.
        </p>
      </div>
    )
  }

  return (
    <div className="space-y-8">
      {data.map((ingredientProducts, idx) => (
        <div key={idx} className="space-y-4">
          <div className="border-b-2 border-border pb-2">
            <h2 className="text-2xl font-heading">
              {ingredientProducts.ingredient.name}
            </h2>
            {ingredientProducts.ingredient.quantityGrams > 0 && (
              <p className="text-sm text-muted-foreground mt-1">
                {ingredientProducts.ingredient.quantityGrams}g needed
              </p>
            )}
            <p className="text-xs text-muted-foreground mt-1">
              Found {ingredientProducts.products.length} product{ingredientProducts.products.length !== 1 ? 's' : ''} from {ingredientProducts.source}
            </p>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {ingredientProducts.products.map((product, productIdx) => (
              <ProductCard
                key={productIdx}
                product={product}
                ingredient={ingredientProducts.ingredient}
              />
            ))}
          </div>
        </div>
      ))}
    </div>
  )
}

