func myFunc(a []int) {
  for i := 0; i < len(a); i++ {
    if a[i] == 0 {
      a = append(a[:i], a[i+1:]...) // Removing element modifies the slice
    }
  }
}