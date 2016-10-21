package main

import (
  "fmt"
  "sort"
)

func pprint(a [][]int) {
  t := len(a)
  for i := 0; i < t; i++ {
    fmt.Println(a[i])
  }
}

func Search(array []int, element int) int {
  index := -1
  for i := 0; i < len(array); i++ {
    if array[i] == element {
      index = i 
    }
  }
  return index
}
func IntMin(n int, m int) int {
  if n < m {
    return n
  } else {
    return m
  }
} 

func Min(n int, m int, z int) int {
  return IntMin(IntMin(n, m), z)
} 

func LevenshteinDistance(s string, t string) int {
  m := len(s) + 1
  n := len(t) + 1
  distance_matrix := make([][]int, m)
  for i := 0; i < m; i++ {
    distance_matrix[i] = make([]int, n)
  }
  for i := 1; i < m; i++ {
    distance_matrix[i][0] = i
  }  
  for j := 1; j < n; j++ {
    distance_matrix[0][j] = j
  }
  for j := 1; j < n; j++ {
    for i := 1; i < m; i++ {
      cost := 0
      if s[i - 1] == t[j - 1] {
        cost = 0
      } else {
        cost = 1
      }
      distance_matrix[i][j] = Min(distance_matrix[i - 1][j] + 1, 
                                  distance_matrix[i][j - 1] + 1, 
                                  distance_matrix[i - 1][j - 1] + cost)
    }
  }
  return distance_matrix[len(s)][len(t)]
}

func SimpleSort(targer string, string_array []string) []string {
  distances := make([]int, len(string_array))
  for i := 0; i < len(string_array); i++ {
    distances[i] = LevenshteinDistance(targer, string_array[i])
  }
  distances_sorted := make([]int, len(string_array))
  copy(distances_sorted, distances)
  sort.Ints(distances_sorted)
  result := make([]string, len(string_array))
  for i := 0; i < len(distances_sorted); i++ {
    index := Search(distances, distances_sorted[i])
    result[i] = string_array[index]
    distances = append(distances[:index], distances[index+1:]...)
    string_array = append(string_array[:index], string_array[index+1:]...)
  }  
  return result
}

func main() {
  fmt.Println(LevenshteinDistance("substitutionCost", "sitting"))
  fmt.Println(LevenshteinDistance("sitting", "substitutionCost"))

  targer := "stat1"
  sample_array := []string{"day", "pay", "kitten", "chick"}
  
  fmt.Println(SimpleSort(targer, sample_array))

}