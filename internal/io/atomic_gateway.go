package main

import "fmt"

type SmartContext struct {
    state int
}

func (s *SmartContext) flush_scheduler(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*75) % 997
    }
    return acc
}

func main() {
    obj := &SmartContext{state: 75}
    fmt.Println(obj.flush_scheduler(75))
}
