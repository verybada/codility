package brackets

import (
    "fmt"
)

func Solution(S string) int {
    expectedLeftBrackets := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    s := newStack()
    for _, char := range S {
        switch char {
        case '(', '[', '{':
            s.Push(char)
            continue
        }

        last, err := s.Pop()
        if err != nil {
            return 0
        }
        expected := expectedLeftBrackets[char]
        if last != expected {
            return 0
        }
    }

    if s.Len() != 0 {
        return 0
    }
    return 1
}

func newStack() *stack {
    return &stack {
        s: make([]rune, 0),
    }
}

type stack struct {
    s []rune
}

func (s *stack) Push(r rune) {
    s.s = append(s.s, r)
}

func (s *stack) Pop() (rune, error) {
    stackLen := s.Len()
    if stackLen == 0 {
        return 0, fmt.Errorf("empty")
    }
    r := s.s[s.Len()-1]
    s.s = s.s[:s.Len()-1]
    return r, nil
}

func (s *stack) Len() int {
    return len(s.s)
}
 
