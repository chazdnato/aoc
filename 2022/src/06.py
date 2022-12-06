from utils.api import get_input

input_str = get_input(6)

# input_str = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"
# WRITE YOUR SOLUTION HERE

def test_day1_solution():
  assert day1_solution("mjqjpqmgbljsphdztnvjfqwrcgsmlb") == 7
  assert day1_solution("bvwbjplbgvbhsrlpgdmjqwftvncz") == 5
  assert day1_solution("nppdvjthqldpwncqszvftbrmjlhg") == 6
  assert day1_solution("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg") == 10
  assert day1_solution("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw") == 11
  return 0

def day1_solution(input_str):
  i = 0
  j = 4
  # go through the string
  while i < len(input_str):
    # stop if we're at the last 4
    if len(input_str[i:j]) < 4:
      break
    # set will have unique characters only
    batch = set(input_str[i:j])
    # if we have 4 unique, we're good to go
    if len(batch) == 4:
      break
    # next batch
    i += 1
    j += 1

  return j

print(day1_solution(input_str))
