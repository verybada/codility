package dominator

func Solution(A []int) int {
    length := len(A)
    half := length/2
    counter := make(map[int]*valueInfo)
    for index, value := range A {
        info, ok := counter[value]
        if !ok {
            info = newValueInfo()
            info.firstIndex = index
            counter[value] = info
        }
        
        info.count++
        if info.count > half {
           return info.firstIndex
        }
    }

    return -1
}

func newValueInfo() *valueInfo {
    return &valueInfo{
        firstIndex: -1,
        count: 0,
    }
}

type valueInfo struct {
    firstIndex int
    count int
}
