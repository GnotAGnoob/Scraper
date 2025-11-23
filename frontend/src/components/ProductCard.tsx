import { ExternalLink, ShoppingCart } from 'lucide-react'

import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import type { ProductFields } from '@/generated/kosik'
import { formatPrice } from '@/lib/format'

type ProductCardProps = {
  product: ProductFields
  ingredientName: string
}

const ProductCard = ({ product, ingredientName }: ProductCardProps) => {
  const imageUrl = product.imageUrl?.value ?? 'https://placehold.co/400x300?text=Kosik'
  const productName = product.name?.value ?? 'Unnamed product'
  const currency = product.currency?.value ?? 'CZK'
  const priceLabel = formatPrice(product.price?.value, currency)
  const pricePerUnitValue = product.pricePerUnit?.value
  const unitLabel = product.unit?.value
  const soldOut = product.isSoldOut?.value ?? false
  const link = product.link?.value ?? undefined

  return (
    <Card className="flex h-full flex-col border-[3px]">
      <CardHeader className="space-y-4 pb-0">
        <div className="flex items-start justify-between gap-3">
          <CardTitle className="text-xl">{productName}</CardTitle>
          {unitLabel && (
            <Badge variant="outline" className="text-[11px]">
              {unitLabel}
            </Badge>
          )}
        </div>
        <CardDescription>Matched for {ingredientName}</CardDescription>
      </CardHeader>
      <CardContent className="flex flex-1 flex-col gap-4">
        <div className="relative aspect-[4/3] overflow-hidden rounded-[var(--radius-base)] border-2 border-border bg-secondary-background">
          <img
            src={imageUrl}
            alt={productName}
            className="size-full object-cover"
            loading="lazy"
          />
          {soldOut && (
            <Badge className="absolute right-3 top-3 bg-red-500 text-white">Sold out</Badge>
          )}
        </div>

        <div>
          <p className="text-sm uppercase text-foreground/70">Price</p>
          <p className="text-2xl font-heading">{priceLabel}</p>
          {pricePerUnitValue && (
            <p className="text-sm text-foreground/70">
              {formatPrice(pricePerUnitValue.value, currency)} / {pricePerUnitValue.unit}
            </p>
          )}
        </div>

        <div className="mt-auto flex flex-wrap items-center gap-2">
          {link ? (
            <Button asChild variant="outline" size="sm">
              <a href={link} target="_blank" rel="noreferrer">
                <ShoppingCart className="size-4" />
                View on Kosik
                <ExternalLink className="size-4" />
              </a>
            </Button>
          ) : (
            <Badge variant="outline" className="text-[11px]">
              Link unavailable
            </Badge>
          )}
        </div>
      </CardContent>
    </Card>
  )
}

export default ProductCard


