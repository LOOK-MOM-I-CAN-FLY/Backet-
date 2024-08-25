package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk проходит по дереву t и отправляет все значения из дерева в канал ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same определяет, содержат ли деревья t1 и t2 одинаковые значения.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Сравниваем значения из обоих каналов.
	for v1 := range ch1 {
		v2, ok := <-ch2
		if !ok || v1 != v2 {
			return false
		}
	}
	// Проверяем, что второй канал тоже пуст.
	_, ok := <-ch2
	return !ok
}

func main() {
	// Тестируем функцию Same
	fmt.Println(Same(tree.New(1), tree.New(1))) // Должно быть true
	fmt.Println(Same(tree.New(1), tree.New(2))) // Должно быть false
}

