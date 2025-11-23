import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "./ui/card"
import { ExternalLink } from "lucide-react"

export interface Product {
  name: string
  price: number
  imageUrl: string
  productUrl: string
  ingredientName: string
  unit?: string
  quantity?: number
}

interface ProductCardProps {
  product: Product
}

export const ProductCard = ({ product }: ProductCardProps) => {
  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('cs-CZ', {
      style: 'currency',
      currency: 'CZK',
    }).format(price)
  }

  return (
    <Card className="group overflow-hidden">
      <CardHeader className="p-0">
        <div className="relative aspect-square overflow-hidden bg-foreground/5">
          <img
            src={product.imageUrl}
            alt={product.name}
            className="h-full w-full object-cover transition-transform group-hover:scale-105"
            loading="lazy"
          />
        </div>
      </CardHeader>
      <CardContent className="p-4">
        <div className="mb-2 inline-block rounded-[--radius-base] border-2 border-border bg-main px-3 py-1 text-sm font-heading text-main-foreground shadow-[2px_2px_0px_0px_var(--border)]">
          {product.ingredientName}
        </div>
        <CardTitle className="mb-2 line-clamp-2 text-lg">
          {product.name}
        </CardTitle>
        {product.quantity && product.unit && (
          <CardDescription className="mb-2">
            {product.quantity} {product.unit}
          </CardDescription>
        )}
      </CardContent>
      <CardFooter className="flex items-center justify-between border-t-4 border-border p-4">
        <div className="text-2xl font-heading">
          {formatPrice(product.price)}
        </div>
        <a
          href={product.productUrl}
          target="_blank"
          rel="noopener noreferrer"
          className="inline-flex items-center gap-2 rounded-[--radius-base] border-4 border-border bg-foreground px-4 py-2 font-heading text-sm text-background shadow-[var(--shadow)] transition-all hover:shadow-[6px_6px_0px_0px_var(--border)] hover:translate-x-[-2px] hover:translate-y-[-2px]"
        >
          View
          <ExternalLink className="h-4 w-4" />
        </a>
      </CardFooter>
    </Card>
  )
}


