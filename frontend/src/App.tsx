import { useState } from 'react'
import { useMutation } from '@tanstack/react-query'
import axios from 'axios'
import './App.css'

type Product = {
  Name: { Value: string }
  Price: { Value: number }
  ImageUrl: { Value: string }
  Link: { Value: string }
}

type IngredientResult = {
  ingredient_name: string
  quantity_grams: number
  product?: Product
  found: boolean
}

type RecipeResponse = {
  ingredients: IngredientResult[]
}

function App() {
  const [text, setText] = useState('')

  const mutation = useMutation({
    mutationFn: async (text: string) => {
      const res = await axios.post<RecipeResponse>('/api/v1/recipe', { text })
      return res.data
    },
  })

  return (
    <div className="container">
      <h1 className="title">Recipe 2 Kosik</h1>
      
      <div className="input-area">
        <textarea 
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Paste your recipe here..."
        />
        <button 
          onClick={() => mutation.mutate(text)}
          disabled={mutation.isPending || !text}
        >
          {mutation.isPending ? 'PROCESSING...' : 'CONVERT TO CART'}
        </button>
      </div>

      {mutation.isError && (
        <div style={{color: 'red', fontWeight: 'bold', border: '3px solid red', padding: '1rem', marginBottom: '2rem'}}>
          ERROR: {mutation.error.message}
        </div>
      )}

      {mutation.data && (
        <div className="grid">
          {mutation.data.ingredients.map((ing, idx) => (
            <div key={idx} className={`card ${!ing.found ? 'not-found' : ''}`}>
              <div style={{ borderBottom: '3px solid black', paddingBottom: '0.5rem', marginBottom: '1rem' }}>
                <strong>Ingredient:</strong> {ing.ingredient_name}
                <br/>
                <small>Est: {ing.quantity_grams}g</small>
              </div>
              
              {ing.found && ing.product ? (
                <>
                  <img src={ing.product.ImageUrl.Value} alt={ing.product.Name.Value} />
                  <h3>{ing.product.Name.Value}</h3>
                  <div className="price">{ing.product.Price.Value} Kč</div>
                  <a href={ing.product.Link.Value} target="_blank" rel="noreferrer" style={{marginTop: '1rem', display: 'block', fontWeight: 'bold', color: 'black'}}>
                    VIEW ON KOSIK &#8594;
                  </a>
                </>
              ) : (
                <div style={{ textAlign: 'center', padding: '2rem', fontStyle: 'italic' }}>
                  NO MATCH FOUND
                </div>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default App

