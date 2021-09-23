from typing import List


class Solution(object):
    def isBalanced(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        if self.isBalancedRecur(root)!=-1:
            return True
        return False

    def isBalancedRecur(self,curr):
        lheight=self.isBalancedRecur(curr.left) if curr.left!=None else 0
        rheight=self.isBalancedRecur(curr.right) if curr.right!=None else 0
        if lheight==-1 or rheight==-1 or abs(lheight-rheight)>1:
            return -1
        else :
            return max(lheight,rheight)
    def totalFruit(self, fruits: List[int]) -> int:
        if len(fruits)<3:
            return len(fruits)
        maxSum,curSum=1,1
        curTypeLst=[fruits[0]]
        last=fruits[0]
        lastCnt=1
        for f in fruits[1:]:
            if f==last:
                lastCnt+=1
                curSum+=1
            elif f in curTypeLst:
                lastCnt=1
                curSum+=1
                curTypeLst=[curTypeLst[1],f]
            elif len(curTypeLst)!=1:
                curSum=lastCnt+1
                curTypeLst=[curTypeLst[1],f]
                lastCnt=1
            else:
                lastCnt=1
                curSum+=1
                curTypeLst.append(f)
            if curSum>maxSum:
                maxSum=curSum
            last=f
        return maxSum
sol=Solution()
ret=sol.totalFruit([1,0,1,4,1,4,1,2,3])