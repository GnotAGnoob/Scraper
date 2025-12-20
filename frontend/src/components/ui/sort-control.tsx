import type { LucideIcon } from "lucide-react";
import { ArrowDown, ArrowUp } from "lucide-react";
import { Button } from "./button";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./select";

export type SortDirection = "asc" | "desc";

export interface SortOption<T extends string> {
	label: string;
	value: T;
	icon?: LucideIcon;
}

interface BaseSortControlProps<T extends string> {
	sortDirection: SortDirection;
	options: SortOption<T>[];
	onDirectionToggle: (direction: SortDirection) => void;
	className?: string;
	defaultIcon?: LucideIcon;
}

interface SortControlPropsWithoutDefault<T extends string> extends BaseSortControlProps<T> {
	sortKey: T;
	onSortChange: (value: T) => void;
	showDefaultOption?: false;
	defaultLabel?: never;
}

interface SortControlPropsWithDefault<T extends string> extends BaseSortControlProps<T> {
	sortKey: T | "default";
	onSortChange: (value: T | "default") => void;
	showDefaultOption: true;
	defaultLabel?: string;
}

type SortControlProps<T extends string> = SortControlPropsWithoutDefault<T> | SortControlPropsWithDefault<T>;

export const SortControl = <T extends string>({
	sortKey,
	sortDirection,
	options,
	onSortChange,
	onDirectionToggle,
	showDefaultOption = false,
	defaultLabel = "Default",
	defaultIcon,
	className,
}: SortControlProps<T>) => {
	const handleDirectionToggle = () => {
		onDirectionToggle(sortDirection === "asc" ? "desc" : "asc");
	};

	return (
		<div className={`flex items-center gap-2 ${className || ""}`}>
			<Select value={sortKey} onValueChange={onSortChange}>
				<SelectTrigger variant="default" className="w-54">
					<SelectValue placeholder="Sort by..." />
				</SelectTrigger>
				<SelectContent>
					{showDefaultOption &&
						(() => {
							const DefaultIcon = defaultIcon;
							return (
								<SelectItem value="default">
									<div className="flex items-center gap-2">
										{DefaultIcon && <DefaultIcon className="h-4 w-4" />}
										<span>{defaultLabel}</span>
									</div>
								</SelectItem>
							);
						})()}
					{options.map((option) => {
						const Icon = option.icon;
						return (
							<SelectItem key={option.value} value={option.value}>
								<div className="flex items-center gap-1.5">
									{Icon && <Icon className="mt-px h-4 w-4" />}
									<span>{option.label}</span>
								</div>
							</SelectItem>
						);
					})}
				</SelectContent>
			</Select>

			<Button
				variant="default"
				size="icon"
				onClick={handleDirectionToggle}
				title={sortDirection === "asc" ? "Ascending" : "Descending"}
			>
				{sortDirection === "asc" ? <ArrowUp className="h-4 w-4" /> : <ArrowDown className="h-4 w-4" />}
			</Button>
		</div>
	);
};
