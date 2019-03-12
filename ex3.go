package main

func Cheat(guess chan Choice) chan Choice {
	choiceSync := make(chan Choice) // Create channel to sync Me and Opponent choice
	go func() {
		for {
			// Waiting to receive 2 choices, one choice is Opponent choice and other one is Me
			choice1 := <-guess
			choice2 := <-guess
			if choice1.Who == 1 { // Is choice1 the Opponent choice? then change Me choice
				if choice1.Guess == SCISSORS {
					choice2 = Choice{0, ROCK}
				} else if choice1.Guess == ROCK {
					choice2 = Choice{0, PAPER}
				} else if choice1.Guess == PAPER {
					choice2 = Choice{0, SCISSORS}
				}
			} else if choice2.Who == 1 { // Is choice2 the Opponent choice? then change Me choice
				if choice2.Guess == SCISSORS {
					choice1 = Choice{0, ROCK}
				} else if choice2.Guess == ROCK {
					choice1 = Choice{0, PAPER}
				} else if choice2.Guess == PAPER {
					choice1 = Choice{0, SCISSORS}
				}
			}
			choiceSync <- choice1
			choiceSync <- choice2
		}

	}()
	return choiceSync // returns channel that will receive Me updated choice and Opponent untouched choice.
}
