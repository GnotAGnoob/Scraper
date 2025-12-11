import { Moon, Sun } from "lucide-react";
import { useCallback, useEffect, useState } from "react";
import { Button } from "@/components/ui/button";
import { cn } from "@/lib/utils";

const DARK_MODE_STORAGE_KEY = "kosik-dark-mode";

export const ThemeSwitcher = () => {
	const [isDark, setIsDark] = useState(false);

	const handleThemeChange = useCallback((isDark: boolean) => {
		if (isDark) {
			document.documentElement.classList.add("dark");
		} else {
			document.documentElement.classList.remove("dark");
		}
		localStorage.setItem(DARK_MODE_STORAGE_KEY, String(isDark));
	}, []);

	useEffect(() => {
		const savedDarkMode = localStorage.getItem(DARK_MODE_STORAGE_KEY);
		const newIsDark = savedDarkMode === "true";
		handleThemeChange(newIsDark);
		setIsDark(newIsDark);
	}, [handleThemeChange]);

	const toggleDarkMode = () => {
		setIsDark((prev) => {
			const newIsDark = !prev;
			handleThemeChange(newIsDark);

			return newIsDark;
		});
	};

	return (
		<Button
			variant="neutral"
			size="icon"
			className={cn("size-12 text-xl")}
			aria-label="Change theme"
			onClick={toggleDarkMode}
		>
			{isDark ? <Moon className="h-5! w-5!" /> : <Sun className="h-5! w-5!" />}
		</Button>
	);
};
