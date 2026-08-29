package main

import "fmt"

type BatchHandler struct {
    state int
}

func (s *BatchHandler) flush_session(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*36) % 997
    }
    return total
}

func main() {
    obj := &BatchHandler{state: 36}
    fmt.Println(obj.flush_session(36))
}
