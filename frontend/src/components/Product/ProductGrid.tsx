import { Globe } from "lucide-react";
import { useState } from "react";
import type { SortDirection } from "../ui/sort-control";
import { SortControl } from "../ui/sort-control";
import { IngredientSection } from "./IngredientSection";
import { ProductEmpty } from "./ProductEmpty";
import type { ProductGridProps, Sort, SortKey } from "./types";
import { SORT_OPTIONS } from "./utils";

export const ProductGrid = ({ results }: ProductGridProps) => {
	const [globalSort, setGlobalSort] = useState<Sort>({ key: "price", direction: "asc" });

	if (results.length === 0) {
		return <ProductEmpty />;
	}

	const setGlobalSortKey = (key: SortKey) => {
		setGlobalSort((prev) => ({ ...prev, key }));
	};

	const setGlobalSortDirection = (direction: SortDirection) => {
		setGlobalSort((prev) => ({ ...prev, direction }));
	};

	return (
		<div className="space-y-6">
			<div className="flex justify-between border-border border-b-4 pb-2">
				<h2 className="font-black text-3xl uppercase">Results</h2>
				<div className="flex items-center gap-2">
					<div className="flex items-center gap-1 text-muted-foreground">
						<Globe className="size-4" />
						Global
					</div>
					<SortControl
						sortKey={globalSort.key}
						sortDirection={globalSort.direction}
						options={SORT_OPTIONS}
						onSortChange={setGlobalSortKey}
						onDirectionToggle={setGlobalSortDirection}
					/>
				</div>
			</div>

			<div className="space-y-12">
				{results.map((item, index) => {
					return (
						<IngredientSection
							key={`${item.ingredient?.name}-${index}`}
							item={item}
							globalSort={globalSort}
						/>
					);
				})}
			</div>
		</div>
	);
};
