import axios from "axios";
import type { RecipeResponse } from "./types";

const apiBase =
  import.meta.env.VITE_API_BASE_URL ?? "http://localhost:8080";

const client = axios.create({
  baseURL: apiBase
});

export async function analyzeRecipe(text: string): Promise<RecipeResponse> {
  const res = await client.post<RecipeResponse>("/api/v1/recipe", { text });
  return res.data;
}


