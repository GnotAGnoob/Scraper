import { createFileRoute } from "@tanstack/react-router";
import { RecipeScraper } from "@/components/RecipeScraper";

const App = () => {
	return (
		<div className="min-h-screen bg-background font-sans">
			<RecipeScraper />
		</div>
	);
};

export const Route = createFileRoute("/")({
	component: App,
});
