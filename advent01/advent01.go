package main;

import (
  "fmt"
  "bufio"
  "os"
  "strconv"
)

var nums map[int]int

func main() {
  nums = make(map[int]int)
  scanner := bufio.NewScanner(os.Stdin);
  for scanner.Scan() {
    s := scanner.Text()
    num, err := strconv.Atoi(s)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    nums[num] = 1
  }

  part1()
  part2()
}

func part1() {
  for num, _ := range nums {
    second := 2020 - num
    _, ok:= nums[second]
    if ok {
      fmt.Println(num,"+",second,"= 2020, so multipled =",num * second)
      return;
    }
  }
  fmt.Println("ERR: none matched?")
}

func part2() {
  for num1, _ := range nums {
    for num2, _ := range nums {
      if num1 == num2 || 2020-num1-num2 == num1 || 2020-num1-num2 == num2 {
        continue
      }
      third := 2020 - num1 - num2
      _, ok:= nums[third]
      if ok {
        fmt.Println(num1,"+",num2,"+",third,"= 2020, so multipled =",num1 * num2 * third)
        return;
      }
    }
  }
  fmt.Println("ERR: none matched?")
}
