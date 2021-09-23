from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        table={}
        maxLen=0
        for n in nums:
            if n not in table:
                table[n]=1
                lEdge=n-table[n-1] if n-1 in table else n
                rEdge=n+table[n+1] if n+1 in table else n
                lenInterval=rEdge-lEdge+1
                maxLen=lenInterval if lenInterval>maxLen else maxLen
                table[lEdge],table[rEdge]=lenInterval,lenInterval
        return maxLen
print(Solution().longestConsecutive([4,0,-4,-2,2,5,2,0,-8,-8,-8,-8,-1,7,4,5,5,-4,6,6,-3]))