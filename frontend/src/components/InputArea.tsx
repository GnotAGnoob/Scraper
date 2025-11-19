import './InputArea.css'

interface InputAreaProps {
  value: string
  onChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void
  placeholder?: string
}

function InputArea({ value, onChange, placeholder }: InputAreaProps) {
  return (
    <textarea
      className="input-area"
      value={value}
      onChange={onChange}
      placeholder={placeholder}
      rows={10}
    />
  )
}

export default InputArea

