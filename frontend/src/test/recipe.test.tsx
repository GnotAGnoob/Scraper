import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { cleanup, fireEvent, render, screen, waitFor } from "@testing-library/react";
import { afterEach, describe, expect, it, vi } from "vitest";
import { ProductGrid } from "@/components/ProductGrid";
import { type IngredientProducts, recipeApi } from "@/lib/api";
import { Route } from "../routes/index";

const product = {
	name: { value: "Butter" },
	price: { value: 50 },
	unit: { value: "kg" },
	imageUrl: { value: "https://example.com/butter.png" },
	link: { value: "https://www.kosik.cz/api/front/product/slug/p123-butter" },
	storefrontUrl: { value: "https://www.kosik.cz/p123-butter" },
};
const groups: IngredientProducts[] = [
	{
		ingredient: { name: "butter", quantityGrams: 200 },
		products: [null, product, product],
	},
	{ ingredient: { name: "salt" }, products: null },
];

afterEach(() => {
	cleanup();
	vi.restoreAllMocks();
	vi.unstubAllGlobals();
});

const renderHome = () => {
	const Home = Route.options.component as React.ComponentType;
	return render(
		<QueryClientProvider client={new QueryClient({ defaultOptions: { mutations: { retry: false } } })}>
			<Home />
		</QueryClientProvider>,
	);
};
const submit = () => {
	fireEvent.change(screen.getByLabelText("Your Recipe"), { target: { value: "200g butter" } });
	fireEvent.click(screen.getByRole("button", { name: "Find Products" }));
};

describe("recipe results", () => {
	it("preserves the textarea DOM node, focus and selection while typing and loading an example", () => {
		renderHome();
		const input = screen.getByLabelText("Your Recipe") as HTMLTextAreaElement;
		input.focus();
		fireEvent.change(input, { target: { value: "b" } });
		fireEvent.change(input, { target: { value: "butter" } });
		input.setSelectionRange(2, 2);
		fireEvent.change(input, { target: { value: "buXtter", selectionStart: 3, selectionEnd: 3 } });
		expect(screen.getByLabelText("Your Recipe")).toBe(input);
		expect(document.activeElement).toBe(input);
		expect(input.selectionStart).toBe(3);
		fireEvent.click(screen.getByRole("button", { name: "Try Example" }));
		expect(screen.getByLabelText("Your Recipe")).toBe(input);
		expect(input.value).toContain("Chocolate Chip Cookies");
	});
	it("renders partial results, counts cards, links the storefront and labels grams", async () => {
		vi.spyOn(recipeApi, "parseRecipe").mockResolvedValue(groups);
		renderHome();
		submit();
		await screen.findByRole("heading", { name: "Found 2 Products" });
		const links = screen.getAllByRole("link", { name: "View" });
		expect(links).toHaveLength(2);
		expect(links[0].getAttribute("href")).toBe(product.storefrontUrl.value);
		expect(screen.getAllByText("Required: 200 g")).toHaveLength(2);
		expect(screen.queryByText("200 kg")).toBeNull();
	});
	it.each([{ data: [] }, { data: [{ products: [null] }] }, { data: [{ products: null }] }])(
		"shows completed empty results for %j",
		async ({ data }) => {
			vi.spyOn(recipeApi, "parseRecipe").mockResolvedValue(data);
			renderHome();
			submit();
			await screen.findByText("No products found");
			expect(screen.getByRole("heading", { name: "Found 0 Products" })).toBeTruthy();
			expect(screen.queryByText('Click "Find Products" to start searching')).toBeNull();
		},
	);
	it("clears previous cards during a retry and after failure, then recovers", async () => {
		let rejectRequest!: (error: Error) => void;
		vi.spyOn(recipeApi, "parseRecipe")
			.mockResolvedValueOnce(groups)
			.mockImplementationOnce(
				() =>
					new Promise((_, reject) => {
						rejectRequest = reject;
					}),
			)
			.mockResolvedValueOnce([]);
		renderHome();
		submit();
		await screen.findByRole("heading", { name: "Found 2 Products" });
		submit();
		await waitFor(() => expect(screen.queryByRole("link", { name: "View" })).toBeNull());
		rejectRequest(new Error("Service unavailable"));
		await screen.findByText("Error: Service unavailable");
		expect(screen.queryByRole("link", { name: "View" })).toBeNull();
		expect(screen.queryByText("No products found")).toBeNull();
		submit();
		await screen.findByText("No products found");
		expect(screen.queryByText("Error: Service unavailable")).toBeNull();
	});
	it("does not render a View link without a storefront URL", () => {
		render(<ProductGrid products={[{ products: [{ ...product, storefrontUrl: undefined }] }]} />);
		expect(screen.queryByRole("link")).toBeNull();
	});
});

describe("recipe API", () => {
	it("posts JSON to the products endpoint and accepts partial results", async () => {
		const fetchMock = vi.fn().mockResolvedValue(new Response(JSON.stringify({ ingredientsProducts: groups })));
		vi.stubGlobal("fetch", fetchMock);
		expect(await recipeApi.parseRecipe("butter")).toEqual(groups);
		expect(fetchMock).toHaveBeenCalledWith("http://localhost:8080/api/v1/text/products", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: '{"text":"butter"}',
		});
	});
	it("rejects HTTP errors so the mutation enters the error state", async () => {
		vi.stubGlobal("fetch", vi.fn().mockResolvedValue(new Response("unavailable", { status: 503 })));
		await expect(recipeApi.parseRecipe("butter")).rejects.toThrow("HTTP 503");
	});
	it("accepts an empty or null result list", async () => {
		const fetchMock = vi
			.fn()
			.mockResolvedValueOnce(new Response("{}"))
			.mockResolvedValueOnce(new Response('{"ingredientsProducts":null}'));
		vi.stubGlobal("fetch", fetchMock);
		expect(await recipeApi.parseRecipe("butter")).toEqual([]);
		expect(await recipeApi.parseRecipe("butter")).toEqual([]);
	});
});
