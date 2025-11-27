import { cva, type VariantProps } from "class-variance-authority";

import type * as React from "react";

import { cn } from "@/lib/utils";

const alertVariants = cva(
	"relative grid w-full grid-cols-[0_1fr] items-start rounded-(--radius) border-2 border-border px-3 py-2.5 text-sm shadow-shadow has-[>svg]:grid-cols-[calc(var(--spacing)*4)_1fr] has-[>svg]:gap-x-3 [&>svg]:size-4 [&>svg]:translate-y-0.5 [&>svg]:text-current",
	{
		variants: {
			variant: {
				default: "bg-main text-main-foreground",
				destructive: "bg-border text-main-foreground",
			},
		},
		defaultVariants: {
			variant: "default",
		},
	},
);

const Alert = ({ className, variant, ...props }: React.ComponentProps<"div"> & VariantProps<typeof alertVariants>) => {
	return <div data-slot="alert" role="alert" className={cn(alertVariants({ variant }), className)} {...props} />;
};

const AlertTitle = ({ className, ...props }: React.ComponentProps<"div">) => {
	return (
		<div
			data-slot="alert-title"
			className={cn("col-start-2 line-clamp-1 min-h-4 font-black text-xl tracking-tight", className)}
			{...props}
		/>
	);
};

const AlertDescription = ({ className, ...props }: React.ComponentProps<"div">) => {
	return (
		<div
			data-slot="alert-description"
			className={cn(
				"col-start-2 grid justify-items-start gap-1 font-mono text-lg [&_p]:leading-relaxed",
				className,
			)}
			{...props}
		/>
	);
};

export { Alert, AlertTitle, AlertDescription };
