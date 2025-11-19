import { MatchedProduct } from '../types'
import ProductCard from './ProductCard'
import './ProductGrid.css'

interface ProductGridProps {
  products: MatchedProduct[]
}

function ProductGrid({ products }: ProductGridProps) {
  return (
    <div className="product-grid">
      {products.map((product, index) => (
        <ProductCard key={index} product={product} />
      ))}
    </div>
  )
}

export default ProductGrid

