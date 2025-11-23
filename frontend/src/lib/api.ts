import axios from 'axios'
import type { Product } from '@/components/ProductCard'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export interface RecipeRequest {
  text: string
}

export interface RecipeResponse {
  products: Product[]
}

export const recipeApi = {
  async parseRecipe(text: string): Promise<Product[]> {
    const response = await apiClient.post<RecipeResponse>('/api/v1/recipe', { text })
    return response.data.products || []
  },
}

export default apiClient


