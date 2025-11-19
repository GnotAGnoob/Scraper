import type { MatchedProduct } from '../api'
import './ProductCard.css'

interface ProductCardProps {
  product: MatchedProduct
}

export default function ProductCard({ product }: ProductCardProps) {
  return (
    <div className="product-card">
      <div className="product-image-container">
        {product.imageUrl ? (
          <img 
            src={product.imageUrl} 
            alt={product.productName}
            className="product-image"
          />
        ) : (
          <div className="product-image-placeholder">No Image</div>
        )}
      </div>
      
      <div className="product-info">
        <div className="ingredient-tag">
          {product.ingredientName}
        </div>
        
        <h3 className="product-name">{product.productName}</h3>
        
        <div className="price-info">
          <div className="price-main">
            {product.price.toFixed(2)} Kč
          </div>
          <div className="price-unit">
            {product.pricePerUnit.toFixed(2)} Kč/{product.pricePerUnitUnit}
          </div>
        </div>
        
        {product.productUrl && (
          <a 
            href={product.productUrl} 
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

