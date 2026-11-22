func isValidSudoku(board [][]byte) bool {
	type Point struct {
		X int
		Y int
	}
	grid_freq := make(map[Point]map[int]bool)
	row_freq := make(map[int]map[int]bool)
	col_freq := make(map[int]map[int]bool)
	for row_idx, row := range board {
		for col_idx, cell := range row {
			if cell != 46 {
				// Checks for submatrix
				grid_row_idx := int(row_idx / 3)
				grid_col_idx := int(col_idx / 3)
				temp_point := Point{
					X: grid_row_idx,
					Y: grid_col_idx,
				}
				_, grid_ok := grid_freq[temp_point]
				if grid_ok {
					_, cell_ok := grid_freq[temp_point][int(cell)]
					if cell_ok {
						return false
					} else {
						grid_freq[temp_point][int(cell)] = true
					}
				} else {
					grid_freq[Point{
						X: grid_row_idx,
						Y: grid_col_idx,
					}] = map[int]bool{int(cell): true}
				}
				// Checks for row duplication
				_, row_ok := row_freq[row_idx]
				if row_ok {
					_, cell_ok := row_freq[row_idx][int(cell)]
					if cell_ok {
						return false
					} else {
						row_freq[row_idx][int(cell)] = true
					}
				} else {
					row_freq[row_idx] = map[int]bool{int(cell): true}
				}
				// Checks for col duplication
				_, col_ok := col_freq[col_idx]
				if col_ok {
					_, cell_ok := col_freq[col_idx][int(cell)]
					if cell_ok {
						return false
					} else {
						col_freq[col_idx][int(cell)] = true
					}
				} else {
					col_freq[col_idx] = map[int]bool{int(cell): true}
				}
			}
		}
	}
	return true
}