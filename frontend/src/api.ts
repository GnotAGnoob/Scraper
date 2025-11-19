import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

export interface MatchedProduct {
  ingredientName: string
  productName: string
  price: number
  unit: string
  imageUrl: string
  productUrl: string
  pricePerUnit: number
  pricePerUnitUnit: string
}

export interface RecipeResponse {
  products: MatchedProduct[]
}

export const processRecipe = async (text: string): Promise<RecipeResponse> => {
  const response = await axios.post<RecipeResponse>(`${API_BASE_URL}/api/v1/recipe`, { text })
  return response.data
}

