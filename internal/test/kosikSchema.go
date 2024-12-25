package main

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Image string  `json:"image"`
	URL   string  `json:"url"`
	Price float64 `json:"price"`
	// ReturnablePackagePrice      float64           `json:"returnablePackagePrice"`
	Unit string `json:"unit"`
	// RecommendedPrice            float64           `json:"recommendedPrice"`
	// PercentageDiscount          int               `json:"percentageDiscount"`
	ProductQuantity productQuantity `json:"productQuantity"`
	// Labels                      []label           `json:"labels"`
	// ActionLabel                 *string           `json:"actionLabel"`
	// CountryCode                 string            `json:"countryCode"`
	// Pictographs                 []string          `json:"pictographs"`
	// MaxInCart                   int               `json:"maxInCart"`
	// LimitInCart                 *int              `json:"limitInCart"`
	// FirstOrderDay               *string           `json:"firstOrderDay"`
	// LastOrderDay                *string           `json:"lastOrderDay"`
	// PlannedStock                *string           `json:"plannedStock"`
	// RelatedProduct              *string           `json:"relatedProduct"`
	// MainCategory                mainCategory      `json:"mainCategory"`
	PricePerUnit     pricePerUnit      `json:"pricePerUnit"`
	CumulativePrices []cumulativePrice `json:"cumulativePrices"`
	// GiftIDs                     []string          `json:"giftIds"`
	// Favorite                    bool              `json:"favorite"`
	// Purchased                   bool              `json:"purchased"`
	// UnitStep                    float64           `json:"unitStep"`
	// VendorID                    int               `json:"vendorId"`
	// PharmacyCertificate         *string           `json:"pharmacyCertificate"`
	// ProductGroups               []interface{}     `json:"productGroups"`
	// RecommendedSellPrice        float64           `json:"recommendedSellPrice"`
	// Detail                      *string           `json:"detail"`
	// HasAssociatedProducts       bool              `json:"hasAssociatedProducts"`
	// ELicence                    bool              `json:"eLicence"`
	// MarketplaceVendor           *string           `json:"marketplaceVendor"`
	// AssociationCode             *string           `json:"associationCode"`
	// LoyaltyClubBenefitsPrice    *float64          `json:"loyaltyClubBenefitsPrice"`
	// HasLoyaltyClubBenefitsPrice bool              `json:"hasLoyaltyClubBenefitsPrice"`
}

type productQuantity struct {
	Prefix string  `json:"prefix"`
	Value  float64 `json:"value"`
	Unit   string  `json:"unit"`
}

// type label struct {
// 	ID             int    `json:"id"`
// 	Name           string `json:"name"`
// 	URL            string `json:"url"`
// 	Background     string `json:"background"`
// 	Priority       int    `json:"priority"`
// 	StyleKey       string `json:"styleKey"`
// 	ExcludeFromBox bool   `json:"excludeFromBox"`
// }

// type mainCategory struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	URL         string  `json:"url"`
// 	Image       *string `json:"image"`
// 	Highlighted bool    `json:"highlighted"`
// 	VendorID    int     `json:"vendorId"`
// }

type pricePerUnit struct {
	Price float64 `json:"price"`
	Unit  string  `json:"unit"`
}

type cumulativePrice struct {
	Quantity        int          `json:"quantity"`
	Price           float64      `json:"price"`
	PricePerUnit    pricePerUnit `json:"pricePerUnit"`
	AssociationCode *string      `json:"associationCode"`
}

// type brand struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

type titleType struct {
	Title string `json:"title"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type nutritionalValue struct {
	Title string `json:"title"`
	// Prefix *string `json:"prefix"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type nutritionalValues struct {
	// ValuesPerGrams int              `json:"valuesPerGrams"`
	// Title          string           `json:"title"`
	Values []nutritionalValue `json:"values"`
}

// type parameterItem struct {
// 	Title string `json:"title"`
// 	Value string `json:"value"`
// }

// type parameterGroup struct {
// 	Title string         `json:"title"`
// 	Items []parameterItem `json:"items"`
// }

// type bestBefore struct {
// 	Usual     int `json:"usual"`
// 	Guaranteed int `json:"guaranteed"`
// }

type productDetail struct {
	// AdultOnly        bool               `json:"adultOnly"`
	// Brand            brand              `json:"brand"`
	// SapID            string             `json:"sapId"`
	// ShoppingListIds  []int              `json:"shoppingListIds"`
	// Photos           []string           `json:"photos"`
	// SupplierInfo     []titleType     `json:"supplierInfo"`
	// Origin           []titleType           `json:"origin"`
	Description       []titleType       `json:"description"`
	Ingredients       []titleType       `json:"ingredients"`
	NutritionalValues nutritionalValues `json:"nutritionalValues"`
	// ParameterGroups  []parameterGroup   `json:"parameterGroups"`
	// BestBefore       bestBefore         `json:"bestBefore"`
	// AssociationCode string  `json:"associationCode"`
	// Unlisted        bool    `json:"unlisted"`
	// MetaDescription *string `json:"metaDescription"` // Nullable in JSON
}

type productWithDetail struct {
	product
	Detail productDetail `json:"detail"`
}

type products struct {
	TotalCount int       `json:"totalCount"`
	Items      []product `json:"items"`
	Cursor     string    `json:"cursor"`
}

// type breadcrumb struct {
// 	ID   *int    `json:"id"`
// 	Name string  `json:"name"`
// 	URL  string  `json:"url"`
// }

// type subCategory struct {
// 	ID          int     `json:"id"`
// 	Name        string  `json:"name"`
// 	URL         string  `json:"url"`
// 	Image       *string `json:"image"`
// 	Highlighted bool    `json:"highlighted"`
// 	VendorID    int     `json:"vendorId"`
// }

// type filter struct {
// 	ID    string        `json:"id"`
// 	Name  string        `json:"name"`
// 	Items []FilterItem  `json:"items"`
// 	Type  string        `json:"type"`
// }

// type filterItem struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Value bool   `json:"value"`
// 	Count int    `json:"count"`
// }

// type orderBy struct {
// 	ID    string `json:"id"`
// 	Name  string `json:"name"`
// 	Value bool   `json:"value"`
// }

// type otherProducts struct {
// 	ProductCount              int           `json:"productCount"`
// 	VendorID                  int           `json:"vendorId"`
// 	NonRegulatedPharmacyProducts []interface{} `json:"nonRegulatedPharmacyProducts"`
// 	MaximumReached            bool          `json:"maximumReached"`
// }

// type Category struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	URL   string `json:"url"`
// 	Level int    `json:"level"`
// }

type SearchResponse struct {
	// Display       string      `json:"display"`
	// ProductGroups []interface{} `json:"productGroups"`
	Products products `json:"products"`
	// Title         string      `json:"title"`
	// DescriptionBefore *string `json:"descriptionBefore"`
	// DescriptionAfter *string  `json:"descriptionAfter"`
	// URL           string      `json:"url"`
	// Breadcrumbs   []Breadcrumb `json:"breadcrumbs"`
	// SubCategories []SubCategory `json:"subCategories"`
	// Widgets       []interface{} `json:"widgets"`
	// Banners       []interface{} `json:"banners"`
	// Filters       []Filter    `json:"filters"`
	// OrderBy       []OrderBy   `json:"orderBy"`
	// ShowProductsCount bool    `json:"showProductsCount"`
	// TotalCount    int         `json:"totalCount"`
	// VendorID      int         `json:"vendorId"`
	// PharmacyCertificates []interface{} `json:"pharmacyCertificates"`
	// OtherProducts OtherProducts `json:"otherProducts"`
	// LongName      *string     `json:"longName"`
	// ShortDescription *string  `json:"shortDescription"`
	// MetaDescription *string   `json:"metaDescription"`
}

type SearchMoreResponse struct {
	// Cursor     string    `json:"cursor"`
	// PharmacyCertificates []interface{} `json:"pharmacyCertificates"`
	Products []product `json:"products"`
}

type ProductDetailResponse struct {
	// Breadcrumbs      []Breadcrumb     `json:"breadcrumbs"`
	// Gifts            []Gift           `json:"gifts"`
	// ShoppingLists    []ShoppingList   `json:"shoppingLists"`
	Product productWithDetail `json:"product"`
	// CategoryTree     []Category       `json:"categoryTree"`
	// ReturnableCarrier *interface{}    `json:"returnableCarrier"`
}
