import { AlertCircle, ExternalLink } from "lucide-react";
import { cn } from "@/lib/utils";
import type { OrchestratorProductFields } from "../../generated/kosik";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";
import { Card } from "../ui/card";
import { TruncatedText } from "../ui/truncated-text";
import { NutritionItem } from "./NutritionItem";

interface ProductCardProps {
	product: OrchestratorProductFields;
}

export const ProductCard = ({ product }: ProductCardProps) => {
	const name = product.name?.value || "Unknown Product";
	const price = product.price?.value;
	const currency = product.currency?.value;
	const imageUrl = product.imageUrl?.value;
	const link = product.link?.value;
	const isSoldOut = product.isSoldOut?.value;
	const pricePerUnit = product.pricePerUnit?.value;
	const ingredients = product.nutrition?.value?.ingredients?.value;
	const nutrition = product.nutrition?.value;

	const formatCurrency = Intl.NumberFormat("cs-CZ", {
		style: "currency",
		currency: currency || "CZK",
		maximumFractionDigits: 2,
	});

	return (
		<Card className="flex h-full flex-col gap-0 overflow-hidden pt-0 pb-2 transition-all hover:shadow-(--shadow-hover)">
			<div className="relative aspect-square w-full shrink-0 overflow-hidden border-border border-b bg-secondary-background">
				{imageUrl ? (
					<img
						src={imageUrl}
						alt={name}
						className={cn(
							"h-full w-full object-contain p-4 transition-transform hover:scale-105",
							isSoldOut && "opacity-70 saturate-50",
						)}
						loading="lazy"
					/>
				) : (
					<div className="flex h-full items-center justify-center bg-muted/20 text-muted-foreground">
						<span className="text-xs">No Image</span>
					</div>
				)}
				{isSoldOut && (
					<div className="absolute inset-0 flex items-center justify-center bg-chart-3/40">
						<Badge variant="neutral" className="bg-chart-3 font-bold text-lg text-white">
							Sold Out
						</Badge>
					</div>
				)}
			</div>

			<div className="flex grow flex-col justify-between p-2">
				<div className="space-y-0.5">
					<TruncatedText
						text={name}
						className="text-left font-semibold text-sm leading-tight tracking-tight"
						maxLines={2}
					/>

					<div className="flex flex-col">
						{price !== undefined ? (
							<span className="font-bold text-foreground text-xl leading-tight tracking-tight">
								{formatCurrency.format(price)}
							</span>
						) : (
							<span className="flex items-center gap-1.5 font-medium text-muted-foreground text-sm leading-tight">
								<AlertCircle className="h-3.5 w-3.5" /> N/A
							</span>
						)}
						{pricePerUnit?.value && (
							<span className="font-medium text-muted-foreground text-xs leading-tight">
								{formatCurrency.format(pricePerUnit.value)} / {pricePerUnit.unit}
							</span>
						)}
					</div>
				</div>

				<div className="mt-1 mb-1 h-px w-full bg-border/20" />

				<div className="flex grow flex-col justify-between gap-3 px-1">
					{nutrition && (
						<div className="flex flex-col gap-0.5">
							<NutritionItem label="Calories" value={nutrition.calories?.value} unit="kcal" />
							<NutritionItem label="Protein" value={nutrition.protein?.value} unit="g" />
							<NutritionItem
								label="Fats"
								value={nutrition.fat?.value}
								unit="g"
								subLabel="Sat."
								subValue={nutrition.saturatedFat?.value}
							/>
							<NutritionItem
								label="Carbs"
								value={nutrition.carbs?.value}
								unit="g"
								subLabel="Sug."
								subValue={nutrition.sugar?.value}
							/>
							<NutritionItem label="Fiber" value={nutrition.fiber?.value} unit="g" />
						</div>
					)}

					{ingredients && (
						<div className="group relative flex rounded-md border border-border/20 p-2 text-xs">
							<span className="-top-2 absolute left-1.5 bg-secondary-background px-1.5 font-medium text-2xs text-muted-foreground">
								Ingredients
							</span>
							<TruncatedText
								text={ingredients}
								maxLines={2}
								className="text-left text-muted-foreground"
								truncatedClassName="cursor-help group-hover:text-foreground transition-colors"
							/>
						</div>
					)}
				</div>

				<div className="mt-4">
					{link ? (
						<Button asChild variant="neutral" size="sm" className="h-8 w-full gap-2 bg-chart-4 text-xs">
							<a href={link} target="_blank" rel="noopener noreferrer">
								View on Shop <ExternalLink size={12} />
							</a>
						</Button>
					) : (
						<Button disabled variant="neutral" size="sm" className="h-8 w-full text-xs opacity-50">
							No Link Available
						</Button>
					)}
				</div>
			</div>
		</Card>
	);
};
