class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        unique = set()
        for _, num in enumerate(nums):
            if num not in unique:
                unique.add(num)
            else:
                return True
        return False