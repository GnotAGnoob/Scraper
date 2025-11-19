import React from "react";

interface Props {
  onClick: () => void;
  disabled?: boolean;
  loading?: boolean;
}

export const SubmitButton: React.FC<Props> = ({
  onClick,
  disabled,
  loading
}) => {
  return (
    <button
      className="btn-primary"
      onClick={onClick}
      disabled={disabled || loading}
    >
      {loading ? "Thinking..." : "Convert to Kosik cart"}
    </button>
  );
};


