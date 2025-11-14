package piscine
func CountAlpha(s string) int {
    count:=0
  for _,char :=range s{
    if ('a' <= char && char <= 'z' )||('A' <= char && char <='Z'){
        count++
    }
  }
  return count
  }