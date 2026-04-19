func groupAnagrams(strs []string) [][]string {
    res := make(map[[26]int][]string)
    for _, str := range strs {

        freq := [26]int{}
        for _, c := range str {
            freq[c - 'a']++
        }
        res[freq] = append(res[freq], str)
    }
    var result [][]string
    for _, group := range res {
        result = append(result, group)
    }
    return result

}
