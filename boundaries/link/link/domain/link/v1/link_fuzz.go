//go:build gofuzz

package v1

func Fuzz(link []byte) int {
	data := &Link{
		Url: string(link),
	}

	err := NewURL(data)
	if err != nil || len(data.Hash) != 7 {
		return -1
	}

	return 1
}
