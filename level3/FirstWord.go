package piscine
func FirstWord(s string) string {
  start:= 0
  for start < len(s) && s[start]==' '{
    start++
  }
  end := start
  for end < len(s) && s[end] != ' '{
    end++
  }
  return s[start:end] +"\n"
}