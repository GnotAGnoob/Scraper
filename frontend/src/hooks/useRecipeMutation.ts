import { useMutation } from "@tanstack/react-query";
import { api } from "../lib/api";

export function useRecipeMutation() {
	return useMutation({
		mutationFn: (text: string) => api.text.v1TextProductsCreate({ text }),
	});
}

export type RecipeMutation = ReturnType<typeof useRecipeMutation>;
