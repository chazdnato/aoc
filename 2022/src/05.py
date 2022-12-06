from utils.api import get_input
from parse import compile

# from more_itertools import batched

input_str = get_input(5)

content = input_str.splitlines()
# width of board (4 spaces per "slot")
number_of_stacks = int(len(content[0]) // 4)
stacks = [[] for i in range(number_of_stacks)]
stacks2 = [[] for i in range(number_of_stacks)]

end_result = ""
end_result2 = ""

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
    stacks2[stack].append(line[index])

# using parse to extract data from the line
p = compile("move {:d} from {:d} to {:d}")
for line in content:
  if line[0] == "m":
    [amount, move_from, move_to] = p.parse(line)

    count = 0
    temp = []
    while count < amount:
      # part 1
      stacks[move_to - 1].insert(0, stacks[move_from - 1].pop(0))

      # part 2
      temp.append(stacks2[move_from - 1].pop(0))
      if count + 1 == amount:
        temp.reverse()
        for element in temp:
          stacks2[move_to - 1].insert(0, element)
          temp = []

      count += 1


# compile results
for stack in stacks:
  end_result += stack[0]

for stack in stacks2:
  end_result2 += stack[0]

print(end_result)
print(end_result2)
