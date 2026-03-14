import { useMutation } from "@tanstack/react-query";
import { api } from "../lib/api";

export function useSearchMutation() {
	return useMutation({
		mutationFn: (queries: string[]) => api.search.v1SearchProductsCreate({ queries }),
	});
}

export type SearchMutation = ReturnType<typeof useSearchMutation>;
