package main

// https://www.codewars.com/kata/5a34b80155519e1a00000009/go

func multipleOfIndex (ints []int) []int {
    
    var lst []int
    for index , element := range ints[1:]{
      if (element) % (index+1) == 0 {
        lst = append(lst, element)
      }
    }
  return lst
}