import type { MatchedProduct } from '../api'
import ProductCard from './ProductCard'
import './ProductGrid.css'

interface ProductGridProps {
  products: MatchedProduct[]
}

export default function ProductGrid({ products }: ProductGridProps) {
  return (
    <div className="products-section">
      <h2 className="products-title">Found Products ({products.length})</h2>
      <div className="product-grid">
        {products.map((product, index) => (
          <ProductCard key={index} product={product} />
        ))}
      </div>
    </div>
  )
}

