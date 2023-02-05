package main

import (
"fmt"
"math"
)

// Найти максимальное количество очков, которое можно получить из набора из десяти случайных карт. 

// Настройка функции
func maxPoints(deck []int, k int) int{
left := 0
right := len(deck) - k
var total, best int
total = 0

// Предполагаем, что k карт с правой стороны дают нам максимальное кол-во очков
for i:= right; i < len(deck); i++ {
total += deck[i]
}
best = total

for i := 0; i<k; i++ {
total += deck[left] - deck[right]

best = int(math.Max(float64(best), float64(total)))
left++
right++
}
return best
}

func main() {
deck := []int{5,2,4,4,2,1,1,1,3}
k := 4
fmt.Println(maxPoints(deck, k))
}