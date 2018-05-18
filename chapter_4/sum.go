package main

func Sum(numbers [5]int) (sum int) {
  for _, num := range numbers {
    sum += num
  }
  return
}
