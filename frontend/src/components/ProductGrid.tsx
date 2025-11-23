import { AlertTriangle, Box } from 'lucide-react'

import ProductCard from '@/components/ProductCard'
import { Badge } from '@/components/ui/badge'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import type { IngredientProducts } from '@/generated/kosik'
import { formatQuantity } from '@/lib/format'

type ProductGridProps = {
  groups?: IngredientProducts[]
  isLoading?: boolean
  error?: string | null
}

const ProductGrid = ({ groups, isLoading, error }: ProductGridProps) => {
  if (error) {
    return (
      <Card className="border-red-500">
        <CardHeader className="flex-row items-center justify-between gap-4">
          <div className="flex items-center gap-3">
            <AlertTriangle className="text-red-500" />
            <CardTitle className="text-red-500">Something went wrong</CardTitle>
          </div>
          <Badge variant="outline" className="text-red-500">
            retry above
          </Badge>
        </CardHeader>
        <CardContent>
          <p className="font-base text-red-600">{error}</p>
        </CardContent>
      </Card>
    )
  }

  if (isLoading) {
    return (
      <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
        {Array.from({ length: 6 }).map((_, index) => (
          <Skeleton key={index} className="h-[420px]" />
        ))}
      </div>
    )
  }

  if (!groups || groups.length === 0) {
    return (
      <div className="rounded-[var(--radius-base)] border-2 border-dashed border-border bg-secondary-background p-10 text-center shadow-[var(--shadow)]">
        <Box className="mx-auto mb-4 size-12" />
        <p className="text-2xl font-heading uppercase">Awaiting your recipe</p>
        <p className="text-sm text-foreground/70">
          Paste text above to let the LLM extract ingredients and Kosik match them.
        </p>
      </div>
    )
  }

  return (
    <div className="space-y-10">
      {groups.map((group) => (
        <section key={group.ingredient.name} className="space-y-4">
          <div className="flex flex-wrap items-center gap-3">
            <div>
              <p className="text-xs uppercase text-foreground/70">Ingredient</p>
              <p className="text-2xl font-heading uppercase">{group.ingredient.name}</p>
            </div>
            <Badge variant="outline">
              {formatQuantity(group.ingredient.quantityGrams)} • {group.source}
            </Badge>
          </div>
          {group.products.length === 0 ? (
            <Card className="border-dashed border-foreground/30">
              <CardHeader>
                <CardTitle className="text-lg">No products matched yet</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-foreground/70">
                  Try rephrasing the ingredient or expand the recipe above.
                </p>
              </CardContent>
            </Card>
          ) : (
            <div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
              {group.products.map((product, index) => (
                <ProductCard
                  key={`${group.ingredient.name}-${index}`}
                  product={product}
                  ingredientName={group.ingredient.name}
                />
              ))}
            </div>
          )}
        </section>
      ))}
    </div>
  )
}

export default ProductGrid


