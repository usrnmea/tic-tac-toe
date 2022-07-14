package game

import "fmt"

type Player interface {
	MakeMove(*Board, Token) Move
}

type (
	RealPlayer	struct {}
	Bot		struct {}
)

func (p *RealPlayer) MakeMove(board *Board, token Token) Move {
	var move string 
	var destination Square

	for {
		fmt.Print("Move: ")
		fmt.Scanln(&move)

		square, err := StringToSquare(move)

		destination = square

		if err == nil && board[destination] == Empty {
			break
		}

		fmt.Println("Bad move format! Example: B2")
	}

	return Move{
		Token: token,
		Destination: destination,
	}
}

func (p *Bot) MakeMove(board *Board, token Token) Move {
	ml := board.GetMoves(token)

	var move Move

	bestEval := -2

	for i := range ml {
		board.DoMove(ml[i])

		newEval := -negamax(board, token.Flip())

		if newEval > bestEval {
			move = ml[i]
			bestEval = newEval
		}

		board.UndoMove(ml[i])
	}

	return move
}

func evaluateEnd(board *Board) int {
	var sums [8]int

	sums[0] = int(board[A1] + board[A2] + board[A3])
	sums[1] = int(board[B1] + board[B2] + board[B3])
	sums[2] = int(board[C1] + board[C2] + board[C3])

	sums[3] = int(board[A1] + board[B1] + board[C1])
	sums[4] = int(board[A2] + board[B2] + board[C2])
	sums[5] = int(board[A3] + board[B3] + board[C3])

	sums[6] = int(board[A1] + board[B2] + board[C3])
	sums[7] = int(board[A3] + board[B2] + board[C1])

	for i := 0; i < 8; i++ {
		if sums[i] == 3 || sums[i] == 12 {
			return 1
		}
	}

	return 0
}

func negamax(board *Board, token Token) int {
	if board.IsEnd() {
		return -evaluateEnd(board)
	}

	ml := board.GetMoves(token)

	eval := -2

	for i := range ml {
		board.DoMove(ml[i])

		newEval := -negamax(board, token.Flip())

		if newEval > eval {
			eval = newEval
		}

		board.UndoMove(ml[i])
	}

	return eval
}
