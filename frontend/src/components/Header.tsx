import { Link } from "@tanstack/react-router";
import { Home, Menu, X } from "lucide-react";
import { useState } from "react";
import { Button } from "./ui/button";

export default function Header() {
	const [isOpen, setIsOpen] = useState(false);

	return (
		<>
			<header className="sticky top-0 z-40 w-full border-b-2 border-border bg-secondary-background shadow-shadow">
				<div className="container mx-auto px-4 py-4 flex items-center justify-between">
					<Link to="/" className="flex items-center gap-2">
						<h1 className="text-2xl font-heading text-main">Recipe to Kosik</h1>
					</Link>
					<Button
						variant="neutral"
						size="icon"
						onClick={() => setIsOpen(true)}
						aria-label="Open menu"
						className="md:hidden"
					>
						<Menu size={24} />
					</Button>
					<nav className="hidden md:flex items-center gap-4">
						<Link to="/">
							<Button variant="neutral" size="sm">
								<Home size={18} className="mr-2" />
								Home
							</Button>
						</Link>
					</nav>
				</div>
			</header>

			<aside
				className={`fixed top-0 left-0 h-full w-80 bg-secondary-background border-r-2 border-border shadow-shadow z-50 transform transition-transform duration-300 ease-in-out flex flex-col ${
					isOpen ? "translate-x-0" : "-translate-x-full"
				}`}
			>
				<div className="flex items-center justify-between p-4 border-b-2 border-border">
					<h2 className="text-xl font-heading">Navigation</h2>
					<Button variant="neutral" size="icon" onClick={() => setIsOpen(false)} aria-label="Close menu">
						<X size={24} />
					</Button>
				</div>

				<nav className="flex-1 p-4 overflow-y-auto">
					<Link
						to="/"
						onClick={() => setIsOpen(false)}
						className="flex items-center gap-3 p-3 rounded-base hover:bg-main hover:text-main-foreground transition-colors mb-2 border-2 border-transparent hover:border-border"
						activeProps={{
							className:
								"flex items-center gap-3 p-3 rounded-base bg-main text-main-foreground border-2 border-border shadow-shadow mb-2",
						}}
					>
						<Home size={20} />
						<span className="font-base">Home</span>
					</Link>
				</nav>
			</aside>

			{isOpen && (
				<div
					className="fixed inset-0 bg-overlay z-40 md:hidden"
					onClick={() => setIsOpen(false)}
					aria-hidden="true"
				/>
			)}
		</>
	);
}
