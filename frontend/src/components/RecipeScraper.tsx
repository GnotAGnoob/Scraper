import { useMutation } from "@tanstack/react-query";
import { Loader2 } from "lucide-react";
import { useState } from "react";
import type { SharedErrorResponse } from "../generated/kosik";
import { api } from "../lib/api";
import { InputArea } from "./InputArea";
import { ProductGrid } from "./ProductGrid";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";
import { Button } from "./ui/button";

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
		<div className="container mx-auto px-4 py-12 max-w-7xl space-y-12">
			<div className="text-center space-y-4">
				<h1 className="text-5xl md:text-7xl font-black bg-chart-1 inline-block px-6 py-2 border-4 border-border shadow-[8px_8px_0px_0px_var(--border)] -rotate-1 transform">
					KOSIK SCRAPER
				</h1>
				<p className="text-xl md:text-2xl font-bold text-muted-foreground mt-4">
					Paste your recipe below and get shopping!
				</p>
			</div>

			<div className="space-y-6 max-w-4xl mx-auto">
				<div className="bg-card p-6 border-4 border-border shadow-[8px_8px_0px_0px_var(--border)] rounded-lg">
					<p className="block text-xl font-bold mb-4 font-mono">{"// RECIPE_INPUT"}</p>
					<InputArea
						value={text}
						onChange={setText}
						placeholder="Paste your ingredients list or full recipe here... e.g., '200g chicken breast, 1 onion, 500ml milk'"
						disabled={mutation.isPending}
					/>

					<div className="mt-6 flex justify-end">
						<Button
							onClick={handleSubmit}
							disabled={mutation.isPending || !text.trim()}
							className="text-xl px-8 py-6 font-black border-2 border-border shadow-[4px_4px_0px_0px_var(--border)] hover:shadow-[6px_6px_0px_0px_var(--border)] hover:-translate-y-1 hover:-translate-x-1 active:shadow-none active:translate-y-[4px] active:translate-x-[4px] transition-all bg-chart-4 hover:bg-chart-4/90 text-main-foreground"
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
				</div>

				{mutation.isError && (
					<Alert
						variant="destructive"
						className="border-4 border-border bg-chart-3/20 shadow-[4px_4px_0px_0px_var(--border)]"
					>
						<AlertTitle className="text-xl font-black">ERROR</AlertTitle>
						<AlertDescription className="font-mono text-lg">{errorMessage}</AlertDescription>
					</Alert>
				)}
			</div>

			{mutation.isSuccess && mutation.data.data?.ingredientsProducts && (
				<div className="animate-in fade-in slide-in-from-bottom-8 duration-700">
					<div className="mb-8 border-b-4 border-border pb-4">
						<h2 className="text-4xl font-black uppercase">Results</h2>
					</div>
					<ProductGrid results={mutation.data.data.ingredientsProducts} />
				</div>
			)}
		</div>
	);
}
