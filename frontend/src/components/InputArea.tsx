import { useCallback, useId, useMemo } from 'react'
import { Wand2 } from 'lucide-react'

import { Badge } from '@/components/ui/badge'
import { Textarea } from '@/components/ui/textarea'
import SubmitButton from '@/components/SubmitButton'
import type { ExampleRecipe } from '@/lib/recipes'
import { cn } from '@/lib/utils'

type InputAreaProps = {
  value: string
  onChange: (nextValue: string) => void
  onSubmit: () => void
  isLoading?: boolean
  disabled?: boolean
  charLimit: number
  examples: ExampleRecipe[]
  activeExampleId?: string | null
  onExampleSelect: (example: ExampleRecipe) => void
}

const InputArea = ({
  value,
  onChange,
  onSubmit,
  isLoading,
  disabled,
  charLimit,
  examples,
  activeExampleId,
  onExampleSelect,
}: InputAreaProps) => {
  const textareaId = useId()
  const length = useMemo(() => value.length, [value])
  const remaining = charLimit - length
  const isOverLimit = remaining < 0

  const handleExampleClick = useCallback(
    (example: ExampleRecipe) => {
      onExampleSelect(example)
    },
    [onExampleSelect],
  )

  return (
    <div className="space-y-4 rounded-[var(--radius-base)] border-2 border-border bg-secondary-background p-6 shadow-[var(--shadow)]">
      <div className="flex items-center justify-between gap-3">
        <div>
          <p className="text-xs font-semibold uppercase tracking-[0.2em] text-foreground/70">
            recipe intelligence
          </p>
          <h2 className="text-3xl font-heading uppercase leading-tight">Paste your recipe</h2>
        </div>
        <Badge variant="outline" className="hidden md:inline-flex">
          LLM + scraper
        </Badge>
      </div>

      <label htmlFor={textareaId} className="sr-only">
        recipe text
      </label>
      <Textarea
        id={textareaId}
        placeholder="Describe the dish or paste raw recipe text..."
        value={value}
        onChange={(event) => onChange(event.target.value)}
        disabled={disabled}
        className={cn(isOverLimit && 'border-red-500 text-red-600')}
      />
      <div className="flex flex-wrap items-center gap-3 text-sm font-semibold uppercase">
        <span className={cn(isOverLimit && 'text-red-600')}>
          {length.toLocaleString()} / {charLimit.toLocaleString()} chars
        </span>
        <span className="text-foreground/60">•</span>
        <span className="text-foreground/80">No formatting required</span>
        {isOverLimit && <span className="text-red-600">Trim {Math.abs(remaining)} chars</span>}
      </div>

      <SubmitButton
        isLoading={isLoading}
        disabled={disabled || isOverLimit || length === 0}
        onClick={onSubmit}
      />

      <div className="space-y-3">
        <p className="text-xs font-semibold uppercase tracking-[0.2em] text-foreground/70">
          quick starters
        </p>
        <div className="flex flex-wrap gap-3">
          {examples.map((example) => (
            <button
              key={example.id}
              type="button"
              onClick={() => handleExampleClick(example)}
              className={cn(
                'flex flex-1 min-w-[140px] items-center justify-between rounded-[var(--radius-base)] border-2 border-border bg-background px-4 py-3 text-left text-sm font-semibold uppercase tracking-wide shadow-[var(--shadow)] transition-all hover:-translate-x-1 hover:-translate-y-1 hover:shadow-[6px_6px_0px_0px_var(--border)]',
                activeExampleId === example.id && 'bg-main text-main-foreground',
              )}
            >
              <span>
                {example.title}
                <span className="block text-[11px] font-normal normal-case text-foreground/70">
                  {example.minutes} min • {example.tone}
                </span>
              </span>
              <Wand2 className="size-4" />
            </button>
          ))}
        </div>
      </div>
    </div>
  )
}

export default InputArea


