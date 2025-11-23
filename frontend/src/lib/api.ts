import axios from 'axios'

import { env } from './env'
import type { RecipeRequest, RecipeResponse } from '@/types/recipe'

const client = axios.create({
  baseURL: env.apiBaseUrl,
  timeout: 15000,
})

export const postRecipe = async (
  payload: RecipeRequest,
): Promise<RecipeResponse> => {
  const { data } = await client.post<RecipeResponse>('/api/v1/text/products', payload)

  return data
}



