import { X } from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "./ui/button";
import { Input } from "./ui/input";

export interface QueryRow {
	id: string;
	value: string;
}

interface SearchInputRowsProps {
	rows: QueryRow[];
	disabled?: boolean;
	onChange: (rows: QueryRow[]) => void;
	onSubmit?: () => void;
}

const MAX_ROWS = 20;
export const createQueryRow = (value = ""): QueryRow => ({ id: crypto.randomUUID(), value });

export const SearchInputRows = ({ rows, onChange, disabled, onSubmit }: SearchInputRowsProps) => {
	const updateRow = (index: number, value: string) => {
		const updated = [...rows];
		updated[index] = { ...updated[index], value };

		// Auto-add a trailing empty row when the user types in the last row
		const isLastRow = index === updated.length - 1;
		if (isLastRow && value.trim() !== "" && updated.length < MAX_ROWS) {
			updated.push(createQueryRow());
		}

		onChange(updated);
	};

	const removeRow = (index: number) => {
		// Always keep at least one row
		if (rows.length <= 1) return;
		const updated = rows.filter((_, i) => i !== index);

		// Ensure there's always a trailing empty row for new input
		const lastHasContent = updated.length > 0 && updated[updated.length - 1].value.trim() !== "";
		if (lastHasContent && updated.length < MAX_ROWS) {
			updated.push(createQueryRow());
		}

		onChange(updated);
	};

	const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>, index: number) => {
		if (e.key === "Enter") {
			e.preventDefault();
			const hasContent = rows.some((r) => r.value.trim() !== "");
			if (hasContent) {
				onSubmit?.();
			}
		} else if (e.key === "Backspace" && rows[index].value === "" && rows.length > 1) {
			e.preventDefault();
			removeRow(index);
		}
	};

	const handlePaste = (e: React.ClipboardEvent<HTMLInputElement>, index: number) => {
		const pastedText = e.clipboardData.getData("text");
		const lines = pastedText.split(/\r?\n/).filter((line) => line.trim() !== "");

		if (lines.length > 1) {
			e.preventDefault();
			const updated = [...rows];
			updated[index] = { ...updated[index], value: lines[0] };

			const newRows = lines.slice(1, MAX_ROWS - updated.length + 1).map((line) => createQueryRow(line));

			updated.splice(index + 1, 0, ...newRows);
			const trimmed = updated.slice(0, MAX_ROWS);

			// Ensure trailing empty row
			const lastHasContent = trimmed.length > 0 && trimmed[trimmed.length - 1].value.trim() !== "";
			if (lastHasContent && trimmed.length < MAX_ROWS) {
				trimmed.push(createQueryRow());
			}

			onChange(trimmed);
		}
	};

	// Count non-empty rows for the display
	const filledCount = rows.filter((r) => r.value.trim() !== "").length;

	return (
		<div className="space-y-2">
			{rows.map((row, index) => (
				<div key={row.id} className="flex items-center gap-3">
					<div className="flex flex-1 items-center gap-2">
						<span className="shrink-0 text-center font-mono text-lg text-muted-foreground">
							{index + 1}.
						</span>

						<Input
							value={row.value}
							onChange={(e) => updateRow(index, e.target.value)}
							onKeyDown={(e) => handleKeyDown(e, index)}
							onPaste={(e) => handlePaste(e, index)}
							placeholder={index === 0 ? "Search term or kosik.cz URL..." : "Add another search..."}
							disabled={disabled}
							className="h-12 font-mono text-base"
						/>
					</div>

					{rows.length > 1 && (
						<Button
							type="button"
							variant="neutral"
							size="icon"
							onClick={() => removeRow(index)}
							disabled={disabled}
							className={cn("shrink-0", index === rows.length - 1 ? "opacity-0" : "")}
							aria-label={`Remove search ${index + 1}`}
						>
							<X />
						</Button>
					)}
				</div>
			))}

			{filledCount > 0 && (
				<p className="text-right font-mono text-muted-foreground text-sm">
					{filledCount} {filledCount === 1 ? "query" : "queries"} / {MAX_ROWS} max
				</p>
			)}
		</div>
	);
};
