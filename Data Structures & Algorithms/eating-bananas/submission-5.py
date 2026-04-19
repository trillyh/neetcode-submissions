class Solution:
    def minEatingSpeed(self, piles: List[int], h: int) -> int:
        # hours it takes to eat a pile = ceil(piles[i] // k)
        # min(piles) is the min, max = max(piles)
        totalPiles =  0
        
        minSpeed, maxSpeed = 1, max(piles)

        result = maxSpeed
        while minSpeed <= maxSpeed:
            currSpeed = minSpeed + (maxSpeed - minSpeed) // 2
            print(currSpeed)

            hoursToFinish = 0 
            for pile in piles:
                hoursToFinish += math.ceil(pile / currSpeed)
                print(f"hours to finish {hoursToFinish} ")
            
            if hoursToFinish <= h:
                result = min(result, currSpeed)
                maxSpeed = currSpeed - 1
            
            else:
                minSpeed = currSpeed + 1
        
        return result

                


        