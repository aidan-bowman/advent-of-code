def get_diff(nums):
    lowest = 9999999
    highest = 0
    for num in nums:
        if num > highest:
            highest = num
        if num < lowest:
            lowest = num
    return highest - lowest

def get_div_pair(nums):
    first = nums[0]
    for num in nums[1:]:
        if first % num == 0:
            return int(first/num)
        elif num % first == 0:
            return int(num/first)
    return get_div_pair(nums[1:])

with open(input("Enter file name: ")) as file:
    part1 = 0
    part2 = 0
    for line in file:
        nums = list(map(int, line.strip().split()))
        part1 += get_diff(nums)
        part2 += get_div_pair(nums)
    print(f"Part 1 solution: {part1}")
    print(f"Part 2 solution: {part2}")
