package datamanager

// only used for visualizing 2D Vectors
type Coordinate struct{ X1, X2 float64 }

// N represents the number of columns (so that a vector can be in one array)
// M represents the number rows (e.g. dimension of a vector)
type Matrix struct {
	N, M   int
	Matrix [][]float64
}

// constructor for Matrix
// makes it easier to create the Matrix array
func NewMatrix(n int, m int) *Matrix {
	if n < 0 || m < 0 {
		return nil
	}

	matrix := Matrix{n, m, make([][]float64, n)}
	for i := 0; i < matrix.N; i++ {
		matrix.Matrix[i] = make([]float64, matrix.M)
	}
	return &matrix
}

type Eigen struct {
	Values  []float64
	Vectors Matrix
}
