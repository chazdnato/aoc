from utils.api import get_input
from pprint import pprint


input_str = get_input(7)

class Path:
  def __init__(self, name, size, parent):
    self.name = name
    self.size = size
    self.parent = parent
    self.children = {}
    self.total_size = 0

# just for diagnostics
def print_tree(node, level):
  if node.parent == None:
    print("TOTAL SIZE " + str(node.total_size))
  for entry, child in node.children.items():
    if child.children:
      print("  " * level + "dir : " + entry + " (total size " + str(child.total_size) + ") (parent: " + child.parent.name + ")")
      level += 1
      print_tree(child, level)
      level -= 1
    else:
      print("  " * level + "file: " + entry + " size: " + str(child.size) + " (parent: " + child.parent.name + ")")

# calculate all the sizes and add total_size to each node
def calculate_sizes(node):
  # for each child ...
  for child in node.children.values():
    if child.children:
      # if a child has children, recurse
      node.total_size += calculate_sizes(child)
    else:
      # otherwise add the size of this node to the total size of the tree
      node.total_size += child.size
  return node.total_size

# for part 1
def find_sizes(node):
  # for each child ...
  for child in node.children.values():
    # if child has children, recurse
    if child.children:
      find_sizes(child)
  # if this node's total size is small enough, count it with a magic number
  if node.total_size <= LARGEST_DIRECTORY:
    part1.append(node.total_size)

# for part 2
def find_directory_sizes(node, used_space):
  if node.parent == None:
    part2.append(node.total_size)
  for child in node.children.values():
    if child.children:
      find_directory_sizes(child, used_space)
      # only consider directory if it's big enough, another magic number
      if child.total_size + used_space > DESIRED_SPACE:
        part2.append(child.total_size)

# initialise
root = Path("/", 0, None)
current_node = root

# Magic numbers
LARGEST_DIRECTORY = 100000
TOTAL_DISK_SIZE = 70000000
DESIRED_SPACE = 30000000

# populate tree
for line in input_str.splitlines():
  line = line.split()
  if line[0] == "$":
    # start at root or ls can be ignored
    if (len(line) > 2 and line[2] == "/") or line[1] == "ls":
      continue
    cmd = line[1]
    dir = line[2]
    # we have a command
    if dir == "..":
      current_node = current_node.parent
    else:
      current_node = current_node.children[dir]
  else:
    size = 0
    # this will be the filesize or string "dir"
    filesize_or_dir = line[0]
    filename = line[1]
    # we have files or directories, we can ignore directories
    if filesize_or_dir != "dir":
      size = int(filesize_or_dir)

    # print("node child " + filename + " getting this path: " + filename, str(size))
    current_node.children[filename] = Path(filename, size, current_node)


# solve 
total_size = calculate_sizes(root)

part1 = []
find_sizes(root)
print(sum(part1))

part2 = []
used_space = TOTAL_DISK_SIZE - total_size # magic number
find_directory_sizes(root, used_space)
print(min(part2))
