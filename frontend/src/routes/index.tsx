import { createFileRoute } from "@tanstack/react-router";
import { RecipeScraper } from "@/components/RecipeScraper";

const App = () => {
	return (
		<div className="min-h-screen bg-background pt-[5vh] font-sans">
			<RecipeScraper />
		</div>
	);
};

export const Route = createFileRoute("/")({
	component: App,
});
