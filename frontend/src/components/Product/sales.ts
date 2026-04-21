import type {
	OrchestratorCumulativePriceFields,
	OrchestratorProductFields,
	OrchestratorSalesFields,
} from "../../generated/kosik";

export interface EffectivePrice {
	unitPrice: number;
	totalPrice: number;
	quantity: number;
	appliedDiscounts: AppliedDiscount[];
}

export interface AppliedDiscount {
	type: "sale" | "membership" | "bulk";
	label: string;
	savingsPerUnit: number;
}

export interface PricingOptions {
	quantity: number;
	isMember: boolean;
}

export const hasAnySale = (sales: OrchestratorSalesFields | undefined): boolean => {
	if (!sales) return false;
	return Boolean(
		sales.percentageDiscount ||
			sales.recommendedPrice ||
			sales.actionLabel ||
			sales.membershipPrice ||
			(sales.cumulativePrices && sales.cumulativePrices.length > 0) ||
			(sales.labels && sales.labels.length > 0),
	);
};

export const getActiveBulkTier = (
	sales: OrchestratorSalesFields | undefined,
	quantity: number,
): OrchestratorCumulativePriceFields | undefined => {
	if (!sales?.cumulativePrices?.length) return undefined;
	const sorted = [...sales.cumulativePrices].sort((a, b) => b.quantity - a.quantity);
	return sorted.find((tier) => quantity >= tier.quantity);
};

export const computeEffectivePrice = (product: OrchestratorProductFields, options: PricingOptions): EffectivePrice => {
	const basePrice = product.price?.value ?? 0;
	const sales = product.sales?.value;
	const quantity = Math.max(1, options.quantity);

	let bestUnitPrice = basePrice;
	const appliedDiscounts: AppliedDiscount[] = [];

	if (sales?.percentageDiscount && sales.percentageDiscount > 0 && sales.recommendedPrice) {
		const savingsPerUnit = sales.recommendedPrice - basePrice;
		if (savingsPerUnit > 0) {
			appliedDiscounts.push({
				type: "sale",
				label: `-${sales.percentageDiscount}%`,
				savingsPerUnit,
			});
		}
	}

	if (options.isMember && sales?.membershipPrice && sales.membershipPrice < bestUnitPrice) {
		const savingsPerUnit = bestUnitPrice - sales.membershipPrice;
		bestUnitPrice = sales.membershipPrice;
		appliedDiscounts.push({
			type: "membership",
			label: "Member",
			savingsPerUnit,
		});
	}

	const bulkTier = getActiveBulkTier(sales, quantity);
	if (bulkTier && bulkTier.price < bestUnitPrice) {
		const savingsPerUnit = bestUnitPrice - bulkTier.price;
		bestUnitPrice = bulkTier.price;
		appliedDiscounts.push({
			type: "bulk",
			label: `Bulk x${bulkTier.quantity}`,
			savingsPerUnit,
		});
	}

	return {
		unitPrice: bestUnitPrice,
		totalPrice: bestUnitPrice * quantity,
		quantity,
		appliedDiscounts,
	};
};
