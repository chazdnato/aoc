from utils.api import get_input

input_str = get_input(1)

# single highest calorie holder
highest = 0
# top three calorie holders
maximums = [0, 0, 0]
count = 0

def calculate_maximums(sum, maximums):
  if sum >= min(maximums):
    # replace lowest maximums with sum
    sorted_maximums = sorted(maximums)
    sorted_maximums[0] = max(sorted_maximums[0], sum)
    maximums = sorted_maximums
  return maximums

for line in input_str.splitlines():

  # empty line means elf is done
  if line == "":
    highest = max(highest, count)

    # calculate the three highest
    maximums = calculate_maximums(count, maximums)

    # reset rolling count
    count = 0
  else:
    count += int(line)

# need to calculate the last elf separately
maximums = calculate_maximums(count, maximums)

print("part one: " + str(highest))

print("part two: " + str(sum(maximums)))
