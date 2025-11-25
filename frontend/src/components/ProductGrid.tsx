import type { OrchestratorIngredientProducts } from "@/generated/kosik";
import { ProductCard } from "./ProductCard";
import { Skeleton } from "./ui/skeleton";

interface ProductGridProps {
	products: OrchestratorIngredientProducts[];
	isLoading?: boolean;
}

const ProductCardSkeleton = () => (
	<div className="overflow-hidden rounded-[--radius-base] border-4 border-border bg-secondary-background shadow-shadow">
		<Skeleton className="aspect-square w-full" />
		<div className="p-4">
			<Skeleton className="mb-2 h-8 w-24" />
			<Skeleton className="mb-2 h-6 w-full" />
			<Skeleton className="mb-2 h-4 w-20" />
		</div>
		<div className="flex items-center justify-between border-t-4 border-border p-4">
			<Skeleton className="h-8 w-20" />
			<Skeleton className="h-10 w-24" />
		</div>
	</div>
);

const SKELETON_KEYS = Array.from({ length: 8 }, (_, i) => `skeleton-${i}`);

export const ProductGrid = ({ products, isLoading }: ProductGridProps) => {
	if (isLoading) {
		return (
			<div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
				{SKELETON_KEYS.map((key) => (
					<ProductCardSkeleton key={key} />
				))}
			</div>
		);
	}

	const allProductCards = products.flatMap((ingredientProduct) =>
		(ingredientProduct.products ?? []).map((product, productIndex) => ({
			ingredientProduct,
			product,
			key: `${ingredientProduct.ingredient?.name ?? "unknown"}-${productIndex}`,
		})),
	);

	if (allProductCards.length === 0) {
		return (
			<div className="flex min-h-[400px] items-center justify-center">
				<div className="text-center">
					<p className="text-2xl font-heading text-foreground/60">No products found</p>
					<p className="mt-2 text-foreground/40">Try entering a recipe to find matching products</p>
				</div>
			</div>
		);
	}

	return (
		<div className="grid gap-6 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
			{allProductCards.map(({ ingredientProduct, product, key }) => (
				<ProductCard key={key} ingredientProduct={ingredientProduct} product={product} />
			))}
		</div>
	);
};
