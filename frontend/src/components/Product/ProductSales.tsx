import { Crown, Minus, Plus, Tag, Users } from "lucide-react";
import { useMemo, useState } from "react";
import { cn } from "@/lib/utils";
import type { OrchestratorProductFields } from "../../generated/kosik";
import { Badge } from "../ui/badge";
import { Button } from "../ui/button";
import { computeEffectivePrice, getActiveBulkTier } from "./sales";

interface ProductSalesProps {
	product: OrchestratorProductFields;
	formatCurrency: Intl.NumberFormat;
}

export const ProductSales = ({ product, formatCurrency }: ProductSalesProps) => {
	const sales = product.sales?.value;
	const basePrice = product.price?.value;
	const [quantity, setQuantity] = useState(1);
	const [isMember, setIsMember] = useState(false);

	const effective = useMemo(
		() => computeEffectivePrice(product, { quantity, isMember }),
		[product, quantity, isMember],
	);

	if (!sales || basePrice === undefined) {
		return null;
	}

	const activeBulkTier = getActiveBulkTier(sales, quantity);
	const hasMembershipOffer = sales.membershipPrice !== undefined;
	const hasBulkOffer = Boolean(sales.cumulativePrices && sales.cumulativePrices.length > 0);
	const hasDirectSale = Boolean(sales.percentageDiscount || sales.recommendedPrice || sales.actionLabel);
	const hasLabels = Boolean(sales.labels && sales.labels.length > 0);

	const totalSavings = basePrice * quantity - effective.totalPrice;

	return (
		<div className="mt-1.5 flex flex-col gap-2 rounded-md border-2 border-border/60 bg-main/5 p-2">
			<div className="flex items-center justify-between gap-1">
				<span className="font-heading text-[10px] text-muted-foreground uppercase tracking-wider">Offers</span>
				{totalSavings > 0 && (
					<span className="font-semibold text-[10px] text-green-700 dark:text-green-500">
						You save {formatCurrency.format(totalSavings)}
					</span>
				)}
			</div>

			{(hasDirectSale || hasLabels) && (
				<div className="flex flex-wrap items-center gap-1">
					{sales.percentageDiscount ? (
						<Badge variant="default" size="sm" className="gap-1 bg-red-500 text-white">
							<Tag className="h-2.5 w-2.5" />-{sales.percentageDiscount}%
						</Badge>
					) : null}
					{sales.actionLabel ? (
						<Badge variant="default" size="sm" className="bg-red-500 text-white">
							{sales.actionLabel}
						</Badge>
					) : null}
					{sales.labels?.map((label) => (
						<Badge
							key={`${label.name}-${label.priority ?? 0}`}
							variant="neutral"
							size="sm"
							style={label.background ? { backgroundColor: label.background, color: "#fff" } : undefined}
						>
							{label.name}
						</Badge>
					))}
					{sales.recommendedPrice && sales.recommendedPrice > basePrice ? (
						<span className="text-[11px] text-muted-foreground line-through">
							{formatCurrency.format(sales.recommendedPrice)}
						</span>
					) : null}
				</div>
			)}

			{hasMembershipOffer && (
				<div className="flex items-center justify-between gap-2 rounded-md bg-yellow-100/60 p-1.5 dark:bg-yellow-950/40">
					<div className="flex min-w-0 items-center gap-1.5">
						<Crown className="h-3 w-3 shrink-0 text-yellow-700 dark:text-yellow-400" />
						<div className="flex flex-col leading-tight">
							<span className="font-semibold text-[11px] text-foreground">Member price</span>
							<span className="text-[10px] text-muted-foreground">
								{formatCurrency.format(sales.membershipPrice as number)} per unit
							</span>
						</div>
					</div>
					<Button
						type="button"
						size="sm"
						variant={isMember ? "default" : "neutral"}
						className="h-6 px-2 text-[10px]"
						onClick={() => setIsMember((v) => !v)}
					>
						<Users className="h-2.5 w-2.5" />
						{isMember ? "Active" : "Apply"}
					</Button>
				</div>
			)}

			{hasBulkOffer && (
				<div className="flex flex-col gap-1 rounded-md bg-blue-100/50 p-1.5 dark:bg-blue-950/40">
					<span className="font-semibold text-[11px] text-foreground">Bulk pricing</span>
					<div className="flex flex-col gap-0.5">
						{sales.cumulativePrices
							?.slice()
							.sort((a, b) => a.quantity - b.quantity)
							.map((tier) => {
								const isActive = activeBulkTier?.quantity === tier.quantity;
								return (
									<div
										key={tier.quantity}
										className={cn(
											"flex items-center justify-between text-[10px]",
											isActive ? "font-semibold text-foreground" : "text-muted-foreground",
										)}
									>
										<span>
											{isActive ? "✓ " : ""}Buy {tier.quantity}+
										</span>
										<span>{formatCurrency.format(tier.price)} / unit</span>
									</div>
								);
							})}
					</div>
				</div>
			)}

			<div className="flex items-center justify-between gap-1.5 border-border/40 border-t pt-1.5">
				<div className="flex items-center gap-1">
					<Button
						type="button"
						size="icon"
						variant="neutral"
						className="h-6 w-6"
						onClick={() => setQuantity((q) => Math.max(1, q - 1))}
						disabled={quantity <= 1}
						aria-label="Decrease quantity"
					>
						<Minus className="h-3 w-3" />
					</Button>
					<span className="w-6 text-center font-semibold text-foreground text-xs tabular-nums">
						{quantity}
					</span>
					<Button
						type="button"
						size="icon"
						variant="neutral"
						className="h-6 w-6"
						onClick={() => setQuantity((q) => q + 1)}
						aria-label="Increase quantity"
					>
						<Plus className="h-3 w-3" />
					</Button>
				</div>
				<div className="flex min-w-0 flex-col items-end leading-tight">
					<span className="font-bold text-foreground text-sm tabular-nums">
						{formatCurrency.format(effective.totalPrice)}
					</span>
					{quantity > 1 && (
						<span className="text-[10px] text-muted-foreground tabular-nums">
							{formatCurrency.format(effective.unitPrice)} × {quantity}
						</span>
					)}
				</div>
			</div>

			{effective.appliedDiscounts.length > 0 && (
				<div className="flex flex-wrap gap-1">
					{effective.appliedDiscounts.map((d) => (
						<Badge key={d.type + d.label} variant="default" size="sm" className="bg-green-600 text-white">
							{d.label}
						</Badge>
					))}
				</div>
			)}
		</div>
	);
};
