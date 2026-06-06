// Given an integer numRows, return the first numRows of Pascal's triangle.

package pascalstriangle

func GeneratePascalsTriangle(n int) [][]int {
	pt := make([][]int, n)
	pt[0] = make([]int, 1)
	pt[0][0] = 1

	for i := 1; i < n; i++ {
		pt[i] = make([]int, i+1)
		pt[i][0] = 1
		pt[i][i] = 1

		for j := 1; j < i; j++ {
			pt[i][j] = pt[i-1][j-1] + pt[i-1][j]
		}
	}
	return pt
}
