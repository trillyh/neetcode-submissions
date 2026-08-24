class KthLargest:
    minQ: List[int]
    k: int
    def __init__(self, k: int, nums: List[int]):
        self.minQ = nums
        self.k = k
        heapq.heapify(self.minQ)
        while len(self.minQ) > k:
            heapq.heappop(self.minQ)

    def add(self, val: int) -> int:
        heapq.heappush(self.minQ, val)
        while len(self.minQ) > self.k:
            heapq.heappop(self.minQ)
        return self.minQ[0]
        
