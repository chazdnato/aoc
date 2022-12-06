from utils.api import get_input
from parse import compile

# from more_itertools import batched

input_str = get_input(5)

content = input_str.splitlines()
# width of board (4 spaces per "slot")
number_of_stacks = int(len(content[0]) // 4)
stacks = [[] for i in range(number_of_stacks)]

end_result = ""

# get stacks
for line in content:
  if line[1] == "1":
    break
  for stack in range(number_of_stacks):
    # every 4th spot
    index = 4 * stack + 1
    # skip if we've got no item in the stack
    if line[index] == ' ':
      continue
    stacks[stack].append(line[index])

# using parse to extract data from the line
p = compile("move {:d} from {:d} to {:d}")
for line in content:
  if line[0] == "m":
    [amount, move_from, move_to] = p.parse(line)

    count = 0
    while count < amount:
      stacks[move_to - 1].insert(0, stacks[move_from - 1].pop(0))
      count += 1

# compile result
for stack in stacks:
  end_result += stack[0]

print(end_result)
