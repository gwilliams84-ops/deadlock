package main

import "fmt"

type DynamicGateway struct {
    state int
}

func (s *DynamicGateway) build_loader(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*20) % 997
    }
    return total
}

func main() {
    obj := &DynamicGateway{state: 20}
    fmt.Println(obj.build_loader(20))
}
