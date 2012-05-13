// Package chipmunk is an interface to Chipmunk Physics library
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

/*
#cgo CFLAGS: -I/usr/include/chipmunk
#cgo LDFLAGS: -lchipmunk -lm

#include <chipmunk.h>
*/
import "C"

import (
  "unsafe"
)

type CollisionType C.cpCollisionType

type Group C.cpGroup

type Layers uint

func (c CollisionType) c() C.cpCollisionType {
  return C.cpCollisionType(c)
}

func (g Group) c() C.cpGroup {
  return C.cpGroup(g)
}

func (l Layers) c() C.cpLayers {
  return C.cpLayers(l)
}

func boolToC(b bool) C.cpBool {
  v := 0

  if b {
    v = 1
  }

  return C.cpBool(v)
}

func cpBool(b C.cpBool) bool {
  return int(b) != 0
}

func MomentForCircle(m, r1, r2 float64, offset Vect) float64 {
  return float64(C.cpMomentForCircle(C.cpFloat(m), C.cpFloat(r1), C.cpFloat(r2), offset.c()))
}

func AreaForCircle(r1, r2 float64) float64 {
  return float64(C.cpAreaForCircle(C.cpFloat(r1), C.cpFloat(r2)))
}

func MomentForSegment(m float64, a, b Vect) float64 {
  return float64(C.cpMomentForSegment(C.cpFloat(m), a.c(), b.c()))
}

func AreaForSegment(a, b Vect, r float64) float64 {
  return float64(C.cpAreaForSegment(a.c(), b.c(), C.cpFloat(r)))
}

func MomentForPoly(m float64, verts []Vect, offset Vect) float64 {
  v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
  return float64(C.cpMomentForPoly(C.cpFloat(m), C.int(len(verts)), v, offset.c()))
}

func AreaForPoly(verts []Vect) float64 {
  v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
  return float64(C.cpAreaForPoly(C.int(len(verts)), v))
}

func CentroidForPoly(verts []Vect) Vect {
  v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
  return cpVect(C.cpCentroidForPoly(C.int(len(verts)), v))
}

func RecenterPoly(verts []Vect) {
  v := (*C.cpVect)(unsafe.Pointer(&verts[0]))
  C.cpRecenterPoly(C.int(len(verts)), v)
}

func MomentForBox(m, width, height float64) float64 {
  return float64(C.cpMomentForBox(C.cpFloat(m), C.cpFloat(width), C.cpFloat(height)))
}

func MomentForBox2(m float64, box BB) float64 {
  return float64(C.cpMomentForBox2(C.cpFloat(m), box.c()))
}
