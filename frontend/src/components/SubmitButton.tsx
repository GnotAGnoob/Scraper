import { Loader2 } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { cn } from '@/lib/utils'

interface SubmitButtonProps {
  onClick: () => void
  disabled?: boolean
  isLoading?: boolean
  children?: React.ReactNode
}

export const SubmitButton = ({ onClick, disabled, isLoading, children }: SubmitButtonProps) => {
  return (
    <Button
      onClick={onClick}
      disabled={disabled || isLoading}
      size="lg"
      className={cn(
        "w-full md:w-auto min-w-[200px] text-lg font-heading py-6 px-8",
        "shadow-shadow hover:shadow-none hover:translate-x-[var(--spacing-boxShadowX)] hover:translate-y-[var(--spacing-boxShadowY)]",
        "transition-all duration-150"
      )}
    >
      {isLoading ? (
        <>
          <Loader2 className="size-5 animate-spin mr-2" />
          Processing...
        </>
      ) : (
        children || 'Find Products'
      )}
    </Button>
  )
}

