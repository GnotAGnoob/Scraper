import * as React from 'react'

import { cn } from '@/lib/utils'

export type TextareaProps = React.TextareaHTMLAttributes<HTMLTextAreaElement>

const Textarea = React.forwardRef<HTMLTextAreaElement, TextareaProps>(
  ({ className, ...props }, ref) => {
    return (
      <textarea
        ref={ref}
        className={cn(
          'flex min-h-[160px] w-full rounded-[var(--radius-base)] border-2 border-border bg-secondary-background px-4 py-3 text-base leading-relaxed shadow-shadow outline-none placeholder:text-neutral-500 focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] disabled:cursor-not-allowed disabled:opacity-60',
          className,
        )}
        {...props}
      />
    )
  },
)

Textarea.displayName = 'Textarea'

export { Textarea }



