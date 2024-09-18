package main

import (
    "fmt"
)

func IsCapitalized(s string) bool {
    start := 0
    end := 0

    for i, r := range s {
        if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
            end = i
            if !isCapitalizedWord(s[start:end]) {
                return false
            }
            start = i + 1
        }
    }

    if start < len(s) {
        if !isCapitalizedWord(s[start:]) {
            return false
        }
    }
    if s == "" {
        return false
    }
    return true
}

func isCapitalizedWord(s string) bool {
    if len(s) == 0 {
        return false
    }
    r := s[0]
    return r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '!'
}

func main() {
    fmt.Println(IsCapitalized("Hello! How are you?"))
    fmt.Println(IsCapitalized("Hello How Are You"))
    fmt.Println(IsCapitalized("Whats 4this 100K?"))
    fmt.Println(IsCapitalized("Whatsthis4"))
    fmt.Println(IsCapitalized("!!!!Whatsthis4"))
    fmt.Println(IsCapitalized(""))
}