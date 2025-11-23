import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080'

export const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

export interface FieldResult<T> {
  value: T
  error?: string
}

export interface NutritionFields {
  calories: FieldResult<number | null>
  protein: FieldResult<number | null>
  fat: FieldResult<number | null>
  saturatedFat: FieldResult<number | null>
  carbs: FieldResult<number | null>
  sugar: FieldResult<number | null>
  fiber: FieldResult<number | null>
  ingredients: FieldResult<string | null>
}

export interface ProductFields {
  name: FieldResult<string>
  price: FieldResult<number>
  currency: FieldResult<string>
  unit: FieldResult<string>
  pricePerUnit: FieldResult<{
    value: number
    unit: string
  }>
  link: FieldResult<string | null>
  imageUrl: FieldResult<string | null>
  isSoldOut: FieldResult<boolean>
  nutrition: FieldResult<NutritionFields | null>
}

export interface Ingredient {
  name: string
  quantityGrams: number
  originalText: string
}

export interface IngredientProducts {
  ingredient: Ingredient
  source: string
  products: ProductFields[]
}

export interface TextProductsRequest {
  text: string
}

export interface TextProductsResponse {
  ingredientsProducts: IngredientProducts[]
}

export const textProductsApi = {
  processText: async (text: string): Promise<TextProductsResponse> => {
    const response = await apiClient.post<TextProductsResponse>(
      '/api/v1/text/products',
      { text } as TextProductsRequest
    )
    return response.data
  },
}

