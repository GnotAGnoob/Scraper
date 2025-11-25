import axios from "axios";
import type { IngredientsRequest, IngredientsResponse, OrchestratorIngredientProducts } from "@/generated/kosik";
import { env } from "./env";

const client = axios.create({
	baseURL: env.apiBaseUrl,
	timeout: 15000,
});

export const postRecipe = async (payload: IngredientsRequest): Promise<OrchestratorIngredientProducts[]> => {
	const { data } = await client.post<IngredientsResponse>("/api/v1/text/products", payload);

	return data.ingredientsProducts ?? [];
};
