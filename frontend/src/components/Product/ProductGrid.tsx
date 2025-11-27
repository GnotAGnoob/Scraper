import { ChevronDown, ChevronUp } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import type { OrchestratorIngredientProducts } from "../../generated/kosik";
import { AutoHeight } from "../ui/auto-height";
import { Button } from "../ui/button";
import { ProductCard } from "./ProductCard";
import { ProductEmpty } from "./ProductEmpty";

interface ProductGridProps {
	results: OrchestratorIngredientProducts[];
}

const MIN_ITEM_WIDTH = 12; // rem
const MIN_ITEM_WIDTH_PX = MIN_ITEM_WIDTH * 16; // 12rem = 192px
const GAP_SIZE = 1 * 16; // 1rem = 16px (gap-4)
const ROWS_TO_SHOW = 2;

export const ProductGrid = ({ results }: ProductGridProps) => {
	if (results.length === 0) {
		return <ProductEmpty />;
	}

	return (
		<div className="space-y-12">
			{results.map((item, index) => {
				return <IngredientSection key={`${item.ingredient?.name}-${index}`} item={item} />;
			})}
		</div>
	);
};

interface IngredientSectionProps {
	item: OrchestratorIngredientProducts;
}

const IngredientSection = ({ item }: IngredientSectionProps) => {
	const [expanded, setExpanded] = useState(false);
	const gridRef = useRef<HTMLDivElement>(null);
	const [maxVisibleItems, setMaxVisibleItems] = useState(0);

	const ingredientName = item.ingredient?.name || "Unknown Ingredient";
	const quantity = item.ingredient?.quantityGrams ? `${item.ingredient.quantityGrams}g` : "";
	const title = quantity ? `${ingredientName} (${quantity})` : ingredientName;

	const products = item.products || [];

	useEffect(() => {
		const grid = gridRef.current;
		if (!grid) return;

		const calculateMaxVisible = () => {
			const containerWidth = grid.offsetWidth;
			if (containerWidth === 0) return;

			// Calculate how many items fit per row
			// Each item is at least MIN_ITEM_WIDTH_PX, with GAP_SIZE between items
			const itemsPerRow = Math.floor((containerWidth + GAP_SIZE) / (MIN_ITEM_WIDTH_PX + GAP_SIZE));
			const calculatedMax = Math.max(1, itemsPerRow * ROWS_TO_SHOW);
			setMaxVisibleItems(calculatedMax);
		};

		calculateMaxVisible();

		const resizeObserver = new ResizeObserver(calculateMaxVisible);
		resizeObserver.observe(grid);

		return () => {
			resizeObserver.disconnect();
		};
	}, []);

	const collapsedProducts = products.slice(0, maxVisibleItems);

	const visibleProducts = expanded ? products : collapsedProducts;
	const hiddenCount = Math.max(0, products.length - maxVisibleItems);

	const handleToggleExpand = () => {
		setExpanded((prev) => !prev);
	};

	return (
		<div className="space-y-4">
			<h3 className="inline-block border-2 border-border bg-chart-5 px-4 py-1 font-black text-2xl shadow-shadow">
				{title}
			</h3>

			{products.length > 0 ? (
				<>
					<AutoHeight>
						<div
							ref={gridRef}
							className="grid gap-4 p-2"
							style={{
								gridTemplateColumns: `repeat(auto-fill, minmax(${MIN_ITEM_WIDTH}rem, 1fr))`,
							}}
						>
							{visibleProducts.map((product, pIndex) => (
								<ProductCard
									key={product.link?.value || pIndex}
									product={product}
									ingredientName={ingredientName}
								/>
							))}
						</div>
					</AutoHeight>

					{products.length > maxVisibleItems && (
						<div className="flex items-center justify-center pt-2">
							<Button
								variant="neutral"
								size="sm"
								onClick={handleToggleExpand}
								className="text-sm"
								aria-label={expanded ? "Show fewer products" : `Show ${hiddenCount} more products`}
								aria-expanded={expanded}
							>
								{expanded ? (
									<>
										Hide products
										<ChevronUp className="size-5!" />
									</>
								) : (
									<>
										+{hiddenCount} more
										<ChevronDown className="size-5!" />
									</>
								)}
							</Button>
						</div>
					)}
				</>
			) : (
				<div className="rounded-lg border-3 border-border/50 border-dashed p-6 font-mono text-muted-foreground">
					No products found for this ingredient.
				</div>
			)}
		</div>
	);
};
