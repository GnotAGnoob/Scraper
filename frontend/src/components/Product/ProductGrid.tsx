import type { OrchestratorIngredientProducts } from "../../generated/kosik";
import { ProductCard } from "./ProductCard";
import { ProductEmpty } from "./ProductEmpty";

interface ProductGridProps {
	results: OrchestratorIngredientProducts[];
}

export const ProductGrid = ({ results }: ProductGridProps) => {
	if (results.length === 0) {
		return <ProductEmpty />;
	}

	return (
		<div className="space-y-12">
			{results.map((item, index) => {
				const ingredientName = item.ingredient?.name || "Unknown Ingredient";
				const quantity = item.ingredient?.quantityGrams ? `${item.ingredient.quantityGrams}g` : "";
				const title = quantity ? `${ingredientName} (${quantity})` : ingredientName;

				return (
					<div key={`${item.ingredient?.name}-${index}`} className="space-y-4">
						<h3 className="inline-block border-2 border-border bg-chart-5 px-4 py-1 font-black text-2xl shadow-shadow">
							{title}
						</h3>

						{item.products && item.products.length > 0 ? (
							<div className="grid grid-cols-[repeat(auto-fill,minmax(12rem,1fr))] gap-4">
								{item.products.map((product, pIndex) => (
									<ProductCard
										key={product.link?.value || pIndex}
										product={product}
										ingredientName={ingredientName}
									/>
								))}
							</div>
						) : (
							<div className="rounded-lg border-3 border-border/50 border-dashed p-6 font-mono text-muted-foreground">
								No products found for this ingredient.
							</div>
						)}
					</div>
				);
			})}
		</div>
	);
};
