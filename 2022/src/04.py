from utils.api import get_input

input_str = get_input(4)

def does_contain(segment1, segment2):
  s1, e1 = [int(x) for x in segment1.split("-")]
  s2, e2 = [int(x) for x in segment2.split("-")]
  # s1 contains s2
  if s1 <= s2 and e2 <= e1:
    return True
  # s2 contains s1
  if s2 <= s1 and e1 <= e2:
    return True
  return False

def does_overlap(segment1, segment2):
  s1, e1 = [int(x) for x in segment1.split("-")]
  s2, e2 = [int(x) for x in segment2.split("-")]
  # s1 overlaps s2
  if s1 <= s2 and e1 >= s2 or e1 >= s2 and e1 <= e2:
    return True
  # s2 overlaps  s1
  if s2 <= s1 and e2 >= s1 or e2 >= s1 and e2 <= e1:
    return True
  return False

count = 0
count2 = 0
for elves in input_str.splitlines():
  elf1, elf2 = elves.split(",")
  if does_contain(elf1, elf2):
    count += 1

  if does_overlap(elf1, elf2):
    count2 += 1

print(count)
print(count2)
