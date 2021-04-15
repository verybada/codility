package fish

func Solution(A []int, B []int) int {
    aliveFishes := newStack()
    upstreamFishes := newStack()

    length := len(A)
    for i:=0; i<length; i++ {
        fish := A[i]
        direction := B[i]

        if direction == 1 {
            upstreamFishes.Push(fish)
        } else if direction == 0 {
            beenEaten := false
            for upstreamFishes.Len() != 0 {
                upstreamFish := upstreamFishes.Last()
                if upstreamFish > fish {
                    beenEaten = true
                    break
                } else {
                    upstreamFishes.Pop()
                }
            }

            if beenEaten {
                continue
            }

            aliveFishes.Push(fish)
        }
    }

    return aliveFishes.Len() + upstreamFishes.Len()
}

func newStack() *stack {
    return &stack{fishes: []int{}}
}

type stack struct {
    fishes []int
}

func (s *stack) Push(fish int) {
    s.fishes = append(s.fishes, fish)
}

func (s *stack) Pop() int {
    l := len(s.fishes)
    fish := s.fishes[l-1]
    s.fishes = s.fishes[:l-1]
    return fish
}

func (s *stack) Last() int {
    l := len(s.fishes)
    return s.fishes[l-1]
}

func (s *stack) Len() int {
    return len(s.fishes)
}
