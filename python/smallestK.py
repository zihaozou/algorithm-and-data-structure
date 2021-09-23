from typing import List


class Solution:
    def smallestK(self, arr: List[int], k: int) -> List[int]:
        if len(arr)==0:
            return []
        self.array=arr
        self.quickFind(0,len(arr)-1,k)
        return self.array[:k]
    def quickFind(self,l,r,k)->int:
        if l==r:
            return l
        p=self.partition(l,r)
        if p==k-1:
            return p
        elif k-1<p:
            return self.quickFind(l,p-1,k)
        else:
            return self.quickFind(p+1,r,k)
    def partition(self,l,r)->int:
        i=l
        for j in range(l,r,1):
            if self.array[j]<self.array[r]:
                self.array[j],self.array[i]=self.array[i],self.array[j]
                i+=1
        self.array[r],self.array[i]=self.array[i],self.array[r]
        return i
sol=Solution()
print(sol.smallestK([1,3,5,7,2,4,6,8],4))           
        