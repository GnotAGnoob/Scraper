import { ChevronDown, ChevronUp } from "lucide-react";
import { useState } from "react";
import { cn } from "@/lib/utils";
import type { SharedIngredient } from "../../generated/kosik";
import { Button } from "../ui/button";
import { Card } from "../ui/card";
import { ScrollArea } from "../ui/scroll-area";
import { RecipeContent } from "./RecipeContent";

interface RecipeDisplayProps {
	recipe?: string;
	ingredients?: SharedIngredient[];
}

export const RecipeDisplay = ({ recipe, ingredients }: RecipeDisplayProps) => {
	const [isOpen, setIsOpen] = useState(false);

	return (
		<Card className={cn("gap-0 rounded-lg border-4 px-2 py-4", isOpen && "pb-0")}>
			<button
				type="button"
				className="flex w-full cursor-pointer items-center justify-between px-4"
				onClick={() => setIsOpen(!isOpen)}
			>
				<p className="font-bold font-mono text-xl">{"// RECIPE_&_INGREDIENTS"}</p>
				<div className="flex items-center gap-2 font-bold text-muted-foreground text-sm uppercase">
					{isOpen ? "Hide" : "Show"}
					<Button variant="neutral" size="icon" className="h-8 w-8">
						{isOpen ? <ChevronUp className="h-4 w-4" /> : <ChevronDown className="h-4 w-4" />}
					</Button>
				</div>
			</button>

			{isOpen && (
				<>
					<div className="mx-4 mt-4 border-border/50 border-t-2 border-dashed" />
					<ScrollArea className={cn("px-6", !!recipe && "h-[min(35rem,50vh)]")}>
						<div className="py-4">
							<RecipeContent recipe={recipe} ingredients={ingredients} />
						</div>
					</ScrollArea>
				</>
			)}
		</Card>
	);
};
