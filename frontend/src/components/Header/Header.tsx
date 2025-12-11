import { ThemeSwitcher } from "@/components/ui/theme-switcher";

export const Header = () => {
	return (
		<header className="flex justify-end p-4">
			<ThemeSwitcher />
		</header>
	);
};
