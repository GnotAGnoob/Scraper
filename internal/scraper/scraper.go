package scraper

import (
	"fmt"
	"log"

	errorsUtil "github.com/GnotAGnoob/kosik-scraper/pkg/utils/errors"
	"github.com/GnotAGnoob/kosik-scraper/pkg/utils/urlParams"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

type nutrition struct {
	Calories     string
	Protein      string
	Fat          string
	SaturatedFat string
	Carbs        string
	Sugar        string
	Fiber        string
}

type ParsedProduct struct {
	Name        string
	Price       string
	PricePerKg  string
	Unit        string
	Link        string
	Image       *[]byte
	Description string
	Nutrition   *nutrition
	AddButton   *rod.Element
}

type returnProduct struct {
	Product *ParsedProduct
	Errors  []error
}

type Scraper struct {
	launcher *launcher.Launcher
}

func NewScraper() *Scraper {
	return &Scraper{
		// leakless is a binary that prevents zombie processes
		// but the problem is that windows defender detects it as a virus
		// because according to internet, it is used in many viruses
		launcher: launcher.New().Leakless(false),
	}
}

// todo separate the code into smaller reusable functions
// todo handle timeout => send what was found and errors for the rest
// todo goroutines for each product and for nutrition page
func (s *Scraper) GetKosikProducts(search string) (*[]*returnProduct, error) {
	searchUrl, err := urlParams.CreateSearchUrl(search)
	if err != nil {
		return nil, err
	}

	controlUrl, err := s.launcher.Launch()
	if err != nil {
		return nil, err
	}

	browser := rod.New().NoDefaultDevice().ControlURL(controlUrl)

	if error := browser.Connect(); error != nil {
		return nil, error
	}
	defer func() {
		err := browser.Close()
		if err != nil {
			log.Fatalf("Error: %v", err)
		}

	}()

	page, err := browser.Page(proto.TargetCreateTarget{
		URL: searchUrl.String(),
	})
	if err != nil {
		return nil, err
	}

	productSelector := "[data-tid='product-box']:not(:has(.product-amount--vendor-pharmacy))"
	err = page.WaitElementsMoreThan(productSelector, 1)
	if err != nil {
		return nil, errorsUtil.ElementNotFoundError(err, productSelector)
	}

	products, err := page.Elements(productSelector)
	if err != nil {
		return nil, errorsUtil.ElementNotFoundError(err, productSelector)
	}

	parsedProducts := make([]*returnProduct, 0, len(products))

	for _, product := range products[:2] {
		errors := []error{}
		parsedProduct := &ParsedProduct{}

		nameSelector := "a[data-tid='product-box__name']"
		nameElement, err := product.Element(nameSelector)
		if err != nil {
			errors = append(errors, errorsUtil.ElementNotFoundError(err, nameSelector))
		} else {
			name, err := nameElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Name = name

			hrefAttribute := "href"
			href, err := nameElement.Attribute(hrefAttribute)
			if err != nil {
				errors = append(errors, errorsUtil.ElementNotFoundError(err, hrefAttribute))
			}
			parsedProduct.Link = *href
		}

		unitSelector := ".attributes"
		unitElement, err := product.Element(unitSelector)
		if err != nil {
			errors = append(errors, errorsUtil.ElementNotFoundError(err, unitSelector))
		} else {
			unit, err := unitElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Unit = unit
		}

		imageSelector := "img"
		imageElement, err := product.Element(imageSelector)
		if err != nil {
			errors = append(errors, errorsUtil.ElementNotFoundError(err, imageSelector))
		} else {
			image, err := imageElement.Resource()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Image = &image
		}

		pricePrefix := ""
		pricePrefixElement, err := product.Element(".price__prefix")
		if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
			errors = append(errors, err)
		} else if err == nil {
			pricePrefix, err = pricePrefixElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			pricePrefix += " "
		}

		priceSelector := "[data-tid='product-price']"
		priceElement, err := product.Element(priceSelector)
		if err != nil {
			errors = append(errors, errorsUtil.ElementNotFoundError(err, priceSelector))
		} else {
			price, err := priceElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.Price = pricePrefix + price
		}

		pricePerKgSelector := "[aria-label='Cena'] > *:last-child"
		pricePerKgElement, err := product.Element(pricePerKgSelector)
		if err != nil {
			errors = append(errors, errorsUtil.ElementNotFoundError(err, pricePerKgSelector))
		} else {
			pricePerKg, err := pricePerKgElement.Text()
			if err != nil {
				errors = append(errors, err)
			}
			parsedProduct.PricePerKg = pricePerKg
		}

		buttonSelector := "[data-tid='product-to-cart__to-cart']"
		buttonElement, err := product.Element(buttonSelector)
		if err != nil {
			errors = append(errors, errorsUtil.ElementNotFoundError(err, buttonSelector))
		}
		parsedProduct.AddButton = buttonElement

		ingredientsUrl, err := urlParams.CreateUrlFromPath(parsedProduct.Link)
		if err != nil {
			errors = append(errors, err)
		}
		ingredientsUrl.Fragment = "ingredients"

		indgredientsPage, err := browser.Page(proto.TargetCreateTarget{
			URL: ingredientsUrl.String(),
		})
		if err != nil {
			errors = append(errors, err)
		} else {
			nutrition := &nutrition{}
			nutritionElement, err := indgredientsPage.Element("[data-tid='product-detail__nutrition_table']")
			if _, ok := err.(*rod.ElementNotFoundError); err != nil && !ok {
				fmt.Printf("xxxxx: %T", err)
				errors = append(errors, err)
			} else if err == nil {
				caloriesRegex := "\\d* kcal"
				caloriesElement, err := nutritionElement.ElementR("td", caloriesRegex)
				if err != nil {
					errors = append(errors, errorsUtil.ElementNotFoundError(err, caloriesRegex))
				}

				calories, err := caloriesElement.Text()
				if err != nil {
					errors = append(errors, err)
				}
				nutrition.Calories = calories

				caloriesParentElement, err := caloriesElement.Parent()
				if err != nil {
					errors = append(errors, err)
				} else {
					// todo better selectors
					// fatElement, err := caloriesParentElement.Next()
					// if err != nil {
					// 	errors = append(errors, err)
					// } else {
					// 	fat, err := fatElement.Text()
					// 	if err != nil {
					// 		errors = append(errors, err)
					// 	}
					// 	nutrition.Fat = fat
					// }

					// saturatedFatElement, err := fatElement.Next()
					// if err != nil {
					// 	errors = append(errors, err)
					// } else {
					// 	saturatedFat, err := saturatedFatElement.Text()
					// 	if err != nil {
					// 		errors = append(errors, err)
					// 	}
					// 	nutrition.SaturatedFat = saturatedFat
					// }

					// carbsElement, err := saturatedFatElement.Next()
					// if err != nil {
					// 	errors = append(errors, err)
					// } else {
					// 	carbs, err := carbsElement.Text()
					// 	if err != nil {
					// 		errors = append(errors, err)
					// 	}
					// 	nutrition.Carbs = carbs
					// }

					// sugarElement, err := carbsElement.Next()
					// if err != nil {
					// 	errors = append(errors, err)
					// } else {
					// 	sugar, err := sugarElement.Text()
					// 	if err != nil {
					// 		errors = append(errors, err)
					// 	}
					// 	nutrition.Sugar = sugar
					// }

					// fiberElement, err := sugarElement.Next()
					// if err != nil {
					// 	errors = append(errors, err)
					// } else {
					// 	fiber, err := fiberElement.Text()
					// 	if err != nil {
					// 		errors = append(errors, err)
					// 	}
					// 	nutrition.Fiber = fiber
					// }

					// proteinElement, err := fiberElement.Next()
					// if err != nil {
					// 	errors = append(errors, err)
					// } else {
					// 	protein, err := proteinElement.Text()
					// 	if err != nil {
					// 		errors = append(errors, err)
					// 	}
					// 	nutrition.Protein = protein
					// }
				}
			}
			parsedProduct.Nutrition = nutrition
		}

		fmt.Println(parsedProduct, errors)

		parsedProducts = append(parsedProducts, &returnProduct{
			Product: parsedProduct,
			Errors:  errors,
		})
	}

	return &parsedProducts, nil
}
