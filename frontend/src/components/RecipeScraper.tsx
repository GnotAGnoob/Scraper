import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import { useState } from "react";
import type { SharedErrorResponse } from "../generated/kosik";
import { api } from "../lib/api";
import { InputArea } from "./InputArea";
import { MonkeyIcon } from "./MonkeyIcon";
import { ProductGrid } from "./Product/ProductGrid";
import { RecipeDisplay } from "./RecipeDisplay";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";
import { Button } from "./ui/button";
import { Card } from "./ui/card";

export const RecipeScraper = () => {
	const [text, setText] = useState("");

	const mutation = useMutation({
		mutationFn: (text: string) => api.text.v1TextProductsCreate({ text }),
	});

	const handleSubmit = () => {
		if (!text.trim()) return;

		mutation.mutate(text);
	};

	const errorMessage =
		(mutation.error as SharedErrorResponse)?.error ||
		(mutation.error instanceof Error ? mutation.error.message : "Something went wrong. Please try again.");

	return (
		<div className="container mx-auto max-w-[1600px] space-y-6 px-4 py-8">
			<div className="text-center">
				<div className="flex justify-center">
					<div className="-rotate-1 flex transform items-center justify-center gap-2 border-4 border-border bg-chart-1 px-6 py-1 pb-2.5 shadow-[8px_8px_0px_0px_var(--border)]">
						<MonkeyIcon className="mt-2" size={56} />
						<h1 className="font-black text-4xl md:text-5xl">Scrape Monkey</h1>
					</div>
				</div>
				<p className="mt-6 font-bold text-muted-foreground text-xl md:text-2xl">
					Paste your recipe below and get shopping!
				</p>
			</div>

			<div className="mx-auto max-w-6xl space-y-6">
				<Card className="rounded-lg border-4 px-6 pt-4 pb-6">
					<p className="font-bold font-mono text-xl">{"// RECIPE_INPUT"}</p>
					<InputArea
						value={text}
						onChange={setText}
						placeholder="Paste your ingredients list or full recipe here... e.g., '200g chicken breast, 1 onion, 500ml milk'"
						disabled={mutation.isPending}
					/>

					<div className="mt-2 flex justify-end">
						<Button
							onClick={handleSubmit}
							disabled={mutation.isPending || !text.trim()}
							className="bg-chart-4 px-8 py-6 font-black text-xl"
						>
							{mutation.isPending ? (
								<>
									<Loader2 className="mr-2 h-6 w-6 animate-spin" />
									PROCESSING...
								</>
							) : (
								"FIND INGREDIENTS ->"
							)}
						</Button>
					</div>
				</Card>

				{mutation.isError && (
					<Alert className="border-4 bg-chart-3/40">
						<AlertTitle>ERROR</AlertTitle>
						<AlertDescription>{errorMessage}</AlertDescription>
					</Alert>
				)}

				{mutation.isSuccess && (mutation.data?.data.recipe || mutation.data?.data.ingredientsProducts) && (
					<div className="fade-in slide-in-from-bottom-8 animate-in duration-700">
						<RecipeDisplay
							recipe={mutation.data.data.recipe}
							ingredients={mutation.data.data.ingredientsProducts
								?.map((ip) => ip.ingredient)
								.filter((i) => !!i)}
						/>
					</div>
				)}
			</div>

			{mutation.isSuccess && mutation.data.data?.ingredientsProducts && (
				<div className="fade-in slide-in-from-bottom-8 animate-in space-y-4 duration-700">
					<div className="border-border border-b-4 pb-2">
						<h2 className="font-black text-3xl uppercase">Results</h2>
					</div>
					<ProductGrid results={mutation.data.data.ingredientsProducts} />
				</div>
			)}
		</div>
	);
};
