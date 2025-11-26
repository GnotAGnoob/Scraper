import { Api } from "../generated/kosik";

export const api = new Api({
	baseUrl: import.meta.env.VITE_API_URL || "http://localhost:8080",
});
