import { Loader2, Sparkles } from 'lucide-react'

import { Button } from '@/components/ui/button'

type SubmitButtonProps = {
  isLoading: boolean
  disabled?: boolean
}

export const SubmitButton = ({ isLoading, disabled }: SubmitButtonProps) => {
  const isDisabled = disabled || isLoading

  return (
    <Button
      type="submit"
      size="lg"
      disabled={isDisabled}
      className="mt-3 w-full justify-center rounded-[var(--radius-base)] border-2 border-border bg-main text-main-foreground text-base font-heading uppercase tracking-wide shadow-shadow transition-transform hover:-translate-y-1 hover:translate-x-1 hover:shadow-[6px_6px_0_0_var(--border)] active:translate-y-0 active:translate-x-0 active:shadow-shadow"
    >
      {isLoading ? (
        <>
          <Loader2 className="animate-spin" />
          Crunching your recipe...
        </>
      ) : (
        <>
          <Sparkles className="text-chart-1" />
          Find products
        </>
      )}
    </Button>
  )
}



