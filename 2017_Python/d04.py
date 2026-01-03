# list eater; if first word matches another word return false, otherwise check words[1:]
# pass in function to be used for validity checking
def check_passphrase_validity(words, equalfunc):
    if len(words) == 1:
        return True
    baseword = words[0]
    for word in words[1:]:
        if equalfunc(baseword, word):
            return False
    return check_passphrase_validity(words[1:], equalfunc)

# anagram checker
def part2_equals(worda, wordb):
    if len(worda) != len(wordb):
        return False
    # list eater: if first char has no match return false, otherwise check worda[1:] and wordb with char deleted
    def rec_check(worda, wordb):
        if len(worda) == 0:
            return True
        lettera = worda[0]
        for pos in range(len(wordb)):
            if lettera == wordb[pos]:
                return rec_check(worda[1:], wordb[:pos] + wordb[pos + 1:])
        return False
    return rec_check(worda, wordb)

with open(input("Enter file name: ")) as file:
    part1 = 0
    part2 = 0
    for line in file:
        words = line.split()
        if check_passphrase_validity(words, lambda x, y: x == y):
            part1 += 1
        if check_passphrase_validity(words, part2_equals):
            part2 += 1
    print(f"Part 1 solution: {part1}")
    print(f"Part 2 solution: {part2}")
