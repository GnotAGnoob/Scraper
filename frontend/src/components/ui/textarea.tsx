import * as React from 'react'

import { cn } from '@/lib/utils'

const Textarea = React.forwardRef<
  HTMLTextAreaElement,
  React.TextareaHTMLAttributes<HTMLTextAreaElement>
>(({ className, ...props }, ref) => (
  <textarea
    ref={ref}
    className={cn(
      'min-h-[200px] w-full rounded-[var(--radius-base)] border-2 border-border bg-secondary-background px-4 py-3 text-base font-base text-foreground shadow-[var(--shadow)] outline-none transition-all placeholder:text-foreground/40 focus:-translate-x-1 focus:-translate-y-1 focus:shadow-[6px_6px_0px_0px_var(--border)] disabled:cursor-not-allowed disabled:opacity-60',
      className,
    )}
    {...props}
  />
))
Textarea.displayName = 'Textarea'

export { Textarea }


