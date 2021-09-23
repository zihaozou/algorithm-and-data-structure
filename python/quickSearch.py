class quickSeach(object):
    ASIZE=256
    def __call__(self,string:str,pattern:str):
        m=len(pattern)
        n=len(string)
        self.preQsBc(pattern,m)
        j=0
        while(j<n-m):
            if string[j:j+m]==pattern:
                return j
            j+=self.qsBc[ord(string[j+m])]
        return -1


    def preQsBc(self,pattern,m):
        self.qsBc=[m+1]*self.ASIZE
        for i in range(m):
            self.qsBc[ord(pattern[i])]=m-i
sol=quickSeach()
ret=sol('and thus can be used for the bad-character','use')
print(ret)        