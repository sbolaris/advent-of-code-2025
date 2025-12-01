#imports 
import os

#example inputs for the lock 
document = [
    "L68",
    "L30",
    "R48",
    "L5",
    "R60",
    "L55",
    "L1",
    "L99",
    "R14",
    "L82"          
]

#safe_cracker function
def safe_cracker(starting_position, document):
    #logic to crack the safe
    password = 0
    combo = []
    current_position = starting_position
    # lock as 100 positions (0-99)
    # if left at <0 go to 99
    # if right at >99 go to 0
    #L for left (lower numbers), R for right (higher number)
    for line in document:
        direction = line[0]
        steps = int(line[1:])
        
        if direction == "L":
            current_position -= steps
            if current_position < 0:
                current_position += 100
        elif direction == "R":
            current_position += steps
            if current_position > 99:
                current_position -= 100
        
        combo.append(current_position)

    #password is the number of zeos in the combo
    for number in combo:
        if number == 0:
            password += 1

    return password

def main():
    starting_position = 50
    with open("input.txt", "r") as file:
        actual_document = [line.strip() for line in file.readlines()]
    password = safe_cracker(starting_position, actual_document)
    print("The password to the safe is:", password)

if __name__ == "__main__":
    main()