package scraper

// https://www.useragentlist.net/
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35"

// selectors
const ingredientsSelector = "[data-tid='product-detail__ingredients'] dd"
const nutritionSelector = "[data-tid='product-detail__nutrition_table'][aria-describedby='Výživové hodnoty (na 100 g)']"
const nameSelector = "a[data-tid='product-box__name']"
const unitSelector = ".attributes"
const pricePrefixSelector = ".price__prefix"
const priceSelector = "[data-tid='product-price']"
const pricePerKgSelector = "[aria-label='Cena'] > *:last-child"
const buttonSelector = "[data-tid='product-to-cart__to-cart']"
