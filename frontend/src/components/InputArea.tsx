import { Textarea } from "./ui/textarea";

interface InputAreaProps {
	value: string;
	onChange: (value: string) => void;
	placeholder?: string;
	disabled?: boolean;
}

export const InputArea = ({ value, onChange, placeholder, disabled }: InputAreaProps) => {
	return (
		<div className="w-full">
			<Textarea
				value={value}
				onChange={(e) => onChange(e.target.value)}
				placeholder={placeholder}
				disabled={disabled}
				className="min-h-[200px] p-4 font-mono text-lg"
			/>
		</div>
	);
};
