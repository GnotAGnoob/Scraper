import { MatchedProduct } from '../types'
import './ProductCard.css'

interface ProductCardProps {
  product: MatchedProduct
}

function ProductCard({ product }: ProductCardProps) {
  return (
    <div className="product-card">
      {product.image_url && (
        <img
          src={product.image_url}
          alt={product.name}
          className="product-image"
        />
      )}
      <div className="product-content">
        <div className="ingredient-badge">{product.ingredient_name}</div>
        <h3 className="product-name">{product.name}</h3>
        <div className="product-price">
          {product.price.toFixed(2)} Kč / {product.unit}
        </div>
        {product.link && (
          <a
            href={product.link}
            target="_blank"
            rel="noopener noreferrer"
            className="product-link"
          >
            View on Kosik →
          </a>
        )}
      </div>
    </div>
  )
}

export default ProductCard

