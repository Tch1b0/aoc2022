import re

with open("days/05/input.txt") as f:
    content = f.readlines()

s = []

for x in [x.replace("    "," [ ] ").strip() for x in content if "move" not in x and x != "\n" and "1" not in x]:
    for i, y in enumerate(x.split("[")):
        if len(s) - 1 < i:
            s.append([])

        y = y.replace("]", "").strip()
        if y != "":
            s[i].append(y)

s = list(map(lambda x: x[::-1], s[1:]))

for x in [x.replace("move ", "").replace("from ", "").replace("to ", "").split(" ") for x in content if "move" in x]:
    n, fr, to = list(map(lambda x: x - 1, map(int, x)))
    s[to].append(s[fr][len(s[fr]) - n - 1:])
    s[fr] = s[fr][:len(s[fr]) - n - 1]

print([x[0] for x in s])
