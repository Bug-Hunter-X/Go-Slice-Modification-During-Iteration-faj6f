func myFunc(a []int) []int {
  var b []int
  for _, v := range a {
    if v != 0 {
      b = append(b, v)
    }
  }
  return b
} 
