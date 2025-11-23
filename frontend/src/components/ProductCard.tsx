import { ExternalLink } from 'lucide-react'
import { Card, CardContent, CardFooter, CardHeader } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import type { ProductFields, Ingredient } from '@/lib/api'
import { cn } from '@/lib/utils'

interface ProductCardProps {
  product: ProductFields
  ingredient: Ingredient
}

const getFieldValue = <T,>(field: { value: T; error?: string }): T | null => {
  if (field.error) return null
  return field.value
}

const formatPrice = (price: number | null, currency: string | null): string => {
  if (price === null) return 'N/A'
  return `${price.toFixed(2)} ${currency || 'CZK'}`
}

export const ProductCard = ({ product, ingredient }: ProductCardProps) => {
  const name = getFieldValue(product.name)
  const price = getFieldValue(product.price)
  const currency = getFieldValue(product.currency)
  const unit = getFieldValue(product.unit)
  const imageUrl = getFieldValue(product.imageUrl)
  const link = getFieldValue(product.link)
  const isSoldOut = getFieldValue(product.isSoldOut) ?? false
  const pricePerUnit = getFieldValue(product.pricePerUnit)

  const kosikBaseUrl = 'https://www.kosik.cz'
  const fullLink = link ? (link.startsWith('http') ? link : `${kosikBaseUrl}${link}`) : null

  return (
    <Card className={cn(
      "transition-transform hover:scale-[1.02]",
      isSoldOut && "opacity-60"
    )}>
      <CardHeader className="p-0">
        {imageUrl ? (
          <div className="relative w-full aspect-square overflow-hidden rounded-t-[var(--radius-base)] border-b-2 border-border">
            <img
              src={imageUrl}
              alt={name || 'Product image'}
              className="w-full h-full object-cover"
              loading="lazy"
            />
            {isSoldOut && (
              <div className="absolute inset-0 bg-overlay flex items-center justify-center">
                <span className="bg-secondary-background border-2 border-border px-4 py-2 rounded-[var(--radius-base)] font-heading text-lg shadow-shadow">
                  SOLD OUT
                </span>
              </div>
            )}
          </div>
        ) : (
          <div className="w-full aspect-square bg-muted flex items-center justify-center border-b-2 border-border">
            <span className="text-muted-foreground font-base">No Image</span>
          </div>
        )}
      </CardHeader>
      <CardContent className="p-4">
        <div className="space-y-2">
          <div>
            <p className="text-xs text-muted-foreground font-base mb-1">
              Matched ingredient:
            </p>
            <p className="text-sm font-heading text-main">
              {ingredient.name}
            </p>
            {ingredient.quantityGrams > 0 && (
              <p className="text-xs text-muted-foreground mt-1">
                {ingredient.quantityGrams}g needed
              </p>
            )}
          </div>
          <div className="border-t-2 border-border pt-2">
            <h3 className="font-heading text-lg leading-tight mb-2 line-clamp-2">
              {name || 'Unknown Product'}
            </h3>
            <p className="text-2xl font-heading text-main">
              {formatPrice(price, currency)}
            </p>
            {pricePerUnit && (
              <p className="text-xs text-muted-foreground mt-1">
                {pricePerUnit.value.toFixed(2)} {pricePerUnit.unit || 'CZK'} / {unit || 'unit'}
              </p>
            )}
          </div>
        </div>
      </CardContent>
      {fullLink && (
        <CardFooter className="p-4 pt-0">
          <Button
            asChild
            variant="default"
            className="w-full"
            disabled={isSoldOut}
          >
            <a
              href={fullLink}
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center justify-center gap-2"
            >
              View on Kosik
              <ExternalLink className="size-4" />
            </a>
          </Button>
        </CardFooter>
      )}
    </Card>
  )
}

