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
			className={`fixed top-0 left-0 h-full w-80 bg-background border-r-4 border-border shadow-[4px_0_0_0_var(--border)] z-50 transform transition-transform duration-300 ease-in-out flex flex-col ${
				isOpen ? "translate-x-0" : "-translate-x-full"
			}`}
		>
			<div className="flex items-center justify-between p-4 border-b-4 border-border bg-secondary-background">
				<h2 className="text-xl font-black">MENU</h2>
				<Button
					variant="neutral"
					size="icon"
					onClick={onClose}
					className="bg-chart-3 border-2 border-transparent hover:border-border transition-colors"
					aria-label="Close menu"
				>
					<X size={24} />
				</Button>
			</div>

			<SidebarNav onNavigate={onClose} />
		</aside>
	);
};
