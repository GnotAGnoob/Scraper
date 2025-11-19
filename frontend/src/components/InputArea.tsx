import React from "react";

interface Props {
  value: string;
  onChange: (value: string) => void;
}

export const InputArea: React.FC<Props> = ({ value, onChange }) => {
  return (
    <div className="panel">
      <label className="panel-title" htmlFor="recipe-input">
        Recipe text
      </label>
      <textarea
        id="recipe-input"
        className="textarea"
        placeholder="Paste your recipe here (CZ or EN)..."
        value={value}
        onChange={(e) => onChange(e.target.value)}
        rows={12}
      />
    </div>
  );
};


