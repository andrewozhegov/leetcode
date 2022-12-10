#!/opt/homebrew/bin/python3

from typing import List

def singleNumber(nums: List[int]) -> int:
    one, two = 0, 0
    for i in range(0, len(nums)):
        one = (one ^ nums[i]) & ~two
        two = (two ^ nums[i]) & ~one
    return one + two

# inp = [2,2,3,2]
inp = [3,2,3,3,3]

print(str(singleNumber(inp)))
