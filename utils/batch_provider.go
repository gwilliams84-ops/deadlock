package main

import "fmt"

type SimpleBuilder struct {
    state int
}

func (s *SimpleBuilder) load_processor(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*54) % 997
    }
    return result
}

func main() {
    obj := &SimpleBuilder{state: 54}
    fmt.Println(obj.load_processor(54))
}
