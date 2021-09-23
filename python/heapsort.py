from typing import List
from typing_extensions import Literal


class Heap(object):
    def __init__(self,array:List[int],order:Literal['max','min']) -> None:
        super().__init__()
        array.insert(0,0)
        self.array=array
        self.size=len(array)-1
        self.order=order
        self.sort()
    def parent(self,i):
        return i//2
    def left(self,i):
        return i*2
    def right(self,i):
        return i*2+1
    def swap(self,a,b):
        self.array[a],self.array[b]=self.array[b],self.array[a]
    def argmax(self,*indice):
        maxi=indice[0]
        for i in indice:
            if i<=self.size and self.array[i]>self.array[maxi]:
                maxi=i
        return maxi
    def argmin(self,*indice):
        mini=indice[0]
        for i in indice:
            if i<=self.size and self.array[i]<self.array[mini]:
                mini=i
        return mini
    def maxHeapify(self,i):
        lc=self.left(i)
        rc=self.right(i)
        maxI=self.argmax(i,lc,rc)
        if maxI!=i:
            self.swap(i,maxI)
            self.maxHeapify(maxI)
    def minHeapify(self,i):
        lc=self.left(i)
        rc=self.right(i)
        minI=self.argmin(i,lc,rc)
        if minI!=i:
            self.swap(i,minI)
            self.minHeapify(minI)
    def buildHeap(self):
        n=self.size
        for i in range(n//2,0,-1):
            if self.order=='max':
                self.maxHeapify(i)
            else:
                self.minHeapify(i)
    def sort(self):
        self.buildHeap()
        for i in range(len(self.array)-1,1,-1):
            self.swap(i,1)
            self.size-=1
            if self.order=='max':
                self.maxHeapify(1)
            else:
                self.minHeapify(1)
        self.size=len(self.array)-1
    def pop(self):
        n=len(self.array)-1
        if n==0:
            return None
        ret=self.array[n]
        del self.array[n]
        return ret
heap=Heap([1,3,5,7,2,4,6,8],'min')
ret=[]
for i in range(4):
    ret.append(heap.pop())
print(ret)