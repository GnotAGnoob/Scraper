import * as React from "react"
import { Slot } from "@radix-ui/react-slot"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const buttonVariants = cva(
	"inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-[var(--radius-base)] border-2 border-border bg-main px-5 py-2 text-base font-semibold uppercase tracking-tight text-main-foreground shadow-[var(--shadow)] transition-all duration-150 disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 [&_svg]:shrink-0 hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[6px_6px_0px_0px_var(--border)] focus-visible:outline-none focus-visible:ring-4 focus-visible:ring-ring/60 active:translate-x-0 active:translate-y-0 active:shadow-[var(--shadow)]",
	{
		variants: {
			variant: {
				default: "bg-main text-main-foreground",
				destructive:
					"bg-red-500 text-white hover:bg-red-500/90 focus-visible:ring-red-500/30",
				outline:
					"bg-secondary-background text-foreground hover:bg-main hover:text-main-foreground",
				secondary:
					"bg-foreground text-background hover:bg-foreground/90 hover:text-background",
				ghost:
					"bg-transparent text-foreground hover:bg-secondary-background",
				link: "bg-transparent text-foreground underline-offset-4 hover:underline",
			},
			size: {
				default: "px-5 py-2",
				sm: "px-4 py-1.5 text-sm",
				lg: "px-8 py-3 text-lg",
				icon: "size-10",
				"icon-sm": "size-8",
				"icon-lg": "size-12",
			},
		},
		defaultVariants: {
			variant: "default",
			size: "default",
		},
	}
)

function Button({
  className,
  variant,
  size,
  asChild = false,
  ...props
}: React.ComponentProps<"button"> &
  VariantProps<typeof buttonVariants> & {
    asChild?: boolean
  }) {
  const Comp = asChild ? Slot : "button"

  return (
    <Comp
      data-slot="button"
      className={cn(buttonVariants({ variant, size, className }))}
      {...props}
    />
  )
}

export { Button, buttonVariants }
