package utils

import (
	"errors"
	route "ova_route_api/internal/models"
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

// SplitToBulks - батчевое разделение слайса на слайс слайсов
func SplitToBulks(entities []route.Route, batchSize uint) ([][]route.Route, error) {
	var (
		resp [][]route.Route
		i    uint
	)

	if batchSize == 0 {
		return nil, errors.New("batchSize should be more than zero")
	}

	tmp := make([]route.Route, len(entities))
	copy(tmp, entities)

	arrLn := uint(len(tmp))
	for i = 0; i < arrLn; i += batchSize {
		if i+batchSize > arrLn {
			resp = append(resp, tmp[i:arrLn])
			break
		}

		resp = append(resp, tmp[i:i+batchSize])
	}

	return resp, nil
}

// ConvertToView - конвертации слайса от структуры в отображение, где ключ
// идентификатор структуры, а значение сама структура
func ConvertToView(entities []route.Route) (map[uint64]route.Route, error) {
	if len(entities) == 0 {
		return nil, errors.New("empty routes list")
	}

	resp := make(map[uint64]route.Route, len(entities))

	for k, v := range entities {
		resp[uint64(k)] = v
	}

	return resp, nil
}
