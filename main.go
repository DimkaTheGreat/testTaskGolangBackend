package main

import (
	"fmt"

	"sync"
)

// Test1 принимает на вход список строк и выводит их в одной строке через пробел, возможно, в случайном порядке.
// Смысл этого кода не в конкатенации слов через пробел, а в понимании работы многопоточности.
//
// Вам необходимо исправить код так, чтобы он работал, не удаляя ничего из него, а только внося исправления.
//
// Доп. задание: напишите тест на эту функцию.
//
// Пример вызова функции:
//     data1 := []string{"hello", "world", "test", "data", "code"}
//     r1 := Test1(data1)
//     fmt.Println(r1)

// func Test1(list []string) string {
// 	var result string

// 	for _, l := range list {
// 		go func() {
// 			result += l + " "
// 		}()
// 	}

// 	return result
// }

//исправленная функция Test1
func Test1(list []string, wg *sync.WaitGroup, mu *sync.Mutex) string {
	var result string

	for _, l := range list {
		go func(l string) {
			mu.Lock()
			result += l + " "
			mu.Unlock()
			wg.Done()
		}(l)

	}
	wg.Wait()

	return result
}

func main() {
	//Test1
	data1 := []string{"hello", "world", "test", "data", "code"}
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(len(data1))
	r1 := Test1(data1, wg, mu)
	fmt.Printf("Test №1 output : %s\n", r1)

}
