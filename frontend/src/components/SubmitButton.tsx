import { Loader2 } from "lucide-react";
import { Button } from "./ui/button";

interface SubmitButtonProps {
	isPending: boolean;
	disabled: boolean;
	label: string;
}

export const SubmitButton = ({ isPending, disabled, label }: SubmitButtonProps) => {
	return (
		<Button type="submit" disabled={isPending || disabled} className="px-8 py-6 font-black text-xl">
			{isPending ? (
				<>
					<Loader2 className="mr-2 h-6 w-6 animate-spin" />
					PROCESSING...
				</>
			) : (
				label
			)}
		</Button>
	);
};
