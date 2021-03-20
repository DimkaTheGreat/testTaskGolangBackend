package main

import (
	"strings"
	"sync"
	"testing"
)

func TestTest1(t *testing.T) {
	testSlice := []string{"one", "two", "three", "four"}
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	result := Test1(testSlice, wg, mu)

	//удаление пробела в конце строки для корректного подсчета элементов
	result = strings.TrimSuffix(result, " ")
	resultSlice := strings.Split(result, " ")

	// проверка на совпадение количества элементов тестового среза и результата действия функции
	if len(resultSlice) != len(testSlice) {
		t.Errorf("Expected %d, got %d", len(testSlice), len(resultSlice))

	}

	//проверка на корректность передачи переменной в цикл for
	for i := range testSlice {

		for j := i + 1; j < len(testSlice); j++ {
			if testSlice[i] == testSlice[j] {
				t.Errorf("Matches on index %d and %d\n Full slice : %s", i, j, testSlice)
			}

		}

	}

}
