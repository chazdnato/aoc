from utils.api import get_input
import re

input_str = get_input(3)

# input_str = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

# WRITE YOUR SOLUTION HERE

# Regular expression to match the pattern
mul_pattern = r"mul\(\d{1,3},\d{1,3}\)"
# Find all matches
matches = re.findall(mul_pattern, input_str)

# Regular expression to extract the numbers from each pattern
number_pattern = r"mul\((\d{1,3}),(\d{1,3})\)"

# Iterate through each match and process
products = []
for match in matches:
    match_groups = re.match(number_pattern, match)
    if match_groups:
        # Extract XX and YY as integers
        xx = int(match_groups.group(1))
        yy = int(match_groups.group(2))
        # Compute the product
        products.append(xx * yy)

# sum the products
print("part 1: ", sum(products))

# Part 2
# input_str="xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

# Regex to match `mul(XX,YY)` or `do()`/`don't()`
mul_pattern = r"mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)"

# Find all matches
matches = re.findall(mul_pattern, input_str)

# Process matches with state tracking
state = True  # Start with "on"
results = []

for match in matches:
    if match == "do()":
        state = True  # Turn matching on
    elif match == "don't()":
        state = False  # Turn matching off
    elif state:  # Only process `mul(XX,YY)` if "on"
        # Extract XX and YY using a separate regex
        number_pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
        match_groups = re.match(number_pattern, match)
        if match_groups:
            xx = int(match_groups.group(1))
            yy = int(match_groups.group(2))
            results.append(xx * yy)

print("part 2: ", sum(results))
