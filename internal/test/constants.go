package main

import "time"

// https://www.useragentlist.net/
const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36 Edg/113.0.1774.35"

// timeouts
const productsPageTimeout = 10 * time.Second
const nutritionPageTimeout = 2 * time.Second
