import './SubmitButton.css'

interface SubmitButtonProps {
  onClick: () => void
  disabled?: boolean
  children: React.ReactNode
}

function SubmitButton({ onClick, disabled, children }: SubmitButtonProps) {
  return (
    <button
      className="submit-button"
      onClick={onClick}
      disabled={disabled}
    >
      {children}
    </button>
  )
}

export default SubmitButton

