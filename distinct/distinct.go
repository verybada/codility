package distinct

func Solution(A []int) int {
    temp := make(map[int]bool)
    for _, a := range A {
        temp[a] = true
    }
    
    return len(temp)
}
