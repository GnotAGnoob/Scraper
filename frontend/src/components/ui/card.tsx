import type * as React from "react";

import { cn } from "@/lib/utils";

const Card = ({ className, ...props }: React.ComponentProps<"div">) => {
	return (
		<div
			data-slot="card"
			className={cn(
				"rounded-(--radius) flex flex-col shadow-shadow border-2 gap-2 py-4 border-border bg-secondary-background text-foreground font-base",
				className,
			)}
			{...props}
		/>
	);
};

export { Card };
