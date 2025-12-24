package utils

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Input prints a prompt and returns the user's input as a string.
func Input(prompt string) string {
fmt.Print(prompt)
text, _ := reader.ReadString('\n')
return strings.TrimSpace(text)
}

// InputInt prints a prompt and returns an integer.
// Returns 0 if the input is invalid.
func InputInt(prompt string) int {
str := Input(prompt)
num, err := strconv.Atoi(str)
if err != nil {
fmt.Println("Invalid number, defaulting to 0")
return 0
}
return num
}

// InputFloat prints a prompt and returns a float64.
// Returns 0.0 if the input is invalid.
func InputFloat(prompt string) float64 {
str := Input(prompt)
num, err := strconv.ParseFloat(str, 64)
if err != nil {
fmt.Println("Invalid number, defaulting to 0.0")
return 0.0
}
return num
}

// InputBool prints a prompt and returns a boolean.
// Accepts: "true", "yes", "y", "1" as true (case insensitive).
func InputBool(prompt string) bool {
str := strings.ToLower(Input(prompt))
return str == "true" || str == "yes" || str == "y" || str == "1"
}
