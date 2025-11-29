import { cva, type VariantProps } from "class-variance-authority";

import type * as React from "react";

import { cn } from "@/lib/utils";

const sectionBadgeVariants = cva("inline-block border-2 border-border font-black shadow-shadow", {
	variants: {
		color: {
			"chart-1": "bg-chart-1",
			"chart-2": "bg-chart-2",
			"chart-5": "bg-chart-5",
		},
		size: {
			sm: "px-2 py-0.5 text-base",
			lg: "px-4 py-1 text-2xl",
		},
		uppercase: {
			true: "uppercase",
			false: "",
		},
	},
	defaultVariants: {
		color: "chart-1",
		size: "sm",
		uppercase: false,
	},
});

const SectionBadge = ({
	className,
	color,
	size,
	uppercase,
	...props
}: React.ComponentProps<"h3"> & VariantProps<typeof sectionBadgeVariants>) => {
	return <h3 className={cn(sectionBadgeVariants({ color, size, uppercase }), className)} {...props} />;
};

export { SectionBadge, sectionBadgeVariants };
