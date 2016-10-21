package main

import ("testing"
        "fmt")

func TestWithCorrecrResult(t *testing.T) {
  expectedResult := []string{"pay", "day", "chick","kitten"}
  result := SimpleSort("pain", []string{"day", "pay", "kitten", "chick"})

  if len(expectedResult) != len(result) {
    t.Errorf("Expected score of ['day', 'pay', 'kitten', 'chick'], but it was %s instead.", result)
  } else {
      for i := 0; i < len(result); i++ {
        if result[i] != expectedResult[i] {
          t.Errorf("Expected score of ['day', 'pay', 'kitten', 'chick'], but it was %s instead.", result)
          break
        }
      }
    }
}

func TestWithEmptyArrayhWords(t *testing.T) {
  expectedResult := []string{}
  result := SimpleSort("pain", []string{})
  fmt.Println(result)

  if len(expectedResult) != len(result) {
    t.Errorf("Expected score of [], but it was %s instead.", result)
  } else {
      for i := 0; i < len(result); i++ {
        if result[i] != expectedResult[i] {
          t.Errorf("Expected score of [], but it was %s instead.", result)
          break
        }
      }
    }
}

func TestWithOneWord(t *testing.T) {
  expectedResult := []string{"cool"}
  result := SimpleSort("pain", []string{"cool"})

  if len(expectedResult) != len(result) {
    t.Errorf("Expected score of [], but it was %s instead.", result)
  } else {
      for i := 0; i < len(result); i++ {
        if result[i] != expectedResult[i] {
          t.Errorf("Expected score of ['cool'], but it was %s instead.", result)
          break
        }
      }
    }
}
