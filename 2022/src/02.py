from utils.api import get_input

input_str = get_input(2)

# input_str = """A Y
# B X
# C Z"""

# ROCK:     A X
# PAPER:    B Y
# SCISSORS: C Z

# A < B, B < C, C < A

player1_moves = {
  "A": "rock",
  "B": "paper",
  "C": "scissors",
}

player2_moves = {
  "X": "rock",
  "Y": "paper",
  "Z": "scissors",
}

move_scores = {
  "X":  1,
  "Y":  2,
  "Z": 3
}

throw_scores = {
  "rock":  0,
  "paper":  3,
  "scissors": 6
}

# did the RPS round result in a win or loss?
def rock_paper_scissors(move1, move2):
  # player 1 plays rock
  if move1 == "rock":
    if move2 == "scissors":
      return True
    else:
      return False
  # player 1 plays paper
  if move1 == "paper":
    if move2 == "rock":
      return True
    else:
      return False
  # player 1 plays scissors
  if move1 == "scissors":
    if move2 == "paper":
      return True
    else:
      return False

# calculate score based on RPS win/draw/loss
def game_score(move1, move2):
  # draw
  if move1 == move2:
    return 3
  # loss if player 1 wins
  if rock_paper_scissors(move1, move2):
    return 0
  # win
  return 6

def what_move_to_make(move1, result):
  # iterate through possible moves until you get the score you want
  if game_score(move1, "rock") == result:
    return move_scores["X"]
  if game_score(move1, "paper") == result:
    return move_scores["Y"]
  if game_score(move1, "scissors") == result:
    return move_scores["Z"]

# must play appropriate move to get the result
# and score that appropriate move appropriately
def throw_score(move1, move2):
  score = 0
  # mandated result
  score += throw_scores[move2]
  # determine which move to get the desired result and score it
  score += what_move_to_make(move1, score)
  return score

part1 = 0
part2 = 0
for line in input_str.splitlines():
  [move1, move2] = line.split(" ")
  part1 += game_score(player1_moves[move1], player2_moves[move2]) + move_scores[move2]
  part2 += throw_score(player1_moves[move1], player2_moves[move2])

print(part1)
print(part2)
