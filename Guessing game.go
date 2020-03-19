package main

import (
   "fmt"
   "math/rand"
   "time"
)

func main() {
    fmt.Println("Guess the number !")

    // Generate a random number
    source := rand.NewSource(time.Now().UnixNano())
    randomizer := rand.New(source)
    HiddenValue := randomizer.Intn(50) // Generates numbers between 0 and n (50)

    var GuessedNumber int

    for {
        fmt.Println("Please input your guess: ")
        fmt.Scan(&GuessedNumber)
        fmt.Printf("You guessed %d\n", GuessedNumber)
        if GuessedNumber > HiddenValue {
          fmt.Println("That's too big, try a smaller one")
        } else if GuessedNumber < HiddenValue {
          fmt.Println("That's too small, try a bigger one")
        } else {
          fmt.Println("Nice one, You win")
          break
        }
    }
}
