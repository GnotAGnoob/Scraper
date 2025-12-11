import * as SelectPrimitive from "@radix-ui/react-select";
import { cva, type VariantProps } from "class-variance-authority";
import { Check, ChevronDown, ChevronUp } from "lucide-react";

import type * as React from "react";

import { cn } from "@/lib/utils";

const selectTriggerVariants = cva(
	"flex h-10 w-full cursor-pointer items-center justify-between gap-2 rounded-base border-2 border-border px-3 py-2 font-base text-sm ring-offset-white transition-all placeholder:text-foreground/50 focus:outline-hidden focus:ring-2 focus:ring-black focus:ring-offset-2 focus-visible:outline-hidden focus-visible:ring-2 focus-visible:ring-black focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 *:data-[slot=select-value]:line-clamp-1 *:data-[slot=select-value]:flex *:data-[slot=select-value]:items-center *:data-[slot=select-value]:gap-2 [&>span]:line-clamp-1 [&_svg:not([class*='size-'])]:size-4 [&_svg]:pointer-events-none [&_svg]:shrink-0",
	{
		variants: {
			variant: {
				default:
					"bg-main shadow-shadow hover:translate-x-boxShadowX hover:translate-y-boxShadowY hover:shadow-none data-[state=open]:translate-x-boxShadowX data-[state=open]:translate-y-boxShadowY data-[state=open]:shadow-none",
				noShadow: "bg-main text-main-foreground",
				neutral:
					"bg-secondary-background text-foreground shadow-shadow hover:translate-x-boxShadowX hover:translate-y-boxShadowY hover:shadow-none data-[state=open]:translate-x-boxShadowX data-[state=open]:translate-y-boxShadowY data-[state=open]:shadow-none",
				reverse:
					"bg-main text-main-foreground hover:translate-x-reverseBoxShadowX hover:translate-y-reverseBoxShadowY hover:shadow-shadow data-[state=open]:translate-x-reverseBoxShadowX data-[state=open]:translate-y-reverseBoxShadowY data-[state=open]:shadow-shadow",
			},
		},
		defaultVariants: {
			variant: "default",
		},
	},
);

const Select = ({ ...props }: React.ComponentProps<typeof SelectPrimitive.Root>) => {
	return <SelectPrimitive.Root data-slot="select" {...props} />;
};

const SelectGroup = ({ ...props }: React.ComponentProps<typeof SelectPrimitive.Group>) => {
	return <SelectPrimitive.Group data-slot="select-group" {...props} />;
};

const SelectValue = ({ ...props }: React.ComponentProps<typeof SelectPrimitive.Value>) => {
	return <SelectPrimitive.Value data-slot="select-value" {...props} />;
};

const SelectTrigger = ({
	className,
	variant,
	children,
	...props
}: React.ComponentProps<typeof SelectPrimitive.Trigger> & VariantProps<typeof selectTriggerVariants>) => {
	return (
		<SelectPrimitive.Trigger
			data-slot="select-trigger"
			className={cn(selectTriggerVariants({ variant, className }))}
			{...props}
		>
			{children}
			<SelectPrimitive.Icon asChild>
				<ChevronDown className="size-4" />
			</SelectPrimitive.Icon>
		</SelectPrimitive.Trigger>
	);
};

const SelectScrollUpButton = ({ className, ...props }: React.ComponentProps<typeof SelectPrimitive.ScrollUpButton>) => {
	return (
		<SelectPrimitive.ScrollUpButton
			data-slot="select-scroll-up"
			className={cn("flex cursor-default items-center justify-center py-1 font-base", className)}
			{...props}
		>
			<ChevronUp className="size-4" />
		</SelectPrimitive.ScrollUpButton>
	);
};

const SelectScrollDownButton = ({
	className,
	...props
}: React.ComponentProps<typeof SelectPrimitive.ScrollDownButton>) => {
	return (
		<SelectPrimitive.ScrollDownButton
			data-slot="select-scroll-down"
			className={cn("flex cursor-default items-center justify-center py-1 font-base", className)}
			{...props}
		>
			<ChevronDown className="size-4" />
		</SelectPrimitive.ScrollDownButton>
	);
};

const SelectContent = ({
	className,
	children,
	position = "popper",
	...props
}: React.ComponentProps<typeof SelectPrimitive.Content>) => {
	return (
		<SelectPrimitive.Portal>
			<SelectPrimitive.Content
				data-slot="select-content"
				className={cn(
					"data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2 relative z-50 max-h-96 min-w-32 origin-(--radix-select-content-transform-origin) overflow-hidden rounded-base border-2 border-border bg-main shadow-shadow data-[state=closed]:animate-out data-[state=open]:animate-in",
					position === "popper" &&
						"data-[side=left]:-translate-x-1 data-[side=top]:-translate-y-1 data-[side=right]:translate-x-1 data-[side=bottom]:translate-y-1",
					className,
				)}
				position={position}
				{...props}
			>
				<SelectScrollUpButton />
				<SelectPrimitive.Viewport
					className={cn(
						"p-1",
						position === "popper" &&
							"h-(--radix-select-trigger-height) w-full min-w-(--radix-select-trigger-width)",
					)}
				>
					{children}
				</SelectPrimitive.Viewport>
				<SelectScrollDownButton />
			</SelectPrimitive.Content>
		</SelectPrimitive.Portal>
	);
};

const SelectLabel = ({ className, ...props }: React.ComponentProps<typeof SelectPrimitive.Label>) => {
	return (
		<SelectPrimitive.Label
			data-slot="select-label"
			className={cn(
				"border-2 border-transparent py-1.5 pr-8 pl-2 font-base text-main-foreground/80 text-sm",
				className,
			)}
			{...props}
		/>
	);
};

const SelectItem = ({ className, children, ...props }: React.ComponentProps<typeof SelectPrimitive.Item>) => {
	return (
		<SelectPrimitive.Item
			data-slot="select-item"
			className={cn(
				"relative flex w-full cursor-pointer select-none items-center gap-2 rounded-base border-2 border-transparent py-1.5 pr-8 pl-2 font-base text-sm outline-none focus:border-border data-disabled:pointer-events-none data-disabled:opacity-50 [&_svg:not([class*='size-'])]:size-4 [&_svg]:pointer-events-none [&_svg]:shrink-0 *:[span]:last:flex *:[span]:last:items-center *:[span]:last:gap-2",
				className,
			)}
			{...props}
		>
			<span className="absolute right-2 flex size-3.5 items-center justify-center">
				<SelectPrimitive.ItemIndicator>
					<Check className="size-4" />
				</SelectPrimitive.ItemIndicator>
			</span>
			<SelectPrimitive.ItemText>{children}</SelectPrimitive.ItemText>
		</SelectPrimitive.Item>
	);
};

const SelectSeparator = ({ className, ...props }: React.ComponentProps<typeof SelectPrimitive.Separator>) => {
	return (
		<SelectPrimitive.Separator
			data-slot="select-separator"
			className={cn("-mx-1 my-1 h-px bg-border", className)}
			{...props}
		/>
	);
};

export {
	Select,
	SelectGroup,
	SelectValue,
	SelectTrigger,
	SelectContent,
	SelectLabel,
	SelectItem,
	SelectSeparator,
	SelectScrollUpButton,
	SelectScrollDownButton,
	selectTriggerVariants,
};
export type { VariantProps };
