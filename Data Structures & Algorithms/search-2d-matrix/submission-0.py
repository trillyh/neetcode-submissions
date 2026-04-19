class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        top, bottom = 0, len(matrix)-1
        lastIdx = len(matrix[0])-1

        while top <= bottom:
            central = top + ((bottom-top)//2)
            print(central)
            if  matrix[central][0] <= target and target <= matrix[central][lastIdx]:
                l, r = 0, len(matrix[0])-1
                while l <= r:
                    m = l + ((r-l)//2)
                    if  target == matrix[central][m]:
                        return True

                    elif  matrix[central][m] < target:
                        l = m + 1
                    else:
                        r = m - 1

                return False
            elif matrix[central][0] < target:
                top = central + 1

            else:
                print("central > target")
                bottom = central - 1
        return False 

            

