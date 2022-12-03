from utils.api import get_input

input_str = get_input(3)

def item_value(item):
  # ord minus 96 gives a-z = 1-26
  # ord minus 38 (aka minutes 64 + 26) gives A-Z = 27-52 
  return ord(item) - 96 if item.islower() else ord(item) - 38

#part 1
sum = 0

# part 2
sum2 = 0
sacks = 1
group = set()

for sack in input_str.splitlines():
  ## PART ONE
  size = int(len(sack) / 2)
  compartment_1 = set(sack[:size])
  compartment_2 = set(sack[size:])
  # extract single string from set
  item = compartment_1.intersection(compartment_2).pop()
  sum += item_value(item)

  ## PART TWO
  sack_set = set(sack[:])
  # first sack, add to empty group
  if sacks %3 != 0:
    # if there is an intersection ...
    if group & sack_set:
      group = group.intersection(sack_set)
    else:
      group = sack_set
    sacks += 1
  # third sack, intersection of set, reset sack/group
  else:
    item = group.intersection(sack_set).pop()
    sum2 += item_value(item)
    group = set()
    sacks = 1

print(sum, sum2)
