package week_6

/*
36. Valid Sudoku
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.

Example 1:

Input: board =
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]

Output: true

Example 2:

Input: board =
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]

Output: false

Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Constraints:

board.length == 9
board[i].length == 9
board[i][j] is a digit 1-9 or '.'.
*/

func isValidSudoku(board [][]byte) bool {
	byteMap := make(map[byte]bool, 9)
	for index := range subBoxCenters {
		i, j := subBoxCenters[index][0], subBoxCenters[index][1]
		for n := range neighbours {
			i += neighbours[n][0]
			j += neighbours[n][1]
			if byteMap[board[i][j]] && board[i][j] != '.' {
				return false
			}
			byteMap[board[i][j]] = true
		}
		byteMap = map[byte]bool{}
	}

	//horizontal checking
	for i := range board {
		for j := 0; j < 9; j++ {
			if byteMap[board[i][j]] && board[i][j] != '.' {
				return false
			}
			byteMap[board[i][j]] = true
		}
		byteMap = map[byte]bool{}
	}

	//vertical checking
	for j := range board[0] {
		for i := 0; i < 9; i++ {
			if byteMap[board[i][j]] && board[i][j] != '.' {
				return false
			}
			byteMap[board[i][j]] = true
		}
		byteMap = map[byte]bool{}
	}

	return true
}

var (
	subBoxCenters = [][]int{
		{1, 1}, {1, 4}, {1, 7},
		{4, 1}, {4, 4}, {4, 7},
		{7, 1}, {7, 4}, {7, 7},
	}

	neighbours = [][]int{
		{0, 0}, {-1, 0}, {0, 1},
		{1, 0}, {1, 0}, {0, -1},
		{0, -1}, {-1, 0}, {-1, 0},
	}
)
