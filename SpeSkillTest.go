package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Narcissistic(angka int) bool {
	string_angka := strconv.Itoa(angka)
	array_angka := strings.Split(string_angka, "")
	int_array_angka := make([]int, len(array_angka))
	for i := range array_angka {
		int_array_angka[i], _ = strconv.Atoi(array_angka[i])
	}

	var hasil int = 0
	for i := 0 ; i < len(int_array_angka) ; i++ {
		hasil = int(hasil + int(math.Pow(float64(int_array_angka[i]), float64(len(int_array_angka)))))
	}
	if hasil == angka {
		return true
	}
	return false
}

func ParityOutlier(angka []int) string {
	var ganjil, genap, indexganjil, indexgenap int
	for i := 0; i < len(angka); i++ {
		if angka[i]%2 == 0 {
			genap++
			indexgenap = i 
		} else {
			ganjil++
			indexganjil = i
		}
	}
	if genap == 0 && ganjil == 0 {
		return "false"
	} else if genap == 1 {
		return fmt.Sprint(angka[indexgenap])
	} else if ganjil == 1 {
		return fmt.Sprint(angka[indexganjil])
	}
	return "false"
}

func FindNeedle(array_kata []string, kata string) int {
	for i, v := range array_kata {
		if v == kata {
			return i
		}
	}
	return 0
}

func BlueOcean(array_int []int, penghapus []int) []int{
	var hasil []int
	for _, v := range array_int {
		if v != penghapus[0] {
			hasil = append(hasil, v)
		}
	}
	return hasil
}