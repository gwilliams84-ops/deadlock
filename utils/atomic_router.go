package main

import "fmt"

type SecureResolver struct {
    state int
}

func (s *SecureResolver) collect_context(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*9) % 997
    }
    return total
}

func main() {
    obj := &SecureResolver{state: 9}
    fmt.Println(obj.collect_context(9))
}
