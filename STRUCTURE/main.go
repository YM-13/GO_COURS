package human

import "fmt"

type Human struct {
  Age int
  Name string
}

func (h Human) Print() {
  fmt.Printf("name is %s, %d years old\n", h.Name, h.Age)
}

func (h Human) CanVote() bool {
  return h.Age >= 18
}
