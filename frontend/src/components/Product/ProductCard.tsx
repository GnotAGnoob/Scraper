import { AlertCircle, ExternalLink } from "lucide-react";
import type { OrchestratorProductFields } from "../../generated/kosik";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";
import { Card } from "../ui/card";

interface ProductCardProps {
	product: OrchestratorProductFields;
	ingredientName?: string;
}

export const ProductCard = ({ product }: ProductCardProps) => {
	const name = product.name?.value || "Unknown Product";
	const price = product.price?.value;
	const currency = product.currency?.value;
	const imageUrl = product.imageUrl?.value;
	const link = product.link?.value;
	const isSoldOut = product.isSoldOut?.value;

	return (
		<Card className="gap-0 overflow-hidden pt-0 pb-2">
			<div className="relative aspect-square w-full overflow-hidden border-border border-b-2 bg-secondary-background">
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
						<Badge variant="neutral" className="bg-chart-3 font-bold">
							Sold Out
						</Badge>
					</div>
				)}
			</div>

			<div className="flex grow flex-col justify-between p-2">
				<h2 className="line-clamp-2 text-sm leading-tight">{name}</h2>
				<div>
					<div className="grow p-4 pt-0">
						<div className="mt-2 flex items-baseline gap-1">
							{price !== undefined ? (
								<>
									<span className="font-black text-2xl">{price}</span>
									<span className="font-bold text-muted-foreground text-sm">{currency}</span>
								</>
							) : (
								<span className="flex items-center gap-1 text-muted-foreground">
									<AlertCircle size={16} /> Price unknown
								</span>
							)}
						</div>
					</div>
					{link && (
						<Button asChild variant="noShadow" className="w-full font-bold">
							<a
								href={link}
								target="_blank"
								rel="noopener noreferrer"
								className="flex items-center justify-center gap-2"
							>
								View on Kosik <ExternalLink size={16} />
							</a>
						</Button>
					)}
				</div>
			</div>
		</Card>
	);
};
