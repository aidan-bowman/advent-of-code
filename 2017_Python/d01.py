def p1_sum(string):
    sum = 0;
    for pos in range(-1, len(string) - 1):
        if string[pos] == string[pos + 1]:
           sum += int(string[pos])
    return sum

def p2_sum(string):
    sum = 0
    length = len(string)
    increment = length >> 1
    # only run through first half
    for pos in range(int(increment)):
        if string[pos] == string[(pos + increment) % length]:
            sum += int(string[pos])
    # double output to count second half
    return (sum << 1)

with open(input("Enter file name: ")) as file:
    contents = file.read().strip()
    print(f"Part 1 solution: {p1_sum(contents)}")
    print(f"Part 2 solution: {p2_sum(contents)}")
