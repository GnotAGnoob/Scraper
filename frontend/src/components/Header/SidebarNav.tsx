import { Link, useRouter } from "@tanstack/react-router";
import { Home, type LucideIcon } from "lucide-react";
import type { FileRoutesByFullPath } from "../../routeTree.gen";

interface SidebarNavProps {
	onNavigate: () => void;
}

const routeIcons: Record<keyof FileRoutesByFullPath, LucideIcon> = {
	"/": Home,
};

const formatRouteName = (path: string): string => {
	if (path === "/") return "HOME";
	return path
		.slice(1)
		.split("/")
		.map((segment) => segment.toUpperCase())
		.join(" ");
};

export const SidebarNav = ({ onNavigate }: SidebarNavProps) => {
	const { routesByPath } = useRouter();
	return (
		<nav className="flex-1 p-4 overflow-y-auto space-y-2">
			{Object.keys(routesByPath).map((path) => {
				const pathKey = path as keyof typeof routesByPath;
				const Icon = routeIcons[pathKey] || Home;
				const routeName = formatRouteName(path);

				return (
					<Link
						key={pathKey}
						to={path}
						onClick={onNavigate}
						className="flex items-center gap-3 p-3 rounded-none border-2 border-transparent hover:border-border hover:bg-chart-2 hover:shadow-[4px_4px_0px_0px_var(--border)] transition-all font-bold"
						activeProps={{
							className:
								"flex items-center gap-3 p-3 rounded-none border-2 border-border bg-chart-2 shadow-[4px_4px_0px_0px_var(--border)] font-bold",
						}}
					>
						<Icon size={20} />
						<span>{routeName}</span>
					</Link>
				);
			})}
		</nav>
	);
};
