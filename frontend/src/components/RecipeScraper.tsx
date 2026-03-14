import { ReceiptText, Search } from "lucide-react";
import { useState } from "react";
import { MonkeyIcon } from "./MonkeyIcon";
import { RecipeTab } from "./RecipeTab";
import { SearchTab } from "./SearchTab";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "./ui/tabs";

type Mode = "recipe" | "search";

const SUBTITLE: Record<Mode, string> = {
	recipe: "Paste your recipe below and get shopping!",
	search: "Search for products directly on Kosik.cz",
};

export const RecipeScraper = () => {
	const [mode, setMode] = useState<Mode>("recipe");

	return (
		<div className="container mx-auto max-w-[1600px] space-y-6 px-4 py-8">
			<div className="text-center">
				<div className="flex justify-center">
					<div className="-rotate-1 flex transform items-center justify-center gap-2 border-4 border-border bg-main px-6 py-1 pb-2.5 shadow-[8px_8px_0px_0px_var(--border)]">
						<MonkeyIcon className="mt-2" size={56} />
						<h1 className="font-black text-4xl md:text-5xl">Scrape Monkey</h1>
					</div>
				</div>
				<p className="mt-6 font-bold text-muted-foreground text-xl md:text-2xl">{SUBTITLE[mode]}</p>
			</div>

			<Tabs value={mode} onValueChange={setMode} className="space-y-6">
				<TabsList className="mx-auto flex w-full max-w-6xl">
					<TabsTrigger value="recipe" className="flex-1 gap-2 font-black">
						<ReceiptText className="h-4 w-4" />
						RECIPE MODE
					</TabsTrigger>
					<TabsTrigger value="search" className="flex-1 gap-2 font-black">
						<Search className="h-4 w-4" />
						DIRECT SEARCH
					</TabsTrigger>
				</TabsList>

				<TabsContent value="recipe">
					<RecipeTab />
				</TabsContent>

				<TabsContent value="search">
					<SearchTab />
				</TabsContent>
			</Tabs>
		</div>
	);
};
