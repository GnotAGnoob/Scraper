import { Loader2, Sparkles } from 'lucide-react'

import { Button } from '@/components/ui/button'

type SubmitButtonProps = {
  isLoading?: boolean
  disabled?: boolean
  label?: string
  onClick?: () => void
}

const SubmitButton = ({
  isLoading,
  disabled,
  label = 'Build my Kosik cart',
  onClick,
}: SubmitButtonProps) => {
  return (
    <Button
      size="lg"
      className="w-full md:w-auto"
      disabled={disabled || isLoading}
      onClick={onClick}
      data-loading={isLoading}
    >
      {isLoading ? (
        <>
          <Loader2 className="animate-spin" />
          crunching
        </>
      ) : (
        <>
          <Sparkles className="text-foreground" />
          {label}
        </>
      )}
    </Button>
  )
}

export default SubmitButton


