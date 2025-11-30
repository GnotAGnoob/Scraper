import type {
	OrchestratorIngredientProducts,
	OrchestratorNutritionFields,
	OrchestratorProductFields,
} from "../../generated/kosik";
import type { SortDirection } from "../ui/sort-control";

export type SortKey =
	| keyof Pick<OrchestratorProductFields, "price" | "pricePerUnit">
	| keyof Omit<OrchestratorNutritionFields, "ingredients">;

export type SortKeyDefault = SortKey | "default";

export interface Sort {
	key: SortKey;
	direction: SortDirection;
}

export interface ProductGridProps {
	results: OrchestratorIngredientProducts[];
}

export interface IngredientSectionProps {
	item: OrchestratorIngredientProducts;
	globalSort: Sort;
}
