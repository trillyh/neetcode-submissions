class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        frequencies = {}
        for task in tasks:
            frequencies[task] = frequencies.get(task, 0) + 1
        
        heap = [freq for task, freq in frequencies.items()]

        heapq.heapify_max(heap)
        queue = deque()

        time = 0
        while heap or queue:
            time += 1
            
            if heap:
                count = heapq.heappop_max(heap) - 1
                if count != 0:
                    queue.append((count, time+n))
            if queue and queue[0][1] == time:
                heapq.heappush_max(heap, queue.popleft()[0])

        return time

        
