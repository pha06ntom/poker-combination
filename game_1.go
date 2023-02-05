package main

import (
"fmt"
"sort"
)

// Определить, возможна ли комбинация из стрита (покер)

// Настройка функции
func isHandOfStraights(hand []int, k int) bool {

 if len(hand) % k != 0 {  // Проверка, делится ли кол-во карт на k
 return false
 }

 count := make(map[int]int)  // Подсчет кол-ва появлений каждой карты в руке
 for _, i := range hand {
  count[i] = count[i]+1
 }

 sort.Ints(hand)  // Сортирока списка и просмотр его с карты с самым низким значением
 i := 0
 n := len(hand)

 for i<n {
  current := hand[i]
  for j := 0; j<k; j++ {

 // Проверка, есть ли в подсчете текущая карта и следующие карты k-1 (в возрастающем рейтинге). Если какой-либо из них не существует, верните значение false.
 if _,ok := count[current + j]; !ok || count[current+j] == 0 {
 return false
 }
 count[current+j]--  // когда каждая из требуемых карт найдена, уменьшаем кол-во ее в count
 }

 // После того, как найдена полная группа, используем цикл for, чтобы найти наименьшую карту следующей группы и определить, у какой из следующих карт в подсчете осталось больше нуля вхождений.
 for i<n && count[hand[i]] ==0 {
 i++
 } 
 }

return true  // возвращаем true, если все карты отсортированы по группам
}

func main() {

hand := []int{5,2,4,7,1,3,5,6,3}
k := 3
fmt.Println(isHandOfStraights(hand, k))
hand2 := []int{1,9,3,5,7,4,2,9,11}
k = 2
fmt.Println(isHandOfStraights(hand2, k))
}