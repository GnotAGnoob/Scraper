import { cn } from "@/lib/utils";

interface NutritionRowProps {
	label: string;
	value: number | undefined | null;
	unit?: string;
	subLabel?: string;
	subValue?: number | undefined | null;
	className?: string;
}

const formatter = new Intl.NumberFormat("en-US", {
	maximumFractionDigits: 1,
	minimumFractionDigits: 0,
});

export const NutritionItem = ({ label, value, unit = "g", subLabel, subValue, className }: NutritionRowProps) => {
	if (value === undefined || value === null) return null;

	return (
		<div className={cn("flex items-center justify-between text-xs", className)}>
			<span className="text-muted-foreground">{label}</span>
			<div className="flex items-baseline gap-1.5">
				<span className="font-medium text-foreground">{`${formatter.format(value)} ${unit}`}</span>
				{subValue && (
					<span className="text-2xs text-muted-foreground">
						({subLabel} {`${formatter.format(subValue)} ${unit}`})
					</span>
				)}
			</div>
		</div>
	);
};
