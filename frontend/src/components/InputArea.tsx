import { useMemo } from 'react'
import { Textarea } from '@/components/ui/textarea'
import { Label } from '@/components/ui/label'
import { cn } from '@/lib/utils'

interface InputAreaProps {
  value: string
  onChange: (value: string) => void
  placeholder?: string
  disabled?: boolean
  onKeyDown?: (e: React.KeyboardEvent<HTMLTextAreaElement>) => void
}

export const InputArea = ({ value, onChange, placeholder, disabled, onKeyDown }: InputAreaProps) => {
  const characterCount = useMemo(() => value.length, [value])

  return (
    <div className="space-y-2">
      <Label htmlFor="recipe-input" className="text-lg font-heading">
        Recipe or Ingredient List
      </Label>
      <Textarea
        id="recipe-input"
        value={value}
        onChange={(e) => onChange(e.target.value)}
        onKeyDown={onKeyDown}
        placeholder={placeholder || "Enter your recipe or list of ingredients here..."}
        disabled={disabled}
        className="min-h-[200px] text-base resize-y"
      />
      <div className="flex justify-between items-center text-sm">
        <p className="text-muted-foreground">
          Paste your recipe text or ingredient list
        </p>
        <span className={cn(
          "font-base",
          characterCount > 10000 && "text-destructive"
        )}>
          {characterCount.toLocaleString()} characters
        </span>
      </div>
    </div>
  )
}

