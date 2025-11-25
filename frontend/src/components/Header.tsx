import { Link } from "@tanstack/react-router";
import { Github, ScrollText } from "lucide-react";
import { useMemo } from "react";

import { Button } from "@/components/ui/button";

const Header = () => {
	const swaggerUrl = useMemo(() => {
		try {
			const base = import.meta.env.VITE_API_URL ?? "http://localhost:8080";
			return new URL("/swagger/index.html", base).toString();
		} catch {
			return "/swagger/index.html";
		}
	}, []);

	return (
		<header className="sticky top-0 z-50 border-b-2 border-border bg-secondary-background shadow-shadow">
			<div className="mx-auto flex max-w-6xl items-center justify-between gap-4 px-4 py-4">
				<Link to="/" className="flex items-center gap-4">
					<div className="flex size-12 items-center justify-center rounded-full border-2 border-border bg-main text-xl font-heading uppercase shadow-shadow">
						KS
					</div>
					<div className="leading-snug">
						<p className="font-heading text-2xl uppercase">Kosik Cart Pilot</p>
						<p className="text-xs font-semibold uppercase tracking-[0.2em] text-foreground/60">
							LLM ingredients → screened products
						</p>
					</div>
				</Link>
				<div className="flex flex-wrap items-center gap-3">
					<Button asChild variant="neutral" size="sm">
						<a href={swaggerUrl} target="_blank" rel="noreferrer">
							<ScrollText className="size-4" />
							Swagger
						</a>
					</Button>
					<Button asChild size="sm">
						<a href="https://github.com/GnotAGnoob/kosik-scraper" target="_blank" rel="noreferrer">
							<Github className="size-4" />
							Repo
						</a>
					</Button>
				</div>
			</div>
		</header>
	);
};

export default Header;
