import { ChevronDown, ChevronUp } from "lucide-react";
import { useState } from "react";
import ReactMarkdown from "react-markdown";
import { cn } from "@/lib/utils";
import type { SharedIngredient } from "../generated/kosik";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";
import { Card } from "./ui/card";
import { ScrollArea } from "./ui/scroll-area";

interface RecipeDisplayProps {
	recipe?: string;
	ingredients?: SharedIngredient[];
}

export const RecipeDisplay = ({ recipe, ingredients }: RecipeDisplayProps) => {
	const [isOpen, setIsOpen] = useState(true);

	return (
		<Card className={cn("rounded-lg border-4 px-6 py-4", isOpen && "pb-0")}>
			<button
				type="button"
				className="flex w-full cursor-pointer items-center justify-between"
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
				<div className="mt-4 border-border/50 border-t-2 border-dashed pt-4">
					<ScrollArea className="h-[500px] pr-4">
						<div className="space-y-8 pb-4">
							{ingredients?.length && (
								<div className="space-y-3">
									<h3 className="inline-block border-2 border-border bg-chart-2 px-2 py-0.5 font-black text-base uppercase shadow-shadow">
										Shopping List
									</h3>
									<div className="grid gap-x-4 gap-y-1 text-sm sm:grid-cols-2 md:grid-cols-3">
										{ingredients.map((ing, idx) => (
											<div
												key={`${ing.originalName || ing.name}-${idx}`}
												className="flex items-center justify-between border-border/10 border-b py-1 last:border-0"
											>
												<span
													className="mr-2 truncate font-medium"
													title={ing.originalName || ing.name}
												>
													{ing.originalName || ing.name}
												</span>
												{ing.quantityGrams && (
													<Badge
														variant="neutral"
														className="h-5 px-1.5 font-mono text-[10px]"
													>
														{ing.quantityGrams}g
													</Badge>
												)}
											</div>
										))}
									</div>
								</div>
							)}

							{recipe && (
								<div className="space-y-3">
									<h3 className="inline-block border-2 border-border bg-chart-1 px-2 py-0.5 font-black text-base uppercase shadow-shadow">
										Instructions
									</h3>
									<div className="prose prose-neutral prose-sm prose-li:my-0 max-w-none prose-headings:font-black prose-p:leading-snug">
										<ReactMarkdown>{recipe}</ReactMarkdown>
									</div>
								</div>
							)}
						</div>
					</ScrollArea>
				</div>
			)}
		</Card>
	);
};
