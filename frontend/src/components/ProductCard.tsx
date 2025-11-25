import { ExternalLink } from "lucide-react";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export type FlattenedProduct = {
	id: string;
	name: string;
	price: number;
	currency: string;
	imageUrl: string | null;
	productUrl: string;
	matchedIngredient: string;
};

type ProductCardProps = {
	product: FlattenedProduct;
};

export const ProductCard = ({ product }: ProductCardProps) => {
	const formattedPrice = `${product.price.toFixed(2)} ${product.currency}`;

	return (
		<Card className="flex flex-col gap-3">
			<CardHeader className="gap-2">
				<CardTitle className="text-base md:text-lg">{product.name}</CardTitle>
				<span className="inline-flex rounded-full border-2 border-border bg-main px-2 py-0.5 text-xs font-heading uppercase shadow-shadow">
					{product.matchedIngredient}
				</span>
			</CardHeader>
			<CardContent className="flex flex-1 flex-col gap-3">
				<div className="flex items-center justify-between gap-3">
					<div className="flex flex-col gap-1">
						<span className="text-lg font-heading">{formattedPrice}</span>
					</div>
					{product.imageUrl ? (
						<img
							src={product.imageUrl}
							alt={product.name}
							className="h-20 w-20 rounded-[var(--radius-base)] border-2 border-border bg-secondary-background object-contain shadow-shadow"
						/>
					) : null}
				</div>

				<a
					href={product.productUrl}
					target="_blank"
					rel="noreferrer"
					className="inline-flex items-center justify-center gap-2 rounded-[var(--radius-base)] border-2 border-border bg-secondary-background px-3 py-1 text-xs font-heading uppercase tracking-wide shadow-shadow transition-transform hover:-translate-y-[2px] hover:translate-x-[2px] hover:shadow-[5px_5px_0_0_var(--border)]"
				>
					Open on Kosik
					<ExternalLink className="h-3 w-3" />
				</a>
			</CardContent>
		</Card>
	);
};
