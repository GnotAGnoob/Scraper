import { cn } from "@/lib/utils"

const Skeleton = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => {
  return (
    <div
      className={cn("animate-pulse rounded-[--radius-base] bg-foreground/10", className)}
      {...props}
    />
  )
}

export { Skeleton }


