import React from "react";
import type { ProductMatch } from "../types";

interface Props {
  match: ProductMatch;
}

export const ProductCard: React.FC<Props> = ({ match }) => {
  return (
    <div className="product-card">
      <div className="product-image-wrap">
        {match.imageUrl ? (
          <img
            src={match.imageUrl}
            alt={match.name}
            className="product-image"
          />
        ) : (
          <div className="product-image placeholder" />
        )}
      </div>
      <div className="product-body">
        <div className="product-title">{match.name}</div>
        <div className="product-ingredient">
          for <span className="badge">{match.ingredient}</span>{" "}
          {Math.round(match.quantityGrams)} g
        </div>
        <div className="product-price">
          {match.price.toFixed(2)} {match.currency}
        </div>
        <a
          className="product-link"
          href={match.url}
          target="_blank"
          rel="noreferrer"
        >
          Open on Kosik
        </a>
      </div>
    </div>
  );
};


