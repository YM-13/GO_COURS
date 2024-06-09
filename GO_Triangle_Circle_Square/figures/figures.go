package figures


var RK float64 = 0.5
var PI float64 = 3.14159

type Figures struct {
	R float64
	A float64
	B float64
	C float64
}

func (f Figures) Circumiference() float64 {
	return 2 * PI * f.R
}

func (f Figures) AreaOfCircle() float64 {
	return PI * f.R * f.R
}


func (f Figures) PSquare() float64 {
	return 4 * f.A
}

func (f Figures) SSquare() float64 {
	return f.A * f.A
}


func (f Figures) PTriangle() float64 {
	return f.A + f.B + f.C
}

func (f Figures) STriangle() float64 {
	return RK * f.A * f.B
}