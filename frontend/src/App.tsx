import { useState } from 'react'
import './App.css'
import { processRecipe, type MatchedProduct } from './api'
import ProductGrid from './components/ProductGrid'

function App() {
  const [recipeText, setRecipeText] = useState('')
  const [products, setProducts] = useState<MatchedProduct[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const handleSubmit = async () => {
    if (!recipeText.trim()) {
      setError('Please enter a recipe')
      return
    }

    setLoading(true)
    setError('')
    setProducts([])

    try {
      const result = await processRecipe(recipeText)
      setProducts(result.products)
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Failed to process recipe')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="app">
      <div className="container">
        <h1 className="title">Recipe to Kosik Cart</h1>
        
        <div className="input-section">
          <textarea
            className="recipe-input"
            placeholder="Paste your recipe here..."
            value={recipeText}
            onChange={(e) => setRecipeText(e.target.value)}
            rows={10}
          />
          
          <button 
            className="submit-btn" 
            onClick={handleSubmit}
            disabled={loading}
          >
            {loading ? 'Processing...' : 'Find Products'}
          </button>
        </div>

        {error && (
          <div className="error-box">
            {error}
          </div>
        )}

        {products.length > 0 && (
          <ProductGrid products={products} />
        )}
      </div>
    </div>
  )
}

export default App

