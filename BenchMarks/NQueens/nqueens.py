def solveNQueens(self, n):
    res = []
    self.dfs([-1]*n, 0, [], res)
    return res

# nums is a one-dimension array, like [1, 3, 0, 2] means
# first queen is placed in column 1, second queen is placed
# in column 3, etc.
def dfs(self, nums, index, path, res):
    if index == len(nums):
        res.append(path)
        return  # backtracking
    for i in xrange(len(nums)):
        nums[index] = i
        if self.valid(nums, index):  # pruning
            tmp = "."*len(nums)
            self.dfs(nums, index+1, path+[tmp[:i]+"Q"+tmp[i+1:]], res)

# check whether nth queen can be placed in that column
def valid(self, nums, n):
    for i in xrange(n):
        if abs(nums[i]-nums[n]) == n -i or nums[i] == nums[n]:
            return False
    return True

    
