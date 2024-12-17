from utils.api import get_input

input_str = get_input(1)

# input_str = """3   4
# 4   3
# 2   5
# 1   3
# 3   9
# 3   3
# """

# WRITE YOUR SOLUTION HERE
array_a = []
array_b = []
for line in input_str.splitlines():
    a, b = line.strip().split("  ")
    array_a.append(int(a))
    array_b.append(int(b))


array_a.sort()
array_b.sort()

total = 0
for i, x in enumerate(array_a):
  total += abs(array_b[i] - x)

print(total)
