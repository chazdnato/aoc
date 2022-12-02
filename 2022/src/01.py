from utils.api import get_input

input_str = get_input(1)

# single highest calorie holder
highest = 0
# top three calorie holders
maximums = []
count = 0

for line in input_str.splitlines():
  # empty line means elf is done
  if line == "":
    highest = max(highest, count)
    # add to array total
    maximums.append(count)
    # reset rolling count
    count = 0
  else:
    count += int(line)

print("part one: " + str(highest))

maximums.sort()
print("part two: " + str(sum(maximums[-3:])))
