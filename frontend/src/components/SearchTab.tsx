import { useSearchMutation } from "../hooks/useSearchMutation";
import { ErrorAlert } from "./ErrorAlert";
import { ProductGrid } from "./Product/ProductGrid";
import { SearchForm } from "./SearchForm";
import { Card } from "./ui/card";

export const SearchTab = () => {
	const mutation = useSearchMutation();
	const data = mutation.isSuccess ? mutation.data?.data : undefined;

	return (
		<div className="space-y-6">
			<div className="mx-auto max-w-6xl space-y-6">
				<Card className="rounded-lg border-4 px-6 pt-4 pb-6">
					<SearchForm mutation={mutation} />
				</Card>

				{mutation.isError && <ErrorAlert error={mutation.error} />}
			</div>

			{data?.ingredientsProducts && (
				<div className="fade-in slide-in-from-bottom-8 animate-in space-y-4 duration-700">
					<ProductGrid results={data.ingredientsProducts} />
				</div>
			)}
		</div>
	);
};
