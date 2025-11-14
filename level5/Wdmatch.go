package main
import(
    "os"
    "github.com/01-edu/z01"
)
func main(){
    if len(os.Args) !=3{
        return
    }
    s1:=os.Args[1]
    s2:=os.Args[2]
    i:=0
    for _,char :=range s2{
        if i<len(s1)&& s1[i] ==byte(char){
            i++
        }
    }
        if i ==len(s1){
            for _,char := range s1{
                z01.PrintRune(char)
            }
  z01.PrintRune('\n')
        }
    }