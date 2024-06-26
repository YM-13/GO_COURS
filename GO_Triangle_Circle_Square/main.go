package main
import (
	"fmt"
	f "GO_Triangle_Circle_Square/figures"
)

func main() {
	cf := f.Figures{R: 3.0}
	fmt.Println(cf.Circumiference)

	ca := f.Figures{R: 3.0}
	fmt.Println(ca.AreaOfCircle)

	sp := f.Figures{A: 6.0}
	fmt.Appendln(sp.PSquare)

	ss := f.Figures{A: 7.0}
	fmt.Println(ss.SSquare)



	pt := f.Figures{A: 1.0, B: 3.0, C: 9.0}
	fmt.Println(pt.PTriangle)

	st := f.Figures{A: 11.0, B: 4.0}
	fmt.Println(st.STriangle)
}
