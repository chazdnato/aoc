from utils.api import get_input

from pprint import pprint

input_str = get_input(8)

input_str = """30373
25512
65332
33549
35390"""

forest = []
visibleTrees = set()
for line in input_str.splitlines():
  numList = [*line]
  forest.append(numList)


def parse_left_to_right(forest):
  for i, row in enumerate(forest):

    # first get edges
    row_width = len(row)
    for j in range(row_width):
      grid_tuple = (i, j)
      if i == 0 or i == len(forest) - 1 or j == 0 or j == row_width - 1:
        visibleTrees.add(grid_tuple)


    # now get inner chunks

    # skip first row as we've already counted it
    if i == 0:
      continue

    for j in range(row_width):
      # skip the first item
      if j == 0:
        continue
      grid_tuple = (i, j)
      if row[j] > row[j-1]:
        print("adding (" + str(i) + "," + str(j) + ") = " + str(row[j]))
        visibleTrees.add(grid_tuple)
        break

pprint(forest)
parse_left_to_right(forest)

rotated1 = list(reversed(list(zip(*forest))))
pprint(rotated1)
parse_left_to_right(rotated1)

rotated2 = list(reversed(list(zip(*rotated1))))
pprint(rotated2)
parse_left_to_right(rotated2)

rotated3 = list(reversed(list(zip(*rotated2))))
pprint(rotated3)
parse_left_to_right(rotated3)

# pprint(visibleTrees)
print(len(visibleTrees))
