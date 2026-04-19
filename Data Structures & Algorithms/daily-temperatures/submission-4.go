
type pair struct {
    idx int
    temp int
}

func dailyTemperatures(temperatures []int) []int {
    res := make([]int, len(temperatures))
    stack := []pair{}

    for i, t := range temperatures {
        for len(stack) > 0 && t > stack[len(stack)-1].temp {
            stackIdx := stack[len(stack)-1].idx
            stack = stack[:len(stack)-1]
            res[stackIdx] = i - stackIdx
        }
        stack = append(stack, pair{idx: i, temp: t})
    }
    return res

}
