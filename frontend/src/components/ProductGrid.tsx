import React from "react";
import type { ProductMatch } from "../types";
import { ProductCard } from "./ProductCard";

interface Props {
  matches: ProductMatch[];
}

export const ProductGrid: React.FC<Props> = ({ matches }) => {
  if (matches.length === 0) {
    return (
      <div className="panel">
        <div className="panel-title">Results</div>
        <div className="panel-body muted">
          No products yet. Paste a recipe and hit the button.
        </div>
      </div>
    );
  }

  return (
    <div className="panel">
      <div className="panel-title">
        Results <span className="badge">{matches.length}</span>
      </div>
      <div className="product-grid">
        {matches.map((m) => (
          <ProductCard key={`${m.ingredient}-${m.url}`} match={m} />
        ))}
      </div>
    </div>
  );
};


