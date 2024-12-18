from utils.api import get_input

input_str = get_input(2)


# input_str = """7 6 4 2 1
# 1 2 7 8 9
# 9 7 6 2 1
# 1 3 2 4 5
# 8 6 4 4 1
# 1 3 6 7 9
# """

# WRITE YOUR SOLUTION HERE
safe = 0
direction = 0
for report in input_str.splitlines():
  array = report.strip().split(" ")

  array_len = len(array)
  for i in range(array_len):
    # reached the end without unsafe
    if i == array_len - 1:
      safe += 1
      break

    a, b = int(array[i]), int(array[(i + 1)])
    # determine if we're ascending or descending, first elements only
    if i == 0:
      if a < b:
        direction = 1
      else:
        direction = -1

    # if we switch directions or the difference is more than 3 or zero, it's unsafe
    if (b - a) * direction <= 0 or (b - a) * direction > 3: 
      break

print("Safe: ", safe)
