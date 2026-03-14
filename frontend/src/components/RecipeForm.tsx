import { useForm } from "@tanstack/react-form";
import type { RecipeMutation } from "../hooks/useRecipeMutation";
import { InputArea } from "./InputArea";
import { SubmitButton } from "./SubmitButton";

interface RecipeFormProps {
	mutation: RecipeMutation;
}

export const RecipeForm = ({ mutation }: RecipeFormProps) => {
	const form = useForm({
		defaultValues: {
			text: "",
		},
		onSubmit: async ({ value }) => {
			await mutation.mutateAsync(value.text);
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
				<p className="font-bold font-mono text-xl">{"// RECIPE_INPUT"}</p>
				<form.Field name="text">
					{(field) => (
						<InputArea
							value={field.state.value}
							onChange={(v) => field.handleChange(v)}
							placeholder="Paste your ingredients list or full recipe here... e.g., '200g chicken breast, 1 onion, 500ml milk'"
							disabled={mutation.isPending}
							onSubmit={() => form.handleSubmit()}
						/>
					)}
				</form.Field>

				<div className="flex justify-end">
					<form.Subscribe selector={(state) => [state.values.text, state.isSubmitting] as const}>
						{([text, isSubmitting]) => (
							<SubmitButton
								isPending={isSubmitting}
								disabled={!text.trim()}
								label="FIND INGREDIENTS ->"
							/>
						)}
					</form.Subscribe>
				</div>
			</div>
		</form>
	);
};
