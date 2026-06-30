package kata
​
​
func Cakes(recipe, available map[string]int) int {
  if len(recipe) == 0 || len(available) == 0 {
    return 0
  }
  maxCake := 0
  for ingredient, amountNeeded := range recipe {
    amountAvailable, exists := available[ingredient]
    if !exists {
      return 0
    }
    cakePossible := amountAvailable/amountNeeded
    if maxCake == 0 || cakePossible < maxCake {
      maxCake = cakePossible
    }
  }
  return maxCake
}
​