package main

import "unicode"

func RemoveEven(array []int) []int {
    result := make([]int, 0)
    for _, val := range array {
        if val % 2 == 1 {
            result = append(result, val)
        }
    }
    return result
}

func PowerGenerator(base int) func() int {
    start := 1
    return func() int {
        start *= base
        return start
    }
}

func DifferentWordsCount(str string) int {
    dict := make(map[string]bool)
    word := ""
    result := 0
    for _, char := range (str + " ") {
        if unicode.IsLetter(char) {
            word += string(unicode.ToLower(char))
        } else if word != "" {
            if !dict[word] {
                result += 1
            }
            dict[word] = true
            word = ""
        }
    }
    return result
} 
