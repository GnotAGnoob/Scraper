import axios from "axios";
import type { IngredientsResponse, OrchestratorIngredientProducts } from "@/generated/kosik";

const API_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";

const apiClient = axios.create({
	baseURL: API_BASE_URL,
	headers: {
		"Content-Type": "application/json",
	},
});

export interface RecipeRequest {
	text: string;
}

export interface RecipeResponse {
	ingredientsProducts: OrchestratorIngredientProducts[];
}

export const recipeApi = {
	async parseRecipe(text: string): Promise<OrchestratorIngredientProducts[]> {
		const response = await apiClient.post<IngredientsResponse>("/api/v1/text/products", { text });
		return response.data.ingredientsProducts || [];
	},
};

export default apiClient;
