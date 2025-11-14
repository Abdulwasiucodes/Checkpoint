package piscine
func LastWord(s string) string {
  start:= -1
  end:=-1
for i:=len(s)-1;i >=0;i--{
    if s[i] != ' '{
        if end== -1{
            end =i+1
        }
        start = i
    }else if end!=-1{
        break
    }
}
if start == -1{
return "\n"
}
return s[start:end]+"\n"
}