/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { api_RecipeRequest } from '../models/api_RecipeRequest';
import type { model_RecipeResponse } from '../models/model_RecipeResponse';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class RecipeService {
    /**
     * Extracts ingredients and matches Kosik products
     * Accepts raw recipe text, extracts structured ingredients via Gemini, and fetches the first Kosik product per ingredient.
     * @returns model_RecipeResponse OK
     * @throws ApiError
     */
    public static postApiV1Recipe({
        request,
    }: {
        /**
         * Recipe input
         */
        request: api_RecipeRequest,
    }): CancelablePromise<model_RecipeResponse> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/api/v1/recipe',
            body: request,
            errors: {
                400: `Bad Request`,
                500: `Internal Server Error`,
            },
        });
    }
}
