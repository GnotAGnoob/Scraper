import axios, { AxiosError } from 'axios'

import type { TextToProductsPayload, TextToProductsResponse } from '@/generated/kosik'

const DEFAULT_BASE_URL = 'http://localhost:8080'

const normalizeBaseUrl = (url: string) => url.replace(/\/+$/, '')

const resolvedBase = import.meta.env.VITE_API_URL ?? DEFAULT_BASE_URL

if (!import.meta.env.VITE_API_URL) {
  console.warn('[Kosik] VITE_API_URL missing, defaulting to http://localhost:8080')
}

const baseURL = normalizeBaseUrl(resolvedBase)

export const apiClient = axios.create({
  baseURL,
  timeout: 45_000,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const submitRecipeText = async (payload: TextToProductsPayload) => {
  const { data } = await apiClient.post<TextToProductsResponse>(
    '/api/v1/text/products',
    payload,
  )
  return data
}

export const extractErrorMessage = (error: unknown) => {
  if (error instanceof AxiosError) {
    return error.response?.data?.message ?? error.message
  }

  if (error instanceof Error) {
    return error.message
  }

  return 'Unexpected error'
}

