import { useForm } from "@tanstack/react-form";
import type { SearchMutation } from "../hooks/useSearchMutation";
import { createQueryRow, SearchInputRows } from "./SearchInputRows";
import { SubmitButton } from "./SubmitButton";

interface SearchFormProps {
	mutation: SearchMutation;
}

export const SearchForm = ({ mutation }: SearchFormProps) => {
	const form = useForm({
		defaultValues: {
			rows: [createQueryRow()],
		},
		onSubmit: async ({ value }) => {
			const queries = value.rows.map((r) => r.value.trim()).filter(Boolean);
			if (queries.length > 0) {
				await mutation.mutateAsync(queries);
			}
		},
	});

	return (
		<form
			onSubmit={(e) => {
				e.preventDefault();
				form.handleSubmit();
			}}
		>
			<div className="space-y-2">
				<p className="font-bold font-mono text-xl">{"// SEARCH_QUERIES"}</p>
				<form.Field name="rows">
					{(field) => (
						<SearchInputRows
							rows={field.state.value}
							onChange={field.handleChange}
							disabled={mutation.isPending}
							onSubmit={form.handleSubmit}
						/>
					)}
				</form.Field>

				<div className="flex justify-end">
					<form.Subscribe selector={(state) => [state.values.rows, state.isSubmitting] as const}>
						{([rows, isSubmitting]) => (
							<SubmitButton
								isPending={isSubmitting}
								disabled={rows.every((r) => !r.value.trim())}
								label="SEARCH PRODUCTS ->"
							/>
						)}
					</form.Subscribe>
				</div>
			</div>
		</form>
	);
};
