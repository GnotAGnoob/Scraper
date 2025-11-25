import { Link } from "@tanstack/react-router";
import { Menu } from "lucide-react";
import { useState } from "react";
import { Overlay } from "../Overlay";
import { Button } from "../ui/button";
import { Sidebar } from "./Sidebar";

const Header = () => {
	const [isOpen, setIsOpen] = useState(false);

	return (
		<>
			<header className="p-4 flex items-center bg-secondary-background border-b-4 border-border shadow-[4px_4px_0px_0px_var(--border)] relative z-50">
				<Button variant="neutral" size="icon" onClick={() => setIsOpen(true)} aria-label="Open menu">
					<Menu size={24} />
				</Button>
				<h1 className="ml-4 text-2xl font-black italic tracking-tighter">
					<Link to="/">
						SCRAPE<span className="text-chart-1">MONKEY</span>
					</Link>
				</h1>
			</header>

			<Overlay isOpen={isOpen} onClose={() => setIsOpen(false)} />
			<Sidebar isOpen={isOpen} onClose={() => setIsOpen(false)} />
		</>
	);
};

export default Header;
