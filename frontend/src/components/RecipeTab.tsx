import { useRecipeMutation } from "../hooks/useRecipeMutation";
import { ErrorAlert } from "./ErrorAlert";
import { ProductGrid } from "./Product/ProductGrid";
import { RecipeDisplay } from "./RecipeDisplay";
import { RecipeForm } from "./RecipeForm";
import { Card } from "./ui/card";

export const RecipeTab = () => {
	const mutation = useRecipeMutation();
	const data = mutation.isSuccess ? mutation.data?.data : undefined;

	return (
		<div className="space-y-6">
			<div className="mx-auto max-w-6xl space-y-6">
				<Card className="rounded-lg border-4 px-6 pt-4 pb-6">
					<RecipeForm mutation={mutation} />
				</Card>

				{mutation.isError && <ErrorAlert error={mutation.error} />}

				{data && (data.recipe || data.ingredientsProducts) && (
					<div className="fade-in slide-in-from-bottom-8 animate-in duration-700">
						<RecipeDisplay
							recipe={data.recipe}
							ingredients={data.ingredientsProducts?.map((ip) => ip.ingredient).filter((i) => !!i)}
						/>
					</div>
				)}
			</div>

			{data?.ingredientsProducts && (
				<div className="fade-in slide-in-from-bottom-8 animate-in space-y-4 duration-700">
					<ProductGrid results={data.ingredientsProducts} />
				</div>
			)}
		</div>
	);
};
