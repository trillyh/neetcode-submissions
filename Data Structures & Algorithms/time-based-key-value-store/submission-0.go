type TimeMap struct {
    timeMap map[string][]pair
}

type pair struct {
    timestamp int
    value string
}

func Constructor() TimeMap {
    return TimeMap{
        timeMap: make(map[string][]pair),
    }
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    newPair := pair{timestamp, value}
    this.timeMap[key] = append(this.timeMap[key], newPair)
}

func (this *TimeMap) Get(key string, timestamp int) string {
    pairs, ok := this.timeMap[key]

    if !ok {
        return ""
    }
    l, r := 0, len(pairs)-1
    
    result := ""
    for l <= r {
        m := l + (r-l)/2
        if pairs[m].timestamp <= timestamp {
            result = pairs[m].value
            l = m + 1
        } else { // time after the desired timestamp
            r = m -1
        }
    }

    return result
}
