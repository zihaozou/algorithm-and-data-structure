from typing import List


class Solution:
    def findMaximizedCapital(self, k: int, w: int, profits: List[int], capital: List[int]) -> int:
        maxBProject={}
        for i in range(len(profits)):
            if capital[i] in maxBProject:
                maxBProject[capital[i]]=profits[i] if profits[i]>maxBProject[capital[i]] else maxBProject[capital[i]]
            else:
                maxBProject[capital[i]]=profits[i]
        for x in range(k):
            for y in range(w,-1,-1):
                if y in maxBProject:
                    w+=maxBProject[y]
                    break
        return w
print(Solution().findMaximizedCapital(10,0,[1,2,3],[0,1,2]))