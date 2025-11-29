import { Slot } from "@radix-ui/react-slot";
import { cva, type VariantProps } from "class-variance-authority";

import type * as React from "react";

import { cn } from "@/lib/utils";

const badgeVariants = cva(
	"inline-flex w-fit shrink-0 items-center justify-center gap-1 whitespace-nowrap rounded-(--radius) border-2 border-border font-base [&>svg]:pointer-events-none",
	{
		variants: {
			variant: {
				default: "bg-main text-main-foreground",
				neutral: "bg-secondary-background text-foreground",
			},
			size: {
				sm: "px-2 py-0.5 text-[10px] [&>svg]:size-2.5",
				default: "px-2.5 py-0.5 text-xs [&>svg]:size-3",
				lg: "px-3 py-1 text-sm [&>svg]:size-3.5",
			},
		},
		defaultVariants: {
			variant: "default",
			size: "default",
		},
	},
);

const Badge = ({
	className,
	variant,
	size,
	asChild = false,
	...props
}: React.ComponentProps<"span"> &
	VariantProps<typeof badgeVariants> & {
		asChild?: boolean;
	}) => {
	const Comp = asChild ? Slot : "span";

	return <Comp data-slot="badge" className={cn(badgeVariants({ variant, size }), className)} {...props} />;
};

export { Badge, badgeVariants };
