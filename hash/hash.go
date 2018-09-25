package hash

type HashService interface {
	Calculate(str string) (string, error)
}

type Hash struct{}

func (h *Hash) Calculate(str string) (string, error) {
	return "", nil
}
