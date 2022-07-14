package game

import "fmt"

func NewGame(crosses Player, noughts Player) {
	var board Board

	for i := 0; !board.IsEnd(); i++ {
		fmt.Println(board)

		if i % 2 == 0 {
			board.DoMove(crosses.MakeMove(&board, Cross))
			continue
		}

		board.DoMove(noughts.MakeMove(&board, Nought))
	}

	fmt.Println(board)
}
