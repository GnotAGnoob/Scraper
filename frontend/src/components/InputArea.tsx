import type { ChangeEvent } from 'react'
import { useCallback } from 'react'

import { Textarea } from '@/components/ui/textarea'

type InputAreaProps = {
  value: string
  onChange: (value: string) => void
  disabled?: boolean
}

export const InputArea = ({ value, onChange, disabled }: InputAreaProps) => {
  const handleChange = useCallback(
    (event: ChangeEvent<HTMLTextAreaElement>) => {
      onChange(event.target.value)
    },
    [onChange],
  )

  const remainingHint =
    value.trim().length === 0
      ? 'Paste or type your recipe here – ingredients, steps, messy text all welcome.'
      : undefined

  return (
    <div className="space-y-3">
      <div className="flex items-baseline justify-between gap-2">
        <h1 className="text-3xl md:text-4xl font-heading">
          Kosik recipe to cart
        </h1>
        <span className="hidden text-xs md:inline-flex rounded-full border-2 border-border bg-secondary-background px-3 py-1 font-base shadow-shadow">
          LLM → ingredients → Kosik search → cart candidates
        </span>
      </div>

      <Textarea
        value={value}
        onChange={handleChange}
        disabled={disabled}
        placeholder="e.g. 500 g chicken breast, 2 onions, 200 g rice, 1 can of tomatoes..."
      />

      <div className="flex items-center justify-between text-xs text-neutral-700 dark:text-neutral-200">
        <span>{remainingHint}</span>
        <span className="rounded-full border-2 border-border bg-secondary-background px-2 py-0.5 font-base shadow-shadow">
          {value.length} chars
        </span>
      </div>
    </div>
  )
}



