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

# determine if an report is "safe"
def is_report_safe(report):
  array_len = len(report)
  for i in range(array_len):
    # reached the end without unsafe
    if i == array_len - 1:
      return True

    a, b = int(report[i]), int(report[(i + 1)])
    # determine if we're ascending or descending, first elements only
    if i == 0:
      if a < b:
        direction = 1
      else:
        direction = -1

    # if we switch directions or the difference is more than 3 or zero, it's unsafe
    if (b - a) * direction <= 0 or (b - a) * direction > 3:
      return False

  return True

safe = 0
safe2 = 0
for report in input_str.splitlines():
  report_array = report.strip().split(" ")

  array_len = len(report_array)
  if is_report_safe(report_array):
    safe += 1

  # take a single element out of the report, one at a time, then do all the same things from here
  is_safe = False
  for i in range(len(report_array)):
    # skip if this report has already been marked safe
    if is_safe:
      break
    report_array_copy = report_array.copy()
    report_array_copy.pop(i) # remove each element one at a time

    if is_report_safe(report_array_copy):
      is_safe = True
      safe2 += 1


print("Safe: ", safe)
print("Safe 2: ", safe2)
