class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        minHeap = []
        for point in points:
            x = point[0]
            y = point[1]
            distance = math.sqrt((x**2) + (y**2))
            minHeap.append([distance, x, y])

        heapq.heapify(minHeap)

        res = []

        for i in range(k):
            dist, x, y = heapq.heappop(minHeap)
            res.append([x,y])
        return res