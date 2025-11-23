import * as React from 'react'

import { cn } from '@/lib/utils'

const badgeVariants = {
  default:
    'bg-main text-main-foreground border-2 border-border shadow-[var(--shadow)]',
  outline: 'bg-secondary-background text-foreground border-2 border-border',
  subtle: 'bg-foreground text-background border-2 border-border',
}

type Variant = keyof typeof badgeVariants

type BadgeProps = React.HTMLAttributes<HTMLDivElement> & {
  variant?: Variant
}

const Badge = React.forwardRef<HTMLDivElement, BadgeProps>(
  ({ className, variant = 'default', ...props }, ref) => (
    <div
      ref={ref}
      className={cn(
        'inline-flex items-center rounded-[var(--radius-base)] px-3 py-1 text-xs font-semibold uppercase tracking-wide',
        badgeVariants[variant],
        className,
      )}
      {...props}
    />
  ),
)
Badge.displayName = 'Badge'

export { Badge }


