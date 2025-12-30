import { ChevronDown, ChevronUp, ExternalLink } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { AutoHeight } from "../ui/auto-height";
import { Button } from "../ui/button";
import { SectionBadge } from "../ui/section-badge";
import { SortControl } from "../ui/sort-control";
import { ProductCard } from "./ProductCard";
import type { IngredientSectionProps, Sort, SortKeyDefault } from "./types";
import { GAP_SIZE, getProductValue, MIN_ITEM_WIDTH, MIN_ITEM_WIDTH_PX, ROWS_TO_SHOW, SORT_OPTIONS } from "./utils";

export const IngredientSection = ({ item, globalSort }: IngredientSectionProps) => {
	const [expanded, setExpanded] = useState(false);
	const [localSort, setLocalSort] = useState<Sort | null>(null);

	const gridRef = useRef<HTMLDivElement>(null);
	const [maxVisibleItems, setMaxVisibleItems] = useState(0);

	const ingredientName = item.ingredient?.name || "Unknown Ingredient";
	const quantity = item.ingredient?.quantityGrams ? `${item.ingredient.quantityGrams}g` : "";
	const title = quantity ? `${ingredientName} (${quantity})` : ingredientName;
	const searchUrl = item.searchUrl;

	const products = item.products || [];

	const activeSort = localSort ?? globalSort;

	const sortedProducts = products.toSorted((a, b) => {
		const valA = getProductValue(a, activeSort.key);
		const valB = getProductValue(b, activeSort.key);

		// Handle undefined/missing values - always put them at the end
		// Using loose equality (== null) to catch both null and undefined
		if (valA == null && valB == null) return 0;
		if (valA == null) return 1;
		if (valB == null) return -1;

		if (valA === valB) return 0;

		const comparison = valA > valB ? 1 : -1;
		return activeSort.direction === "asc" ? comparison : -comparison;
	});

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

	const collapsedProducts = sortedProducts.slice(0, maxVisibleItems);

	const visibleProducts = expanded ? sortedProducts : collapsedProducts;
	const hiddenCount = Math.max(0, sortedProducts.length - maxVisibleItems);

	const handleToggleExpand = () => {
		setExpanded((prev) => !prev);
	};

	const globalOption = SORT_OPTIONS.find((o) => o.value === globalSort.key);
	const globalLabel = globalOption?.label || globalSort.key;
	const defaultLabel = `Default (${globalLabel})`;
	const defaultIcon = globalOption?.icon;

	const handleSortChange = (key: SortKeyDefault) => {
		if (key === "default") {
			setLocalSort(null);
		} else {
			setLocalSort({ key, direction: localSort?.direction || globalSort.direction });
		}
	};

	const handleDirectionToggle = () => {
		const key = localSort?.key || globalSort.key;
		const direction = localSort?.direction || globalSort.direction;
		setLocalSort({ key, direction: direction === "asc" ? "desc" : "asc" });
	};

	return (
		<div className="space-y-4">
			<div className="flex flex-wrap items-center justify-between gap-4">
				{/* // create a linkbutton that takes the style from the button. @frontend/src/components/Product/IngredientSection.tsx:99-109 then use the linkbutton instead of the sectionBadge */}
				<a
					href={searchUrl}
					target="_blank"
					rel="noopener noreferrer"
					className="inline-flex items-center gap-2 no-underline"
				>
					<SectionBadge color="chart-2" size="lg" className="flex items-center gap-2 py-2 leading-none">
						{title}
						<ExternalLink className="mt-px size-6" strokeWidth={3} />
					</SectionBadge>
				</a>

				<SortControl
					sortKey={localSort?.key ?? "default"}
					sortDirection={localSort?.direction || globalSort.direction}
					options={SORT_OPTIONS}
					showDefaultOption={true}
					defaultLabel={defaultLabel}
					defaultIcon={defaultIcon}
					onSortChange={handleSortChange}
					onDirectionToggle={handleDirectionToggle}
				/>
			</div>

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
								<ProductCard key={product.link?.value || pIndex} product={product} />
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
