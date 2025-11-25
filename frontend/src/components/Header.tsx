import { Link } from '@tanstack/react-router'
import { useCallback, useEffect, useState } from 'react'
import { ChefHat, Moon, Sun } from 'lucide-react'

export default function Header() {
  const [isDark, setIsDark] = useState(false)

  useEffect(() => {
    const stored = localStorage.getItem('theme')
    const isDarkMode = stored === 'dark' || (!stored && window.matchMedia('(prefers-color-scheme: dark)').matches)
    setIsDark(isDarkMode)
    document.documentElement.classList.toggle('dark', isDarkMode)
  }, [])

  const toggleTheme = useCallback(() => {
    const newIsDark = !isDark
    setIsDark(newIsDark)
    document.documentElement.classList.toggle('dark', newIsDark)
    localStorage.setItem('theme', newIsDark ? 'dark' : 'light')
  }, [isDark])

  return (
    <header className="border-b-4 border-border bg-secondary-background shadow-[var(--shadow)]">
      <div className="container mx-auto flex items-center justify-between p-4">
        <Link
          to="/"
          className="flex items-center gap-3 transition-transform hover:scale-105"
        >
          <div className="rounded-[--radius-base] border-4 border-border bg-main p-2 shadow-[var(--shadow)]">
            <ChefHat className="h-8 w-8 text-main-foreground" />
          </div>
          <h1 className="text-2xl font-heading">
            Recipe to Kosik
          </h1>
        </Link>

        <button
          onClick={toggleTheme}
          className="rounded-[--radius-base] border-4 border-border bg-foreground p-3 text-background shadow-[var(--shadow)] transition-all hover:shadow-[6px_6px_0px_0px_var(--border)] hover:translate-x-[-2px] hover:translate-y-[-2px]"
          aria-label="Toggle theme"
        >
          {isDark ? <Sun className="h-5 w-5" /> : <Moon className="h-5 w-5" />}
        </button>
      </div>
    </header>
  )
}
