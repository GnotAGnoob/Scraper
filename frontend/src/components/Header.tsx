import { Link } from "@tanstack/react-router";
import { Home, Menu, X } from "lucide-react";
import { useState } from "react";
import { Button } from "./ui/button";

export default function Header() {
	const [isOpen, setIsOpen] = useState(false);

	return (
		<>
			<header className="p-4 flex items-center bg-white border-b-4 border-black shadow-[4px_4px_0px_0px_var(--border)] relative z-50">
				<Button
					variant="neutral"
					size="icon"
					onClick={() => setIsOpen(true)}
					className="border-2 border-black shadow-[2px_2px_0px_0px_rgba(0,0,0,1)] hover:translate-x-px hover:translate-y-px hover:shadow-none transition-all"
					aria-label="Open menu"
				>
					<Menu size={24} />
				</Button>
				<h1 className="ml-4 text-2xl font-black italic tracking-tighter">
					<Link to="/">
						KOSIK<span className="text-chart-1">SCRAPER</span>
					</Link>
				</h1>
			</header>

			{/* Overlay */}
			{isOpen && (
				<button
					type="button"
					className="fixed inset-0 bg-black/50 z-40 w-full h-full cursor-default"
					onClick={() => setIsOpen(false)}
					onKeyDown={(e) => {
						if (e.key === "Escape") setIsOpen(false);
					}}
					aria-label="Close menu"
				/>
			)}

			<aside
				className={`fixed top-0 left-0 h-full w-80 bg-background border-r-4 border-black shadow-[4px_0_0_0_var(--border)] z-50 transform transition-transform duration-300 ease-in-out flex flex-col ${
					isOpen ? "translate-x-0" : "-translate-x-full"
				}`}
			>
				<div className="flex items-center justify-between p-4 border-b-4 border-black bg-secondary-background">
					<h2 className="text-xl font-black">MENU</h2>
					<Button
						variant="neutral"
						size="icon"
						onClick={() => setIsOpen(false)}
						className="hover:bg-destructive hover:text-white border-2 border-transparent hover:border-black transition-colors"
						aria-label="Close menu"
					>
						<X size={24} />
					</Button>
				</div>

				<nav className="flex-1 p-4 overflow-y-auto space-y-2">
					<Link
						to="/"
						onClick={() => setIsOpen(false)}
						className="flex items-center gap-3 p-3 rounded-none border-2 border-transparent hover:border-black hover:bg-chart-2 hover:shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] transition-all font-bold"
						activeProps={{
							className:
								"flex items-center gap-3 p-3 rounded-none border-2 border-black bg-chart-2 shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] font-bold",
						}}
					>
						<Home size={20} />
						<span>HOME</span>
					</Link>
				</nav>
			</aside>
		</>
	);
}
