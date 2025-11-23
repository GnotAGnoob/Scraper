export type ExampleRecipe = {
  id: string
  title: string
  tone: 'weekday' | 'mealprep' | 'weekend'
  text: string
  minutes: number
}

export const exampleRecipes: ExampleRecipe[] = [
  {
    id: 'weekday-bowl',
    title: 'Crispy tofu power bowl',
    tone: 'weekday',
    minutes: 25,
    text: `Quick tofu bowl for two:
- 200 g firm tofu, cubed
- 1 head broccoli, florets only
- 1 red pepper
- 100 g jasmine rice
- 1 avocado
- Dressing = 2 tbsp soy, 1 tbsp sesame oil, 1 tsp honey, lime juice`,
  },
  {
    id: 'mealprep-soup',
    title: 'Sunday soup prep',
    tone: 'mealprep',
    minutes: 15,
    text: `Make a big pot of veggie soup:
- 3 carrots, 3 celery sticks, 2 onions, 3 garlic cloves
- 400 g canned tomatoes, 1 L vegetable stock
- 200 g red lentils
- Handful of parsley, thyme, bay leaves
- Finish with lemon zest + olive oil`,
  },
  {
    id: 'weekend-dinner',
    title: 'Weekend date night pasta',
    tone: 'weekend',
    minutes: 35,
    text: `Rich mushroom pasta for two:
- 250 g tagliatelle
- 300 g mixed mushrooms, 1 shallot, 2 cloves garlic
- 150 ml cooking cream, splash of white wine
- 50 g parmesan, fresh thyme
- Simple salad: rocket, cherry tomatoes, balsamic`,
  },
]


