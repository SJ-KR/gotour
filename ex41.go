package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    split := strings.Fields(s)
    m := make(map[string]int)

    for i := 0; i < len(split); i++ {
        m[split[i]]++
    }
    return m
}

func main() {
    wc.Test(WordCount)
}

