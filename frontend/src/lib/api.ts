import type { IngredientsResponse, OrchestratorIngredientProducts, OrchestratorProductFields } from "@/generated/kosik";

// Failed scraper entries are serialized as null, which Swagger does not model.
export type IngredientProducts = Omit<OrchestratorIngredientProducts, "products"> & {
	products?: (OrchestratorProductFields | null)[] | null;
};
type RecipeResponse = Omit<IngredientsResponse, "ingredientsProducts"> & {
	ingredientsProducts?: IngredientProducts[] | null;
};

const API_BASE_URL = (import.meta.env.VITE_API_URL || "http://localhost:8080").replace(/\/$/, "");

export const recipeApi = {
	async parseRecipe(text: string): Promise<IngredientProducts[]> {
		const response = await fetch(`${API_BASE_URL}/api/v1/text/products`, {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify({ text }),
		});
		if (!response.ok) {
			throw new Error(`Failed to find products (HTTP ${response.status})`);
		}
		const data: RecipeResponse = await response.json();
		return data.ingredientsProducts ?? [];
	},
};
