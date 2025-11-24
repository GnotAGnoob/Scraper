import { Textarea } from "./ui/textarea";

interface InputAreaProps {
	value: string;
	onChange: (value: string) => void;
	placeholder?: string;
	disabled?: boolean;
}

export function InputArea({ value, onChange, placeholder, disabled }: InputAreaProps) {
	return (
		<div className="w-full">
			<Textarea
				value={value}
				onChange={(e) => onChange(e.target.value)}
				placeholder={placeholder}
				disabled={disabled}
				className="min-h-[200px] text-lg p-4 border-2 border-border shadow-shadow focus-visible:ring-offset-2 font-mono"
			/>
		</div>
	);
}
