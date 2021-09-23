from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        row=[[False]*9]*10
        column=[[False]*9]*10
        subgrids=[[[False]*3]*3]*10
        for i in range(9):
            for j in range(9):
                num=int(board[i][j]) if board[i][j]!='.' else 0
                if num!=0:
                    temp=i//3
                    temp=j//3
                    if row[i][num]==True or column[j][num]==True or subgrids[i//3][j//3][num]==True:
                        return False
                    row[i][board[i][j]]=True
                    column[j][board[i][j]]=True
                    subgrids[i//3][j//3][board[i][j]]=True
        return True
print(Solution().isValidSudoku([["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]))