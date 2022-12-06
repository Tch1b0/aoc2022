with open("days/06/input.txt") as f:
    content = f.read()


def part_1():
    for i in range(0, len(content) - 3):
        x = [x for x in content[i:i+4]]
        if len(set(x)) == 4:
            print(i + 4)
            exit()

def part_2():
    for i in range(0, len(content) - 13):
        x = [x for x in content[i:i+14]]
        if len(set(x)) == 14:
            print(i + 14)
            exit()
