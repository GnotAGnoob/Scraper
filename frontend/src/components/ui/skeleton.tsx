import { cn } from '@/lib/utils'

type SkeletonProps = React.HTMLAttributes<HTMLDivElement>

const Skeleton = ({ className, ...props }: SkeletonProps) => {
  return (
    <div
      className={cn(
        'animate-pulse rounded-[var(--radius-base)] border-2 border-border bg-secondary-background shadow-[var(--shadow)]',
        className,
      )}
      {...props}
    />
  )
}

export { Skeleton }


