package main

func Sum(numbers []int) (sum int) {
  for _, num := range numbers {
    sum += num
  }
  return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
  for _, nums := range numbersToSum {
    sums = append(sums, Sum(nums))
  }
  return
}

func SumAllTails(numbersToSum ...[]int) (tailSums []int) {
  for _, nums := range numbersToSum {
    if len(nums) == 0 {
      tailSums = append(tailSums, 0)
    } else {
      tail := nums[1:]
      tailSums = append(tailSums, Sum(tail))
    }
  }
  return
}
