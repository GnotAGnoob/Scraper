import React, { useState } from "react";
import { useMutation } from "@tanstack/react-query";
import { analyzeRecipe } from "./api";
import { InputArea } from "./components/InputArea";
import { SubmitButton } from "./components/SubmitButton";
import { ProductGrid } from "./components/ProductGrid";

function App() {
  const [text, setText] = useState("");

  const mutation = useMutation({
    mutationFn: analyzeRecipe
  });

  const handleSubmit = () => {
    if (!text.trim()) return;
    mutation.mutate(text);
  };

  return (
    <div className="app-root">
      <header className="app-header">
        <h1>Recipe → Kosik Cart</h1>
        <p>Paste a recipe, get matching Kosik.cz products.</p>
      </header>

      <main className="app-main">
        <section className="left-column">
          <InputArea value={text} onChange={setText} />
          <SubmitButton
            onClick={handleSubmit}
            disabled={!text.trim()}
            loading={mutation.isPending}
          />
          {mutation.isError && (
            <div className="error-banner">
              Something went wrong. Check backend logs.
            </div>
          )}
        </section>

        <section className="right-column">
          <ProductGrid matches={mutation.data?.matches ?? []} />
        </section>
      </main>

      <footer className="app-footer">
        <span>Backend: http://localhost:8080</span>
        <span>Swagger: /swagger/index.html</span>
      </footer>
    </div>
  );
}

export default App;


