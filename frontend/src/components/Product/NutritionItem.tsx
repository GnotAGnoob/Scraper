interface NutritionItemProps {
	label: string;
	value?: number;
	unit?: string;
}

export const NutritionItem = ({ label, value, unit = "g" }: NutritionItemProps) => {
	if (value === undefined || value === null) return null;
	return (
		<div className="flex flex-col items-center justify-center rounded-(--radius) border border-border/20 p-1 text-center">
			<span className="font-bold text-xs">{label}</span>
			<span className="text-2xs text-muted-foreground">
				{new Intl.NumberFormat("en-US", {
					maximumFractionDigits: 1,
					minimumFractionDigits: 0,
				}).format(value)}
				{unit}
			</span>
		</div>
	);
};
