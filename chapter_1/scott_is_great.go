package main

import "fmt"

const spanish = "Spanish"
const spanishIsGreat = " es fantástica!"

const french = "French"
const frenchIsGreat = " est Remarquable!"

const pigLatin = "Pig Latin"
const pigLatinIsGreat = " isway ate-gray!"

const isGreat = " is Great!"
const defaultName = "Scott"

func Hello(name string, language string) string {
  if name == "" {
    name = defaultName
  }

  return name + greatSuffix(language)
}

func greatSuffix(language string) (suffix string) {
  switch language {
  case spanish:
    suffix = spanishIsGreat
  case french:
    suffix = frenchIsGreat
  case pigLatin:
    suffix = pigLatinIsGreat
  default:
    suffix = isGreat
  }
  return
}

func main() {
  fmt.Println(Hello("Scott", ""))
}
