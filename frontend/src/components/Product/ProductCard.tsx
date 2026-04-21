import { AlertCircle, ExternalLink } from "lucide-react";
import { cn } from "@/lib/utils";
import type { OrchestratorProductFields } from "../../generated/kosik";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";
import { Card } from "../ui/card";
import { TruncatedText } from "../ui/truncated-text";
import { NutritionItem } from "./NutritionItem";
import { ProductSales } from "./ProductSales";
import { hasAnySale } from "./sales";

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
	const productQuantity = product.productQuantity?.value;
	const ingredients = product.nutrition?.value?.ingredients?.value;
	const nutrition = product.nutrition?.value;
	const sales = product.sales?.value;
	const showSales = hasAnySale(sales);
	const recommendedPrice = sales?.recommendedPrice;
	const hasPriceStrikethrough = price !== undefined && recommendedPrice !== undefined && recommendedPrice > price;

	const formatCurrency = Intl.NumberFormat("cs-CZ", {
		style: "currency",
		currency: currency || "CZK",
		maximumFractionDigits: 2,
	});

	return (
		<Card className="flex h-full flex-col gap-0 overflow-hidden pt-0 pb-2 transition-all hover:shadow-(--shadow-hover)">
			<div className="relative aspect-square w-full shrink-0 overflow-hidden border-border border-b bg-white">
				{imageUrl ? (
					<a href={link} target="_blank" rel="noopener noreferrer">
						<img
							src={imageUrl}
							alt={name}
							className={cn(
								"h-full w-full object-contain p-4 transition-transform hover:scale-105",
								isSoldOut && "opacity-70 saturate-50",
							)}
							loading="lazy"
						/>
					</a>
				) : (
					<div className="flex h-full items-center justify-center bg-muted/20 text-muted-foreground">
						<span className="text-xs">No Image</span>
					</div>
				)}
				{isSoldOut && (
					<div className="absolute inset-0 flex items-center justify-center bg-main/20">
						<Badge className="font-bold text-lg">Sold Out</Badge>
					</div>
				)}
				{!isSoldOut && sales?.percentageDiscount ? (
					<Badge className="absolute top-2 left-2 bg-red-500 font-bold text-white" size="lg">
						-{sales.percentageDiscount}%
					</Badge>
				) : null}
				{!isSoldOut && !sales?.percentageDiscount && sales?.actionLabel ? (
					<Badge className="absolute top-2 left-2 bg-red-500 font-bold text-white" size="lg">
						{sales.actionLabel}
					</Badge>
				) : null}
			</div>

			<div className="flex grow flex-col justify-between px-2 pb-2">
				<div className="space-y-0.5">
					<div className="flex flex-col">
						<div className="flex items-center justify-between">
							{price !== undefined ? (
								<span className="flex items-baseline gap-1.5">
									<span
										className={cn(
											"font-bold leading-tight tracking-tight",
											hasPriceStrikethrough
												? "text-red-600 text-xl dark:text-red-400"
												: "text-foreground text-xl",
										)}
									>
										{formatCurrency.format(price)}
									</span>
									{hasPriceStrikethrough && recommendedPrice !== undefined ? (
										<span className="text-muted-foreground text-xs line-through">
											{formatCurrency.format(recommendedPrice)}
										</span>
									) : null}
								</span>
							) : (
								<span className="flex items-center gap-1.5 font-medium text-muted-foreground text-sm leading-tight">
									<AlertCircle className="h-3.5 w-3.5" /> N/A
								</span>
							)}

							{productQuantity && (
								<span className="font-medium text-muted-foreground text-xs leading-tight">
									{[productQuantity.prefix, productQuantity.value, productQuantity.unit]
										.filter(Boolean)
										.join(" ")}
								</span>
							)}
						</div>
						{pricePerUnit?.value && (
							<span className="font-medium text-muted-foreground text-xs leading-tight">
								{formatCurrency.format(pricePerUnit.value)} / {pricePerUnit.unit}
							</span>
						)}
					</div>

					<TruncatedText
						text={name}
						className="text-left font-semibold text-sm leading-tight tracking-tight"
						maxLines={2}
					/>

					{showSales && <ProductSales product={product} formatCurrency={formatCurrency} />}
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
						<Button asChild size="sm" className="h-8 w-full gap-2 text-xs">
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
