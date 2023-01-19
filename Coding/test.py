import random

# Create the game board with 10 spaces for each player
player1_board = ["_"] * 10
player2_board = ["_"] * 10

# Set the number of lives for each player
player1_lives = 5
player2_lives = 5

# Function to check if a player has lost
def check_loss(player_lives):
    if player_lives == 0:
        return True
    else:
        return False

# Main game loop
while True:
    # Player 1's turn
    print("Player 1's turn")
    player1_guess = int(input("Guess a space (0-9): "))
    if player2_board[player1_guess] == "X":
        print("Hit!")
        player2_board[player1_guess] = "H"
    else:
        print("Miss!")
        player2_board[player1_guess] = "M"
        player1_lives -= 1

    # Check if Player 1 has lost
    if check_loss(player1_lives):
        print("Player 2 wins!")
        break

    # Player 2's turn
    print("Player 2's turn")
    player2_guess = random.randint(0, 9)
    if player1_board[player2_guess] == "X":
        print("Hit!")
        player1_board[player2_guess] = "H"
    else:
        print("Miss!")
        player1_board[player2_guess] = "M"
        player2_lives -= 1

    # Check if Player 2 has lost
    if check_loss(player2_lives):
        print("Player 1 wins!")
        break