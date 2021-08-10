package utils

import (
	"errors"
)

// SplitSlice - разделение на слайса на батчи - исходный слайс конвертировать в слайс слайсов -
// чанки одинкового размера (кроме последнего)
func SplitSlice(array []string, batchSize uint) ([][]string, error) {
	var (
		resp [][]string
		i    uint
	)

	if batchSize == 0 {
		return nil, errors.New("batchSize should be more than zero")
	}

	arrLn := uint(len(array))
	for i = 0; i < arrLn; i += batchSize {
		if i+batchSize > arrLn {
			resp = append(resp, array[i:arrLn])
			break
		}

		resp = append(resp, array[i:i+batchSize])
	}

	return resp, nil
}

// ReverseKey - происходит конвертация отображения (“ключ-значение“) в отображение (“значение-ключ“)
func ReverseKey(m map[int]string) (map[string]int, error) {

	resp := make(map[string]int)
	for k, v := range m {
		if _, ok := resp[v]; ok {

			return nil, errors.New("key is duplicated")
		}

		resp[v] = k
	}

	return resp, nil
}

var Ints = []int{1, 3, 5, 7, 13} // Захардкоженный слайс со знаениями

// FilterSlice - фильтрация по списку - нужно отфильтровать входной слайс по критерию отсутствия элемента в списке
func FilterSlice(slice []int, filter []int) ([]int, error) {
	hash := make(map[int]interface{})
	res := make([]int, 0)

	for _, v := range filter {
		hash[v] = nil
	}

	for v := range slice {
		if _, ok := hash[v]; ok {
			res = append(res, v)
		}
	}

	return res, nil
}
