export interface ProductMatch {
  ingredient: string;
  quantityGrams: number;
  name: string;
  url: string;
  imageUrl: string;
  price: number;
  currency: string;
}

export interface RecipeResponse {
  matches: ProductMatch[];
}


