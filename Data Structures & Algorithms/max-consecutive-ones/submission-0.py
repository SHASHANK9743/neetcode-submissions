class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        counter = 0
        max_counter = -1
        for _, itm in enumerate(nums):
            if itm == 1:
                counter += 1
            else:
                max_counter = max(counter, max_counter)
                counter = 0
        max_counter = max(counter, max_counter)
        return max_counter