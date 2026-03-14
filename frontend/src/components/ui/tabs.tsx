import * as TabsPrimitive from "@radix-ui/react-tabs";

import type * as React from "react";

import { cn } from "@/lib/utils";

type TabsRootProps<T extends string> = Omit<
	React.ComponentProps<typeof TabsPrimitive.Root>,
	"value" | "onValueChange"
> & {
	value?: T;
	onValueChange?: (value: T) => void;
};

function Tabs<T extends string = string>({ className, value, onValueChange, ...props }: TabsRootProps<T>) {
	return (
		<TabsPrimitive.Root
			data-slot="tabs"
			className={cn("w-full", className)}
			value={value}
			onValueChange={onValueChange as (value: string) => void}
			{...props}
		/>
	);
}

function TabsList({ className, ...props }: React.ComponentProps<typeof TabsPrimitive.List>) {
	return (
		<TabsPrimitive.List
			data-slot="tabs-list"
			className={cn(
				"inline-flex items-center justify-center gap-3 rounded-base border-2 border-border bg-secondary-background px-2 pt-1.5 pb-2 text-foreground shadow-shadow",
				className,
			)}
			{...props}
		/>
	);
}

function TabsTrigger({ className, ...props }: React.ComponentProps<typeof TabsPrimitive.Trigger>) {
	return (
		<TabsPrimitive.Trigger
			data-slot="tabs-trigger"
			className={cn(
				"inline-flex cursor-pointer items-center justify-center gap-1.5 whitespace-nowrap rounded-base border-2 border-transparent px-3 py-2 font-heading text-sm transition-all hover:border-border focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:border-border data-[state=active]:bg-main data-[state=active]:text-main-foreground data-[state=active]:shadow-shadow data-[state=active]:hover:bg-main",
				className,
			)}
			{...props}
		/>
	);
}

function TabsContent({ className, ...props }: React.ComponentProps<typeof TabsPrimitive.Content>) {
	return (
		<TabsPrimitive.Content
			data-slot="tabs-content"
			className={cn(
				"mt-2 ring-offset-white focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
				className,
			)}
			{...props}
		/>
	);
}

export { Tabs, TabsList, TabsTrigger, TabsContent };
