import { useMutation } from '@tanstack/react-query'
import { textProductsApi, type TextProductsResponse } from '@/lib/api'

export const useTextProducts = () => {
  return useMutation<TextProductsResponse, Error, string>({
    mutationFn: (text: string) => textProductsApi.processText(text),
  })
}

