import { Textarea } from "./ui/textarea";

interface InputAreaProps {
	value: string;
	onChange: (value: string) => void;
	placeholder?: string;
	disabled?: boolean;
	onSubmit?: () => void;
}

export const InputArea = ({ value, onChange, placeholder, disabled, onSubmit }: InputAreaProps) => {
	const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
		if (e.key === "Enter" && !e.shiftKey && onSubmit) {
			e.preventDefault();
			onSubmit();
		}
	};

	return (
		<div className="w-full">
			<Textarea
				value={value}
				onChange={(e) => onChange(e.target.value)}
				onKeyDown={handleKeyDown}
				placeholder={placeholder}
				disabled={disabled}
				className="min-h-60 p-4 font-mono text-lg"
			/>
		</div>
	);
};
