package main

import "fmt"

func main() {
	// Test 1 - List
	//list := NewList()
	//
	//list.PushBack("A")
	//list.PushBack("B")
	//list.PushBack("C")
	//list.PushBack("D")
	//list.PushBack("E")
	//list.PushBack("F")
	//list.PushBack("G")
	//
	//for item := list.Front(); item != nil; item = item.Next {
	//	if item.Value == "D" {
	//		list.Remove(item)
	//	}
	//	if item.Value == "G" {
	//		list.MoveToFront(item)
	//	}
	//	fmt.Println(item.Value)
	//}
	//
	//fmt.Printf("Длина списка: %d\n", list.Len())

	// Test 2 - Cache
	cache := NewCache(3)

	cache.Set("A", 1)
	cache.Set("B", 2)
	cache.Set("C", 3)

	fmt.Printf("1. Объявлем размер кэша - 3 и проставляем A,B,C\n")
	for item := cache.(*lruCache).queue.Front(); item != nil; item = item.Next {
		key := cache.(*lruCache).keys[item]
		fmt.Printf("Key: %s, Value: %v\n", key, item.Value)
	}

	fmt.Printf("\n2. Вставляем новое значение - D\n")
	cache.Set("D", 4)
	for item := cache.(*lruCache).queue.Front(); item != nil; item = item.Next {
		key := cache.(*lruCache).keys[item]
		fmt.Printf("Key: %s, Value: %v\n", key, item.Value)
	}

	fmt.Printf("\n3. Ищем значение A\n")
	if val, ok := cache.Get("A"); ok {
		fmt.Println("A найден:", val)
	} else {
		fmt.Println("A не найден")
	}

	fmt.Printf("\n4. Ищем значение B\n")
	if val, ok := cache.Get("B"); ok {
		fmt.Println("B найден:", val)
	}

	fmt.Printf("\n5. Вставляем новое значение - C\n")
	cache.Set("C", 100)
	for item := cache.(*lruCache).queue.Front(); item != nil; item = item.Next {
		key := cache.(*lruCache).keys[item]
		fmt.Printf("Key: %s, Value: %v\n", key, item.Value)
	}

	fmt.Printf("\n6. Ищем значение C\n")
	if val, ok := cache.Get("C"); ok {
		fmt.Println("C найден: ", val)
	}

	fmt.Println("\n7. Очищаем кэш")
	cache.Clear()

	for item := cache.(*lruCache).queue.Front(); item != nil; item = item.Next {
		key := cache.(*lruCache).keys[item]
		fmt.Printf("Key: %s, Value: %v\n", key, item.Value)
	}
}
