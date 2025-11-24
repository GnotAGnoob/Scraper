import { createFileRoute } from '@tanstack/react-router'
import { RecipeScraper } from '@/components/RecipeScraper'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  return (
    <div className="min-h-screen bg-background font-sans">
      <RecipeScraper />
    </div>
  )
}
