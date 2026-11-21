class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        if len(nums) == 1:
            return nums[0]
        hashmap = {}
        for _, num in enumerate(nums):
            if num not in hashmap:
                hashmap[num] = 1
            else:
                hashmap[num] += 1
                if hashmap[num] > (len(nums) / 2):
                    return num