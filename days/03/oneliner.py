print(sum([(ord(x) - (38 if x.isupper() else 96)) for x in (line[: len(line) // 2])if x in line[len(line) // 2 :]][0] for line in open("days/03/input.txt").readlines()), sum([(ord(x) - (38 if x.isupper() else 96)) for x in a if x in b and x in c][0] for a, b, c in zip((x for i, x in enumerate(open("days/03/input.txt").readlines()) if i % 3 == 0),(x for i, x in enumerate(open("days/03/input.txt").readlines()) if i % 3 == 1), (x for i, x in enumerate(open("days/03/input.txt").readlines()) if i % 3 == 2))))