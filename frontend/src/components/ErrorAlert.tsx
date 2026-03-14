import type { SharedErrorResponse } from "../generated/kosik";
import { Alert, AlertDescription, AlertTitle } from "./ui/alert";

interface ErrorAlertProps {
	error: Error | null;
}

export const ErrorAlert = ({ error }: ErrorAlertProps) => {
	if (!error) return null;

	const message =
		(error as unknown as SharedErrorResponse)?.error ||
		(error instanceof Error ? error.message : "Something went wrong. Please try again.");

	return (
		<Alert variant="destructive">
			<AlertTitle>ERROR</AlertTitle>
			<AlertDescription>{message}</AlertDescription>
		</Alert>
	);
};
