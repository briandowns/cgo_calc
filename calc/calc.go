package calc

/*
static double
addition(const double x, const double y)
{
	return x + y;
}

static double
subtract(const double x, const double y)
{
	return x - y;
}

static double
multiply(const double x, const double y)
{
	return x * y;
}

static double
divide(const double x, const double y)
{
	return x / y;
}
*/
import "C"

// Think of this as Go's standard library. These
// are the crypto functions that are called from
// the primary Go code.

var z C.double

func Add(x, y float64) float64 {
	z = C.addition(C.double(x), C.double(y))
	return float64(z)
}

func Sub(x, y float64) float64 {
	z = C.subtract(C.double(x), C.double(y))
	return float64(z)
}

func Mul(x, y float64) float64 {
	z = C.multiply(C.double(x), C.double(y))
	return float64(z)
}

func Div(x, y float64) float64 {
	z = C.divide(C.double(x), C.double(y))
	return float64(z)
}
