package game

import (
	"fmt"
	"errors"
	"strings"
)

type (
	Square int
	Token int
)

type Move struct {
	Token Token
	Destination Square
}

type Board [SquareNumber]Token
type MoveList []Move

const (
	A1 Square = iota
	A2
	A3
	B1
	B2
	B3
	C1
	C2
	C3
	SquareNumber
)

const (
	Empty Token = 0
	Cross = 1
	Nought = 4
)

func (t Token) Flip() Token {
	if t == Cross {
		return Nought
	} else if t == Nought {
		return Cross
	}

	return Empty
}

func (t Token) String() string {
	if t == Cross {
		return "X"
	} else if t == Nought {
		return "O"
	}

	return " "
}

func StringToSquare(str string) (Square, error) {
	var square Square

	str = strings.ToUpper(str)

	converter := map[string]Square {
		"A1": A1, "A2": A2, "A3": A3,
		"B1": B1, "B2": B2, "B3": B3,
		"C1": C1, "C2": C2, "C3": C3,
	}

	square, ok := converter[str]

	if !ok {
		return square, errors.New("Bad string format")
	}

	return square, nil
}

func (board *Board) DoMove(move Move) {
	board[move.Destination] = move.Token
}

func (board *Board) UndoMove(move Move) {
	board[move.Destination] = Empty
}

func (board *Board) GetMoves(token Token) MoveList {
	moveList := make(MoveList, 0, 9)

	for sq := A1; sq < SquareNumber; sq++ {
		if board[sq] != Empty {
			continue
		}

		move := Move{
			Token: token,
			Destination: sq,
		}

		moveList = append(moveList, move)
	}

	return moveList
}

func (board *Board) IsEnd() bool {
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
			return true
		}
	}

	if len(board.GetMoves(Cross)) == 0 {
		return true
	}

	return false
}

func (board Board) String() string {
	var str string

	str += "    1   2   3\n"
	str += "  +---+---+---+\n"

	for i := 0; i < 3; i++ {
		str += fmt.Sprintf(
			"%v | %v | %v | %v |\n",
			[3]string{"A", "B", "C"}[i],
			board[i * 3],
			board[i * 3 + 1],
			board[i * 3 + 2],
		)
		
		str += "  +---+---+---+\n"
	}

	return str
}
