import type { OrchestratorIngredientProducts } from "../generated/kosik";
import { ProductCard } from "./ProductCard";

interface ProductGridProps {
	results: OrchestratorIngredientProducts[];
}

export const ProductGrid = ({ results }: ProductGridProps) => {
	if (!results || results.length === 0) {
		return null;
	}

	return (
		<div className="space-y-12">
			{results.map((item, index) => {
				const ingredientName = item.ingredient?.name || item.ingredient?.originalText || "Unknown Ingredient";
				const quantity = item.ingredient?.quantityGrams ? `${item.ingredient.quantityGrams}g` : "";
				const title = quantity ? `${ingredientName} (${quantity})` : ingredientName;

				return (
					<div key={item.ingredient?.name || index} className="space-y-4">
						<h3 className="text-2xl font-black bg-chart-2 inline-block px-4 py-1 border-2 border-border shadow-[4px_4px_0px_0px_var(--border)]">
							{title}
						</h3>

						{item.products && item.products.length > 0 ? (
							<div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
								{item.products.map((product, pIndex) => (
									<ProductCard
										key={product.link?.value || pIndex}
										product={product}
										ingredientName={ingredientName}
									/>
								))}
							</div>
						) : (
							<div className="p-6 border-2 border-dashed border-border rounded-lg bg-muted/20 text-muted-foreground font-mono">
								No products found for this ingredient.
							</div>
						)}
					</div>
				);
			})}
		</div>
	);
}
