from utils.api import get_input

input_str = get_input(6)

# input_str = "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

def test_day1_solution():
  batch_size = 4
  assert solution("mjqjpqmgbljsphdztnvjfqwrcgsmlb", batch_size) == 7
  assert solution("bvwbjplbgvbhsrlpgdmjqwftvncz", batch_size) == 5
  assert solution("nppdvjthqldpwncqszvftbrmjlhg", batch_size) == 6
  assert solution("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", batch_size) == 10
  assert solution("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", batch_size) == 11
  return 0

def test_day2_solution():
  batch_size = 14
  assert solution("mjqjpqmgbljsphdztnvjfqwrcgsmlb", batch_size) == 19
  assert solution("bvwbjplbgvbhsrlpgdmjqwftvncz", batch_size) == 23
  assert solution("nppdvjthqldpwncqszvftbrmjlhg", batch_size) == 23
  assert solution("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", batch_size) == 29
  assert solution("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", batch_size) == 26

def solution(input_str, batch_size):
  i = 0
  j = batch_size
  # go through the string
  while i < len(input_str):
    # stop if we're at the end of the batch
    if len(input_str[i:j]) < batch_size:
      break
    # set will have unique characters only
    batch = set(input_str[i:j])
    # if we have unique set, we're good to go
    if len(batch) == batch_size:
      break
    # next batch
    i += 1
    j += 1

  return j

print(solution(input_str, 4))
print(solution(input_str, 14))
