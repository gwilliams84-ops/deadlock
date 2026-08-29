package main

import "fmt"

type SmartMonitor struct {
    state int
}

func (s *SmartMonitor) resolve_cache(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*41) % 997
    }
    return value
}

func main() {
    obj := &SmartMonitor{state: 41}
    fmt.Println(obj.resolve_cache(41))
}
