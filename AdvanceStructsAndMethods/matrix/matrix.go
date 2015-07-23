package matrix

// Visibility is private which enforces factory method.
type matrix struct{
    a int
}

func NewMatrix(a int) *matrix {
    m := new(matrix)
    m.a = a
    return m
}
