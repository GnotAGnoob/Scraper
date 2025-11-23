import { useCallback, useState } from 'react'

import { Link } from '@tanstack/react-router'
import { Home, Menu, Moon, Sun } from 'lucide-react'

const Header = () => {
  const [isOpen, setIsOpen] = useState(false)
  const [isDark, setIsDark] = useState(
    typeof document !== 'undefined'
      ? document.documentElement.classList.contains('dark')
      : false,
  )

  const openMenu = useCallback(() => {
    setIsOpen(true)
  }, [])

  const closeMenu = useCallback(() => {
    setIsOpen(false)
  }, [])

  const toggleTheme = useCallback(() => {
    if (typeof document === 'undefined') {
      return
    }

    const root = document.documentElement
    root.classList.toggle('dark')
    setIsDark(root.classList.contains('dark'))
  }, [])

  return (
    <>
      <header className="flex items-center justify-between gap-4 border-b-2 border-border bg-secondary-background px-4 py-3 shadow-shadow">
        <div className="flex items-center gap-3">
          <button
            type="button"
            onClick={openMenu}
            className="inline-flex h-10 w-10 items-center justify-center rounded-[var(--radius-base)] border-2 border-border bg-background text-foreground shadow-shadow transition-transform hover:-translate-y-[2px] hover:translate-x-[2px] hover:shadow-[5px_5px_0_0_var(--border)] md:hidden"
            aria-label="Open menu"
          >
            <Menu className="h-5 w-5" />
          </button>

          <Link
            to="/"
            className="inline-flex items-baseline gap-2 rounded-[var(--radius-base)] border-2 border-border bg-main px-3 py-1 font-heading text-sm uppercase tracking-wide shadow-shadow"
          >
            <Home className="h-4 w-4" />
            <span>Kosik scraper</span>
          </Link>
        </div>

        <button
          type="button"
          onClick={toggleTheme}
          className="inline-flex items-center gap-2 rounded-[var(--radius-base)] border-2 border-border bg-background px-3 py-1 text-xs font-heading uppercase shadow-shadow transition-transform hover:-translate-y-[2px] hover:translate-x-[2px] hover:shadow-[5px_5px_0_0_var(--border)]"
          aria-label="Toggle theme"
        >
          {isDark ? (
            <>
              <Sun className="h-4 w-4 text-chart-1" />
              Light
            </>
          ) : (
            <>
              <Moon className="h-4 w-4 text-chart-5" />
              Dark
            </>
          )}
        </button>
      </header>

      {isOpen ? (
        <aside className="fixed inset-y-0 left-0 z-50 flex w-72 flex-col border-r-2 border-border bg-secondary-background p-4 shadow-shadow">
          <div className="mb-4 flex items-center justify-between gap-2">
            <span className="font-heading text-lg">Navigation</span>
            <button
              type="button"
              onClick={closeMenu}
              className="inline-flex h-8 w-8 items-center justify-center rounded-[var(--radius-base)] border-2 border-border bg-background text-foreground shadow-shadow"
            >
              ✕
            </button>
          </div>

          <nav className="flex flex-col gap-2 text-sm">
            <Link
              to="/"
              onClick={closeMenu}
              className="inline-flex items-center gap-2 rounded-[var(--radius-base)] border-2 border-border bg-background px-3 py-2 font-base shadow-shadow"
              activeProps={{
                className:
                  'inline-flex items-center gap-2 rounded-[var(--radius-base)] border-2 border-border bg-main px-3 py-2 font-heading shadow-shadow',
              }}
            >
              <Home className="h-4 w-4" />
              Home
            </Link>
          </nav>
        </aside>
      ) : null}
    </>
  )
}

export default Header
