package chipmunk

/*
Copyright (c) 2012 Serge Zirukin

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// #include <chipmunk.h>
import "C"

import (
  "fmt"
  "math"
)

// Vect is a 2D vector type.
type Vect struct {
  X, Y float64
}

// VectNew returns a new 2D vector.
func VectNew(x, y float64) Vect {
  return Vect{x, y}
}

// VectForAngle returns the unit length vector for the given angle (radians).
func VectForAngle(a float64) Vect {
  return VectNew(math.Cos(a), math.Sin(a))
}

// Origin returns zero vector.
func Origin() Vect {
  return VectNew(0.0, 0.0)
}

func (v Vect) String() string {
  return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func (a Vect) Add(b Vect) Vect {
  return VectNew(a.X+b.X, a.Y+b.Y)
}

func (a Vect) Sub(b Vect) Vect {
  return VectNew(a.X-b.X, a.Y-b.Y)
}

func (v Vect) Mul(x float64) Vect {
  return VectNew(v.X*x, v.Y*x)
}

func (v Vect) Div(x float64) Vect {
  return VectNew(v.X/x, v.Y/x)
}

// Dot returns a dot product of two vectors.
func (v1 Vect) Dot(v2 Vect) float64 {
  return v1.X*v2.X + v1.Y*v2.Y
}

// Length returns the length of vector.
func (v Vect) Length() float64 {
  return math.Sqrt(v.Dot(v))
}

func (v Vect) Neg() Vect {
  return VectNew(-v.X, -v.Y)
}

// Lerp does a spherical linear interpolation between two vectors.
func (v1 Vect) Lerp(v2 Vect, t float64) Vect {
  omega := math.Acos(v1.Dot(v2))

  if omega > 0.0 {
    denom := 1.0 / math.Sin(omega)
    return v1.Mul(math.Sin((1.0-t)*omega) * denom).Add(v2.Mul(math.Sin(t*omega) * denom))
  }

  return v1
}

// LerpConst does a spherical linear interpolation between two vectors
// by no more than specific angle (radians).
func (v1 Vect) LerpConst(v2 Vect, a float64) Vect {
  angle := math.Acos(v1.Dot(v2))
  return v1.Lerp(v2, math.Min(a, angle)/angle)
}

// ToAngle returns the angular direction vector is pointing in (radians).
func (v Vect) ToAngle() float64 {
  return math.Atan2(v.Y, v.X)
}

func (v Vect) c() C.cpVect {
  return C.cpVect{x: C.cpFloat(v.X), y: C.cpFloat(v.Y)}
}

func cpVect(v C.cpVect) Vect {
  return VectNew(float64(v.x), float64(v.y))
}
