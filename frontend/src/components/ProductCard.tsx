import type { OrchestratorProductFields } from "../generated/kosik";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "./ui/card";
import { Badge } from "./ui/badge";
import { Button } from "./ui/button";
import { ExternalLink, AlertCircle } from "lucide-react";

interface ProductCardProps {
  product: OrchestratorProductFields;
  ingredientName?: string;
}

export function ProductCard({ product, ingredientName }: ProductCardProps) {
  const name = product.name?.value || "Unknown Product";
  const price = product.price?.value;
  const currency = product.currency?.value || "CZK"; // Defaulting to CZK if missing
  const imageUrl = product.imageUrl?.value;
  const link = product.link?.value;
  const isSoldOut = product.isSoldOut?.value;

  return (
    <Card className="flex flex-col h-full overflow-hidden hover:translate-x-[-2px] hover:translate-y-[-2px] hover:shadow-[6px_6px_0px_0px_var(--border)] transition-all bg-card text-card-foreground border-2 border-border shadow-shadow">
      <div className="relative aspect-square w-full overflow-hidden bg-secondary border-b-2 border-border">
        {imageUrl ? (
          <img
            src={imageUrl}
            alt={name}
            className="h-full w-full object-cover transition-transform hover:scale-105"
            loading="lazy"
          />
        ) : (
          <div className="flex h-full items-center justify-center text-muted-foreground">
            No Image
          </div>
        )}
        {isSoldOut && (
          <div className="absolute top-2 right-2">
            <Badge variant="destructive" className="border-2 border-black text-white font-bold">Sold Out</Badge>
          </div>
        )}
      </div>
      
      <CardHeader className="p-4 pb-2 space-y-2">
        {ingredientName && (
          <Badge variant="outline" className="w-fit bg-chart-1 text-black border-2 border-black font-bold">
            Matches: {ingredientName}
          </Badge>
        )}
        <CardTitle className="line-clamp-2 text-lg leading-tight min-h-[3rem]">
          {name}
        </CardTitle>
      </CardHeader>

      <CardContent className="p-4 pt-0 flex-grow">
        <div className="flex items-baseline gap-1 mt-2">
          {price !== undefined ? (
            <>
              <span className="text-2xl font-black">{price}</span>
              <span className="text-sm font-bold text-muted-foreground">{currency}</span>
            </>
          ) : (
             <span className="text-muted-foreground flex items-center gap-1">
               <AlertCircle size={16} /> Price unknown
             </span>
          )}
        </div>
      </CardContent>

      <CardFooter className="p-4 pt-0">
        {link && (
          <Button 
            asChild 
            className="w-full font-bold border-2 border-black shadow-[2px_2px_0px_0px_rgba(0,0,0,1)] hover:shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] hover:translate-x-[-1px] hover:translate-y-[-1px] active:shadow-none active:translate-x-[2px] active:translate-y-[2px] transition-all bg-main hover:bg-main/90 text-main-foreground"
          >
            <a href={link} target="_blank" rel="noopener noreferrer" className="flex items-center gap-2 justify-center">
              View on Kosik <ExternalLink size={16} />
            </a>
          </Button>
        )}
      </CardFooter>
    </Card>
  );
}

