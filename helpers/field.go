package helpers

import (
	"reflect"
)

func IsEmptyField(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func Difference(a, b []int) []int {
	m := make(map[int]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	for _, item := range b {
		if _, ok := m[item]; ok {
			delete(m, item)
		}
	}
	diff := make([]int, 0, len(m))
	for item := range m {
		diff = append(diff, item)
	}
	return diff
}

func Intersection(a, b []int) []int {
	m := make(map[int]struct{})
	for _, item := range a {
		m[item] = struct{}{}
	}
	intersection := make([]int, 0)
	for _, item := range b {
		if _, ok := m[item]; ok {
			intersection = append(intersection, item)
		}
	}
	return intersection
}
