import type { SharedIngredient } from "../../generated/kosik";
import { Badge } from "../ui/badge";
import { Markdown } from "../ui/markdown";
import { SectionBadge } from "../ui/section-badge";

interface RecipeContentProps {
	recipe?: string;
	ingredients?: SharedIngredient[];
}

export const RecipeContent = ({ recipe, ingredients }: RecipeContentProps) => {
	return (
		<div className="space-y-6">
			{ingredients?.length && (
				<div className="space-y-2">
					<SectionBadge color="chart-2" size="sm" uppercase>
						Shopping List
					</SectionBadge>
					<div className="grid gap-x-4 gap-y-1 text-sm sm:grid-cols-2 md:grid-cols-3">
						{ingredients.map((ing, idx) => (
							<div
								key={`${ing.originalName || ing.name}-${idx}`}
								className="flex items-center justify-between border-border/10 border-b py-1 last:border-0"
							>
								<span className="mr-2 truncate font-medium" title={ing.originalName || ing.name}>
									{ing.originalName || ing.name}
								</span>
								{ing.quantityGrams && (
									<Badge variant="neutral" className="px-1 py-0.5 font-mono text-2xs">
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
					<SectionBadge color="chart-2" size="sm" uppercase>
						Instructions
					</SectionBadge>
					<Markdown>{recipe}</Markdown>
				</div>
			)}
		</div>
	);
};
