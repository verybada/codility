package maxprofit

func Solution(A []int) int {
    if len(A) == 0 {
        return 0
    }
    
    maxProfit := 0
    lowestPrice := A[0]
    highestPrice := A[0]
    for index, price := range A {
        if index == 0 {
            continue
        }

        if price < lowestPrice {
            lowestPrice = price
            highestPrice = -1
        } else if price > highestPrice {
            highestPrice = price
            maxProfit = max(maxProfit, highestPrice - lowestPrice)
        }
    }

    return maxProfit
}

func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b

