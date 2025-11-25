import { ExternalLink } from "lucide-react";
import type { OrchestratorIngredientProducts, OrchestratorProductFields } from "@/generated/kosik";
import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "./ui/card";

export interface Product {
	name: string;
	price: number;
	imageUrl: string;
	productUrl: string;
	ingredientName: string;
	unit?: string;
	quantity?: number;
}

interface ProductCardProps {
	ingredientProduct: OrchestratorIngredientProducts;
	product: OrchestratorProductFields;
}

export const ProductCard = ({ ingredientProduct, product }: ProductCardProps) => {
	const formatPrice = (price: number) => {
		return new Intl.NumberFormat("cs-CZ", {
			style: "currency",
			currency: "CZK",
		}).format(price);
	};

	const imageUrl = product.imageUrl?.value ?? "";
	const name = product.name?.value ?? "";
	const price = product.price?.value ?? 0;
	const productUrl = product.link?.value ?? "";
	const ingredientName = ingredientProduct.ingredient?.name ?? "";
	const unit = product.unit?.value;
	const quantity = ingredientProduct.ingredient?.quantityGrams;

	return (
		<Card className="group overflow-hidden">
			<CardHeader className="p-0">
				<div className="relative aspect-square overflow-hidden bg-foreground/5">
					<img
						src={imageUrl}
						alt={name}
						className="h-full w-full object-cover transition-transform group-hover:scale-105"
						loading="lazy"
					/>
				</div>
			</CardHeader>
			<CardContent className="p-4">
				<div className="mb-2 inline-block rounded-[--radius-base] border-2 border-border bg-main px-3 py-1 text-sm font-heading text-main-foreground shadow-[2px_2px_0px_0px_var(--border)]">
					{ingredientName}
				</div>
				<CardTitle className="mb-2 line-clamp-2 text-lg">{name}</CardTitle>
				{quantity && unit && (
					<CardDescription className="mb-2">
						{quantity} {unit}
					</CardDescription>
				)}
			</CardContent>
			<CardFooter className="flex items-center justify-between border-t-4 border-border p-4">
				<div className="text-2xl font-heading">{formatPrice(price)}</div>
				<a
					href={productUrl}
					target="_blank"
					rel="noopener noreferrer"
					className="inline-flex items-center gap-2 rounded-[--radius-base] border-4 border-border bg-foreground px-4 py-2 font-heading text-sm text-background shadow-shadow transition-all hover:shadow-[6px_6px_0px_0px_var(--border)] hover:translate-x-[-2px] hover:translate-y-[-2px]"
				>
					View
					<ExternalLink className="h-4 w-4" />
				</a>
			</CardFooter>
		</Card>
	);
};
