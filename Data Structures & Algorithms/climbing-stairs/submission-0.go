func climbStairs(n int) int {
    cache := make([]int, n+1)

    for i := range cache {
        cache[i] = -1
    }



    var step func(int) int
    step = func(at int) int {
        if at == n {
            return 1
        }
        if at > n {
            return 0
        }

        if cache[at] != -1 {
            return cache[at]
        }

        stepAt := step(at+1) + step(at+2)
        cache[at] = stepAt
        return stepAt
    }
    res := step(0)
    return res
}