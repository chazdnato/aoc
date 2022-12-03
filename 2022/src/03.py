from utils.api import get_input

input_str = get_input(3)

def item_value(item):
  if ord(item) > 96:
    return ord(item) - 96
  else:
    return ord(item) - 38

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
  (item, ) = compartment_1.intersection(compartment_2)
  # ord minus 96 means a-z = 1-26
  # ord minus 38 means A-Z = 27-52
  sum += item_value(item)

  ## PART TWO
  sack_set = set(sack[:])
  # first sack, add to group
  if sacks == 1:
    group = sack_set
    sacks += 1
  # second sack, intersection
  elif sacks == 2:
    group = group.intersection(sack_set)
    sacks += 1
  # third sack, intersection of set, reset sack/group
  elif sacks == 3:
    (item, ) = group.intersection(sack_set)
    sum2 += item_value(item)
    group = set()
    sacks = 1

print(sum, sum2)
