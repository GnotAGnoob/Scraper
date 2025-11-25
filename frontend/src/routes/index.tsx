import { useMutation } from "@tanstack/react-query";
import { createFileRoute } from "@tanstack/react-router";
import { ChefHat, Sparkles } from "lucide-react";
import { useCallback, useState } from "react";
import { ProductGrid } from "@/components/ProductGrid";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import type { OrchestratorIngredientProducts } from "@/generated/kosik";
import { recipeApi } from "@/lib/api";

export const Route = createFileRoute("/")({
	component: HomePage,
});

const EXAMPLE_RECIPE = `Chocolate Chip Cookies

Ingredients:
- 200g butter
- 150g sugar
- 2 eggs
- 300g flour
- 200g chocolate chips
- 1 tsp vanilla extract
- 1 tsp baking powder
- pinch of salt`;

function HomePage() {
	const [recipeText, setRecipeText] = useState("");
	const [products, setProducts] = useState<OrchestratorIngredientProducts[]>([]);

	const mutation = useMutation({
		mutationFn: (text: string) => recipeApi.parseRecipe(text),
		onSuccess: (data) => {
			setProducts(data);
		},
	});

	const handleSubmit = useCallback(
		(e: React.FormEvent) => {
			e.preventDefault();
			if (recipeText.trim()) {
				mutation.mutate(recipeText);
			}
		},
		[recipeText, mutation],
	);

	const handleTryExample = useCallback(() => {
		setRecipeText(EXAMPLE_RECIPE);
	}, []);

	return (
		<div className="container mx-auto px-4 py-8">
			<div className="mb-12 text-center">
				<h1 className="mb-4 text-5xl font-heading">
					Turn Recipes into{" "}
					<span className="inline-block rounded-[--radius-base] border-4 border-border bg-main px-4 py-2 text-main-foreground shadow-shadow">
						Kosik Cart
					</span>
				</h1>
				<p className="text-lg text-foreground/70">
					Paste your recipe and we'll find matching products on Kosik.cz
				</p>
			</div>

			<form onSubmit={handleSubmit} className="mb-12">
				<div className="mb-4">
					<label htmlFor="recipe" className="mb-2 block text-lg font-heading">
						Your Recipe
					</label>
					<Textarea
						key={recipeText}
						value={recipeText}
						onChange={(e) => setRecipeText(e.target.value)}
						placeholder="Paste your recipe here... Include ingredients with quantities!"
						rows={12}
						className="resize-none text-base"
					/>
				</div>

				<div className="flex flex-wrap gap-4">
					<Button
						type="submit"
						disabled={mutation.isPending || !recipeText.trim()}
						className="group relative overflow-hidden rounded-[--radius-base] border-4 border-border bg-main px-8 py-6 text-lg font-heading text-main-foreground shadow-shadow transition-all hover:shadow-[8px_8px_0px_0px_var(--border)] hover:translate-x-[-2px] hover:translate-y-[-2px] disabled:opacity-50 disabled:hover:translate-x-0 disabled:hover:translate-y-0 disabled:hover:shadow-shadow"
					>
						{mutation.isPending ? (
							<>
								<Sparkles className="mr-2 inline h-5 w-5 animate-spin" />
								Finding Products...
							</>
						) : (
							<>
								<ChefHat className="mr-2 inline h-5 w-5" />
								Find Products
							</>
						)}
					</Button>

					<Button
						type="button"
						onClick={handleTryExample}
						className="rounded-[--radius-base] border-4 border-border bg-secondary-background px-6 py-6 text-lg font-heading text-foreground shadow-shadow transition-all hover:shadow-[6px_6px_0px_0px_var(--border)] hover:translate-x-[-2px] hover:translate-y-[-2px]"
					>
						Try Example
					</Button>
				</div>

				{mutation.isError && (
					<div className="mt-4 rounded-[--radius-base] border-4 border-border bg-chart-3 p-4 font-heading text-main-foreground shadow-shadow">
						Error: {mutation.error instanceof Error ? mutation.error.message : "Failed to parse recipe"}
					</div>
				)}
			</form>

			{(mutation.isPending || products.length > 0 || mutation.isError) && (
				<div>
					<h2 className="mb-6 text-3xl font-heading">
						{mutation.isPending ? "Finding Products..." : `Found ${products.length} Products`}
					</h2>
					<ProductGrid products={products} isLoading={mutation.isPending} />
				</div>
			)}

			{!mutation.isPending && !products.length && !mutation.isError && recipeText && (
				<div className="text-center">
					<p className="text-lg text-foreground/60">Click "Find Products" to start searching</p>
				</div>
			)}
		</div>
	);
}
