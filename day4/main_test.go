package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXmas(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := xmas(input)

	assert.Equal(t, 9, result)
}

func TestSearch(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

	result := search(views(input))

	assert.Equal(t, 18, result)
}

func TestViews(t *testing.T) {
	// M M M S X M
	// M S A M X M
	// A M X S X M
	// M S A M A S
	// X M A S A M

	input :=
		`MMMSXM
MSAMXM
AMXSXM
MSAMAS
XMASAM`

	vertical :=
		`MMAMX
MSMSM
MAXAA
SMSMS
XXXAA
MMMSM
`

	diagonalLtr :=
		`MSXMA
MASAM
MMXS
SXM
XM
M
MMAS
ASA
MM
X
`

	diagonalRtl :=
		`MXSAM
XMXSX
SAMM
MSA
MM
M
MXMA
MAS
SA
M
`

	views := views(input)

	assert.Equal(t, input, views[0])
	assert.Equal(t, vertical, views[1])
	assert.Equal(t, diagonalLtr, views[2])
	assert.Equal(t, diagonalRtl, views[3])
}
