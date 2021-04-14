package genomicRangeQuery

func Solution(S string, P []int, Q []int) []int {
    s := []rune(S)
    M := len(P)
    results := make([]int, M)
    for i:=0; i<M; i++ {
        p := P[i]
        q := Q[i]
        if p == q {
            results[i] = typeToInt(s[p])
            continue
        } 

        minValue := typeToInt(s[p])
        for j:=p+1; j<=q; j++ {
            minValue = min(minValue, typeToInt(s[j]))
        }
        results[i] = minValue
    }
    return results
}

func typeToInt(t rune) int {
    switch t {
    case 'A':
        return 1
    case 'C':
        return 2
    case 'G':
        return 3
    case 'T':
        return 4
    }
    panic("should not be here")
}

func min(a int, b int) int{
    if a <= b {
        return a
    }
    return b
}
