package generics

import "reflect"

func FirstOrDefault[T any](slice []T, filter func(*T) bool) (element *T) {
	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			return &slice[i]
		}
	}

	return nil
}

func Where[T any](slice []T, filter func(*T) bool) []*T {
	var ret []*T = make([]*T, 0)

	for i := 0; i < len(slice); i++ {
		if filter(&slice[i]) {
			ret = append(ret, &slice[i])
		}
	}

	return ret
}

func Any(slice interface{}, f func(value interface{}) bool) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return true
			}
		}
	}

	return false
}

func Find(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}

	return -1
}

func FindAndDelete[T any](s []T, filter func(*T) bool) []T {
	var new []T = make([]T, 0)
	index := 0
	for j, i := range s {
		if !filter(&s[j]) {
			new = append(new, i)
			index++
		}
	}

	return new[:index]
}

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
