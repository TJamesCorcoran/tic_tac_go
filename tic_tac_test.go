package main

import (
	"testing"
)


func Test_mark_to_s(t *testing.T) {
	ret := mark_to_s(m_EMPTY)
	if ret != " "  {
		t.Fatalf("expected ' ' got '%q'", ret)
	}
	ret = mark_to_s(m_X)
	if ret != "X"  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}
	ret = mark_to_s(m_O)
	if ret != "O"  {
		t.Fatalf("expected 'O' got '%q'", ret)
	}
}


func Test_eval(t *testing.T) {

	// empty board
	//
	bb := Board{
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
	}
	ret := eval(bb)
	if ret != m_EMPTY  {
		t.Fatalf("expected m_EMPTY got '%q'", ret)
	}

	// incomplete row
	//
	bb = Board{
		{m_X,     m_X,     m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
	}
	ret = eval(bb)
	if ret != m_EMPTY  {
		t.Fatalf("expected m_EMPTY got '%q'", ret)
	}


	// horiz row
	bb = Board{
		{m_X,     m_X,     m_X},
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}
	bb = Board{
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_X,     m_X,     m_X},
		{m_EMPTY, m_EMPTY, m_EMPTY},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}
	bb = Board{
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
		{m_X,     m_X,     m_X},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}

	// vert col
	bb = Board{
		{m_X,     m_EMPTY, m_EMPTY},
		{m_X,     m_EMPTY, m_EMPTY},
		{m_X,     m_EMPTY, m_EMPTY},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}
	bb = Board{
		{m_EMPTY, m_X,     m_EMPTY},
		{m_EMPTY, m_X,     m_EMPTY},
		{m_EMPTY, m_X,     m_EMPTY},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}
	bb = Board{
		{m_EMPTY, m_EMPTY, m_X},
		{m_EMPTY, m_EMPTY, m_X},
		{m_EMPTY, m_EMPTY, m_X},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)
	}

	// diag
	bb = Board{
		{m_X,      m_EMPTY, m_EMPTY},
		{m_EMPTY,  m_X,     m_EMPTY},
		{m_EMPTY,  m_EMPTY, m_X},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)

	}
	bb = Board{
		{m_EMPTY,  m_EMPTY, m_X},
		{m_EMPTY,  m_X,     m_EMPTY},
		{m_X,      m_EMPTY, m_EMPTY},
	}
	ret = eval(bb)
	if ret != m_X  {
		t.Fatalf("expected 'X' got '%q'", ret)

	}

}

func Test_mark_alt(t *testing.T) {
	ret := mark_alt(m_X)
	if ret != m_O  {
		t.Fatalf("expected 'O'")
	}

	ret = mark_alt(m_O)
	if ret != m_X  {
		t.Fatalf("expected 'X'")
	}

}

func Test_play(t *testing.T) {
	bb := Board{
		{m_X,     m_EMPTY, m_EMPTY},
		{m_X,     m_EMPTY, m_EMPTY},
		{m_EMPTY, m_EMPTY, m_EMPTY},
	}
	x, y := play(bb, m_X, 0)
	if x != 0 || y != 2   {
		t.Fatalf("expected x == 0 (was %d), y == 2 (was %d)", x, y )
	}

	
}
