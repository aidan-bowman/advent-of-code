import array

def part1(nums_list):
    nums = array.array('i', nums_list)
    
    index = 0
    counter = 0
    while 0 <= index < nums.buffer_info()[1]:
        temp = nums[index]
        nums[index] += 1
        index = index + temp
        counter += 1
    return counter

def part2(nums_list):
    nums = array.array('i', nums_list)
    
    index = 0
    counter = 0
    while 0 <= index < nums.buffer_info()[1]:
        temp = nums[index]
        if temp >= 3:
            nums[index] -= 1
        else:
            nums[index] += 1
        index = index + temp
        counter += 1
    return counter

with open(input("Enter file name: ")) as file:
    nums = []
    for line in file:
        nums.append(int(line))
    print(f"Part 1 solution: {part1(nums)}")
    print(f"Part 2 solution: {part2(nums)}")
