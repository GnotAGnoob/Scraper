export type MatchedProduct = {
  id: string
  name: string
  price: number
  currency: string
  imageUrl: string
  productUrl: string
  matchedIngredient: string
}

export type RecipeRequest = {
  text: string
}

export type RecipeResponse = MatchedProduct[]



