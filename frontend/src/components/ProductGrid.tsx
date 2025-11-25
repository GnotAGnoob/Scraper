import type { OrchestratorIngredientProducts } from "@/generated/kosik";
import { type FlattenedProduct, ProductCard } from "./ProductCard";

type ProductGridProps = {
	products: OrchestratorIngredientProducts[];
	isLoading: boolean;
	isError: boolean;
};

const flattenProducts = (ingredientsProducts: OrchestratorIngredientProducts[]): FlattenedProduct[] => {
	const flattened: FlattenedProduct[] = [];

	for (const ingredientProduct of ingredientsProducts) {
		const matchedIngredient = ingredientProduct.ingredient?.name ?? "";
		const products = ingredientProduct.products ?? [];

		for (const product of products) {
			const name = product.name?.value;
			const price = product.price?.value;
			const currency = product.currency?.value;
			const imageUrl = product.imageUrl?.value ?? null;
			const link = product.link?.value;

			if (name && price !== undefined && currency && link) {
				flattened.push({
					id: `${matchedIngredient}-${name}-${link}`,
					name,
					price,
					currency,
					imageUrl,
					productUrl: link,
					matchedIngredient,
				});
			}
		}
	}

	return flattened;
};

const SKELETON_KEYS = ["skeleton-0", "skeleton-1", "skeleton-2", "skeleton-3", "skeleton-4", "skeleton-5"];

export const ProductGrid = ({ products = [], isLoading, isError }: ProductGridProps) => {
	const flattenedProducts = flattenProducts(products);
	if (isLoading) {
		return (
			<div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
				{SKELETON_KEYS.map((key) => (
					<div
						key={key}
						className="h-40 animate-pulse rounded-[var(--radius-base)] border-2 border-dashed border-border bg-secondary-background/70 shadow-shadow"
					/>
				))}
			</div>
		);
	}

	if (isError) {
		return (
			<div className="rounded-[var(--radius-base)] border-2 border-border bg-secondary-background px-4 py-3 text-sm font-base text-red-700 shadow-shadow dark:text-red-300">
				Something went wrong while talking to the server. Try again in a bit.
			</div>
		);
	}

	if (flattenedProducts.length === 0) {
		return (
			<div className="rounded-[var(--radius-base)] border-2 border-dashed border-border bg-secondary-background/60 px-4 py-6 text-sm font-base text-neutral-700 shadow-shadow dark:text-neutral-200">
				Matched products will appear here. Submit a recipe to see Kosik suggestions.
			</div>
		);
	}

	return (
		<div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3">
			{flattenedProducts.map((product) => (
				<ProductCard key={product.id} product={product} />
			))}
		</div>
	);
};
