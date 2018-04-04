package main

type move struct {
	x int
	y int
}

type ChessPiece interface {
	validMove(from, to move) bool
	// validMoves(m move) []move
}

type Piece struct {
	coloured bool
	symbol string
}

func (p *Piece) validMove(from, to move) bool {
	if p.coloured {
		return to.y == from.y - 1 || to.y == from.y - 2
	}
	return to.y == from.y + 1 || to.y == from.y + 2
}

var symbols = [...]string{"♙", "♟", "♘", "♞", "♗", "♝", "♖", "♜", "♕", "♛", "♔", "♚"}

func getSymbol(idx int, coloured bool) string {
	if coloured {
		return symbols[idx]
	} else {
		return symbols[idx + 1]
	}
}

func newPawn(coloured bool) (p Piece) {
	return Piece{coloured, getSymbol(0, coloured)}
}

// func newKnight(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(2, coloured)}
// }

// func newBishop(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(4, coloured)}
// }

// func newRook(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(6, coloured)}
// }

// func newQueen(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(8, coloured)}
// }

// func newKing(coloured bool) (p Piece) {
// 	return Piece{coloured, getSymbol(10, coloured)}
// }

