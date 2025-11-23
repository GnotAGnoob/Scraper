import { createFileRoute } from '@tanstack/react-router'
import { useState, useCallback } from 'react'
import { InputArea } from '@/components/InputArea'
import { SubmitButton } from '@/components/SubmitButton'
import { ProductGrid } from '@/components/ProductGrid'
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useTextProducts } from '@/hooks/useTextProducts'
import { AlertCircle } from 'lucide-react'

export const Route = createFileRoute('/')({
  component: App,
})

function App() {
  const [text, setText] = useState('')
  const { mutate, data, isPending, error } = useTextProducts()

  const handleSubmit = useCallback(() => {
    if (!text.trim()) return
    mutate(text)
  }, [text, mutate])

  const handleKeyDown = useCallback((e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) {
      e.preventDefault()
      handleSubmit()
    }
  }, [handleSubmit])

  return (
    <div className="min-h-screen bg-background">
      <main className="container mx-auto px-4 py-8 max-w-7xl">
        <div className="space-y-8">
          <div className="text-center space-y-4">
            <h1 className="text-5xl md:text-6xl font-heading">
              Recipe to Kosik Cart
            </h1>
            <p className="text-lg text-muted-foreground max-w-2xl mx-auto">
              Paste your recipe or ingredient list and we'll find matching products from Kosik
            </p>
          </div>

          <Card className="max-w-4xl mx-auto">
            <CardHeader>
              <CardTitle>Enter Your Recipe</CardTitle>
            </CardHeader>
            <CardContent className="space-y-6">
              <InputArea
                value={text}
                onChange={setText}
                disabled={isPending}
                onKeyDown={handleKeyDown}
              />
              <div className="flex flex-col sm:flex-row gap-4 items-start sm:items-center justify-between">
                <p className="text-sm text-muted-foreground">
                  Press <kbd className="px-2 py-1 bg-secondary-background border-2 border-border rounded-[var(--radius-base)] font-mono text-xs">Ctrl+Enter</kbd> or <kbd className="px-2 py-1 bg-secondary-background border-2 border-border rounded-[var(--radius-base)] font-mono text-xs">Cmd+Enter</kbd> to submit
                </p>
                <SubmitButton
                  onClick={handleSubmit}
                  disabled={!text.trim() || isPending}
                  isLoading={isPending}
                >
                  Find Products
                </SubmitButton>
              </div>
            </CardContent>
          </Card>

          {error && (
            <Alert variant="destructive" className="max-w-4xl mx-auto">
              <AlertCircle className="size-4" />
              <AlertTitle>Error</AlertTitle>
              <AlertDescription>
                {error.message || 'Failed to process your recipe. Please try again.'}
              </AlertDescription>
            </Alert>
          )}

          {data && (
            <div className="space-y-6">
              <div className="text-center">
                <h2 className="text-3xl font-heading">
                  Found Products
                </h2>
                <p className="text-muted-foreground mt-2">
                  {data.ingredientsProducts.length} ingredient{data.ingredientsProducts.length !== 1 ? 's' : ''} matched
                </p>
              </div>
              <ProductGrid data={data.ingredientsProducts} />
            </div>
          )}
        </div>
      </main>
    </div>
  )
}
