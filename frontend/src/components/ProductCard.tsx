import { AlertCircle, ExternalLink } from "lucide-react";
import type { OrchestratorProductFields } from "../generated/kosik";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";
import { Card } from "./ui/card";

interface ProductCardProps {
	product: OrchestratorProductFields;
	ingredientName?: string;
}

export const ProductCard = ({ product, ingredientName }: ProductCardProps) => {
	const name = product.name?.value || "Unknown Product";
	const price = product.price?.value;
	const currency = product.currency?.value;
	const imageUrl = product.imageUrl?.value;
	const link = product.link?.value;
	const isSoldOut = product.isSoldOut?.value;

	return (
		<Card className="pt-0 gap-0">
			<div className="relative aspect-square w-full overflow-hidden bg-secondary-background border-b-2 border-border">
				{imageUrl ? (
					<img
						src={imageUrl}
						alt={imageUrl}
						className="h-full w-full object-contain transition-transform hover:scale-105"
						loading="lazy"
					/>
				) : (
					<div className="flex h-full items-center justify-center text-muted-foreground">No Image</div>
				)}
				{isSoldOut && (
					<div className="absolute top-2 right-2">
						<Badge variant="neutral" className="font-bold bg-chart-3">
							Sold Out
						</Badge>
					</div>
				)}
			</div>

			<div className="p-4 pb-2">
				<h2 className="line-clamp-2 text-lg leading-tight min-h-12">{name}</h2>
				<div className="p-4 pt-0 grow">
					<div className="flex items-baseline gap-1 mt-2">
						{price !== undefined ? (
							<>
								<span className="text-2xl font-black">{price}</span>
								<span className="text-sm font-bold text-muted-foreground">{currency}</span>
							</>
						) : (
							<span className="text-muted-foreground flex items-center gap-1">
								<AlertCircle size={16} /> Price unknown
							</span>
						)}
					</div>
				</div>
				{link && (
					<Button
						asChild
						className="w-full font-bold border-2 border-border shadow-[2px_2px_0px_0px_var(--border)] hover:shadow-[4px_4px_0px_0px_var(--border)] hover:-translate-x-px hover:-translate-y-px active:shadow-none active:translate-x-[2px] active:translate-y-[2px] transition-all bg-main hover:bg-main/90 text-main-foreground"
					>
						<a
							href={link}
							target="_blank"
							rel="noopener noreferrer"
							className="flex items-center gap-2 justify-center"
						>
							View on Kosik <ExternalLink size={16} />
						</a>
					</Button>
				)}
			</div>
		</Card>
	);
};
