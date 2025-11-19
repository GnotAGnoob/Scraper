import { useState } from 'react'
import InputArea from './components/InputArea'
import SubmitButton from './components/SubmitButton'
import ProductGrid from './components/ProductGrid'
import { MatchedProduct } from './types'
import './App.css'

function App() {
  const [recipeText, setRecipeText] = useState('')
  const [products, setProducts] = useState<MatchedProduct[]>([])
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const handleSubmit = async () => {
    if (!recipeText.trim()) {
      setError('Please enter a recipe')
      return
    }

    setLoading(true)
    setError(null)
    setProducts([])

    try {
      const response = await fetch('/api/v1/recipe', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ text: recipeText }),
      })

      if (!response.ok) {
        const errorData = await response.json()
        throw new Error(errorData.error || 'Failed to process recipe')
      }

      const data = await response.json()
      setProducts(data.products || [])
    } catch (err) {
      setError(err instanceof Error ? err.message : 'An error occurred')
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="app">
      <div className="container">
        <h1 className="title">Recipe to Kosik Cart</h1>
        <p className="subtitle">Paste your recipe and we'll find the ingredients!</p>
        
        <InputArea
          value={recipeText}
          onChange={(e) => setRecipeText(e.target.value)}
          placeholder="Enter your recipe here..."
        />
        
        <SubmitButton onClick={handleSubmit} disabled={loading}>
          {loading ? 'Processing...' : 'Find Products'}
        </SubmitButton>

        {error && <div className="error">{error}</div>}

        {products.length > 0 && (
          <ProductGrid products={products} />
        )}
      </div>
    </div>
  )
}

export default App

