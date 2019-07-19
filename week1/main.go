package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func eval(text string) {

  parts := strings.Split(text, " ")

  if len(parts) != 3 {
    fmt.Println("Invalid input!")
    return
  }

  firstNum, err := strconv.ParseFloat(parts[0], 64)
  secondNum, err := strconv.ParseFloat(parts[2], 64)
  if err != nil {
    fmt.Println("Invalid input!")
  }

  operator := parts[1]

  switch operator {
  case "+":
    fmt.Println(firstNum, operator, secondNum, "=", firstNum+secondNum)
  case "-":
    fmt.Println(firstNum, operator, secondNum, "=", firstNum-secondNum)
  case "*":
    fmt.Println(firstNum, operator, secondNum, "=", firstNum*secondNum)
  case "/":
    if secondNum == 0 {
      fmt.Println("Cannot divide by zero!")
    } else {
      fmt.Println(firstNum, operator, secondNum, "=", firstNum/secondNum)
    }
  default:
    fmt.Println("Invalid input!")
  }
}

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("> ")
  for scanner.Scan() {
    text := scanner.Text()
    eval(text)
    fmt.Print("> ")
  }
}
