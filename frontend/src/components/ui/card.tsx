import type * as React from "react";

import { cn } from "@/lib/utils";

const Card = ({ className, ...props }: React.ComponentProps<"div">) => {
	return (
		<div
			data-slot="card"
			className={cn(
				"flex flex-col gap-2 rounded-(--radius) border-2 border-border bg-secondary-background py-4 font-base text-foreground shadow-shadow",
				className,
			)}
			{...props}
		/>
	);
};

export { Card };
