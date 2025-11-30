import { Beef, Candy, Carrot, Coins, Droplet, Droplets, Flame, Tags, Wheat } from "lucide-react";
import type { OrchestratorProductFields } from "../../generated/kosik";
import type { SortOption } from "../ui/sort-control";
import type { SortKey } from "./types";

export const MIN_ITEM_WIDTH = 12; // rem
export const MIN_ITEM_WIDTH_PX = MIN_ITEM_WIDTH * 16; // 12rem = 192px
export const GAP_SIZE = 1 * 16; // 1rem = 16px (gap-4)
export const ROWS_TO_SHOW = 2;

export const SORT_OPTIONS: SortOption<SortKey>[] = [
	{ label: "Price", value: "price", icon: Coins },
	{ label: "Price per Unit", value: "pricePerUnit", icon: Tags },
	{ label: "Calories", value: "calories", icon: Flame },
	{ label: "Protein", value: "protein", icon: Beef },
	{ label: "Fat", value: "fat", icon: Droplet },
	{ label: "Saturated Fat", value: "saturatedFat", icon: Droplets },
	{ label: "Carbs", value: "carbs", icon: Wheat },
	{ label: "Sugar", value: "sugar", icon: Candy },
	{ label: "Fiber", value: "fiber", icon: Carrot },
];

export const getProductValue = (product: OrchestratorProductFields, key: SortKey): number | undefined => {
	let value: number | undefined | null;

	switch (key) {
		case "price":
			value = product.price?.value;
			break;
		case "pricePerUnit":
			value = product.pricePerUnit?.value?.value;
			break;
		case "calories":
			value = product.nutrition?.value?.calories?.value;
			break;
		case "protein":
			value = product.nutrition?.value?.protein?.value;
			break;
		case "fat":
			value = product.nutrition?.value?.fat?.value;
			break;
		case "saturatedFat":
			value = product.nutrition?.value?.saturatedFat?.value;
			break;
		case "fiber":
			value = product.nutrition?.value?.fiber?.value;
			break;
		case "sugar":
			value = product.nutrition?.value?.sugar?.value;
			break;
		case "carbs":
			value = product.nutrition?.value?.carbs?.value;
			break;
	}

	if (value === null || value === undefined) {
		return undefined;
	}

	return value;
};
