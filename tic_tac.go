package main

import (
	"fmt"
	//	"log"
	"os"
)


type BoardMark uint8

const (
	m_EMPTY BoardMark = 0
	m_X = 1
	m_O = 2
) 

type Board [3][3]BoardMark

// eval a board and return the winner (EMPTY = no winner)

func mark_to_s(mm BoardMark) string{
	var s = []string{ " ", "X", "O" }[mm]
	return s
}


func eval(bb Board) BoardMark {
	var mark BoardMark = m_EMPTY
	for mark = m_X; mark <= m_O ; mark +=1 {
		// fmt.Println("eval mark = " +  mark_to_s(mark))
		// check rows and cols
		for ii := 0; ii <= 2 ; ii += 1 {
			if (bb[ii][0] == mark && bb[ii][1] == mark && bb[ii][2] == mark) ||
				(bb[0][ii] == mark && bb[1][ii] == mark && bb[2][ii] == mark) {
				return mark
			}
		}
		// check diagonals
		if (bb[0][0] == mark && bb[1][1] == mark && bb[2][2] == mark) ||
			(bb[2][0] == mark && bb[1][1] == mark && bb[0][2] == mark) {
				return mark
		}
	}
	return m_EMPTY
}


func mark_alt(mm BoardMark) BoardMark {
	switch mm {
	case m_X:
		return m_O
	case m_O:
		return m_X
	default:
		fmt.Println("error in mark_alt")
		os.Exit(-1)
	}
	return m_EMPTY
}

func print_board(bb Board) {
	for ii := 0; ii <= 2 ; ii += 1 {
		fmt.Println(" "  + mark_to_s(bb[ii][0]) + " | " + mark_to_s(bb[ii][1]) + " | " + mark_to_s(bb[ii][2]) )
		if ii !=2 {
			fmt.Println("---+---+---")
		}
	}
}

// returns an X, Y pair that guarantees that 'O' player can win, no matter what X does
// return 0,0 if no such move exits
//
func play(bb Board, next_move BoardMark, depth int) (int, int) {
	possible_moves := 0
	for ii := 0; ii <= 2 ; ii += 1 {
		for jj := 0; jj <= 2 ; jj += 1 {
			if bb[ii][jj] == m_EMPTY {
				possible_moves += 1
				candidate := Board(bb)
				candidate[ii][jj] = next_move
				fmt.Printf("candidate move: x = %d, y =%d" , ii, jj)
				// print_board(candidate)
				winner := eval(candidate)
				fmt.Println("winner = " + mark_to_s(winner))

				// recurse
				if winner == m_EMPTY {
					play(candidate, mark_alt(next_move), depth + 1)
				} else {
					return ii, jj
				}
				fmt.Println()
				fmt.Println()
			}
		}
	}
	if possible_moves == 0 {
		fmt.Println("exiting")
		os.Exit(0)
	}
	return 0,0
}

func main() {

	fmt.Println("Would you like to play a game?")

	
	bb := Board{
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
	}
	
	play(bb, m_O, 0)
	
	// fmt.Println("current winner of this board is " + mark_to_s(eval(bb)))
}
