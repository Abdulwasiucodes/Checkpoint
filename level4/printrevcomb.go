package main

import (
	"github.com/01-edu/z01"
)

func main() {
for a:='9';a >='0';a--{
    for b:='9';b >= '0';b--{
        for c:='9';c>= '0';c--{
            if a > b && b > c{
                z01.PrintRune(a)
                z01.PrintRune(b)
                z01.PrintRune(c)
                if a =='2'&& b == '1'&& c =='0'{
                }else{
                    z01.PrintRune(',')
                    z01.PrintRune(' ')
                }
            }
        }
    }
}
z01.PrintRune('\n')
}