package scraper

import "time"

// https://www.useragentlist.net/
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35"

// selectors
const productPageWaitSelector = ".product-box-listing .product"
const productPageNotFoundWaitSelector = "[data-tid='product-no-results-title']"
const nutritionPageWaitSelector = "[data-tid='product-detail__tab-content']"
const nutritionPageNotFoundWaitSelector = ".no-results"
const ingredientsSelector = "[data-tid='product-detail__ingredients'] dd"
const nutritionSelector = "[data-tid='product-detail__nutrition_table'][aria-describedby='Výživové hodnoty (na 100 g)']"
const linkSelector = "a[data-tid='product-box__name']"
const nameSelector = ".name"
const unitSelector = ".attributes"
const pricePrefixSelector = ".price__prefix"
const priceSelector = "[data-tid='product-price']"
const pricePerKgSelector = "[aria-label='Cena'] > *:last-child"
const buttonSelector = "[data-tid='product-to-cart__to-cart']"

// timeouts
const productsPageTimeout = 5 * time.Second
const perProductTimeout = 1 * time.Second
const nutritionPageTimeout = 15 * time.Second
