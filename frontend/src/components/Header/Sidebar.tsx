import { X } from "lucide-react";
import { Button } from "../ui/button";
import { SidebarNav } from "./SidebarNav";

interface SidebarProps {
	isOpen: boolean;
	onClose: () => void;
}

export const Sidebar = ({ isOpen, onClose }: SidebarProps) => {
	return (
		<aside
			className={`fixed top-0 left-0 z-50 flex h-full w-80 transform flex-col border-border border-r-4 bg-background shadow-[4px_0_0_0_var(--border)] transition-transform duration-300 ease-in-out ${
				isOpen ? "translate-x-0" : "-translate-x-full"
			}`}
		>
			<div className="flex items-center justify-between border-border border-b-4 bg-secondary-background p-4">
				<h2 className="font-black text-xl">MENU</h2>
				<Button
					variant="neutral"
					size="icon"
					onClick={onClose}
					className="border-2 border-transparent bg-chart-3 transition-colors hover:border-border"
					aria-label="Close menu"
				>
					<X size={24} />
				</Button>
			</div>

			<SidebarNav onNavigate={onClose} />
		</aside>
	);
};
