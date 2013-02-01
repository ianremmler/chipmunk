package chipmunk

/*
Copyright © 2012 Serge Zirukin

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

// #include <chipmunk/chipmunk.h>
import "C"

import (
  "fmt"
  "math"
)

////////////////////////////////////////////////////////////////////////////////

// BB is an axis-aligned 2D bounding box type (left, bottom, right, top).
type BB struct {
  l, b, r, t float64
}

////////////////////////////////////////////////////////////////////////////////

// Area returns the area of the bounding box.
func (b BB) Area() float64 {
  return (b.r - b.l) * (b.t - b.b)
}

// BBNew creates a 2D bounding box.
func BBNew(l, b, r, t float64) BB {
  return BB{l, b, r, t}
}

// BBNewForCircle constructs a BB for a circle with the given position and radius.
func BBNewForCircle(p Vect, r float64) BB {
  return BBNew(p.X-r, p.Y-r, p.X+r, p.Y+r)
}

// ClampVect clamps a vector to a bounding box.
func (bb BB) ClampVect(v Vect) Vect {
  return Vect{X: math.Min(math.Max(bb.l, v.X), bb.r), Y: math.Min(math.Max(bb.b, v.Y), bb.t)}
}

// Contains returns true if other bounding box lies completely within.
func (b BB) Contains(other BB) bool {
  return b.l <= other.l && b.r >= other.r && b.b <= other.b && b.t >= other.t
}

// ContainsVect returns true if the bounding box contains a vector.
func (b BB) ContainsVect(v Vect) bool {
  return b.l <= v.X && b.r >= v.X && b.b <= v.Y && b.t >= v.Y
}

// Expand returns a bounding box that holds both bounding box and a vector.
func (b BB) Expand(v Vect) BB {
  return BB{
    math.Min(b.l, v.X),
    math.Min(b.b, v.Y),
    math.Max(b.r, v.X),
    math.Max(b.t, v.Y)}
}

// Intersects returns true if two bounding boxes intersect.
func (a BB) Intersects(b BB) bool {
  return a.l <= b.r && b.l <= a.r && a.b <= b.t && b.b <= a.t
}

// IntersectsSegment returns true if the bounding box intersects the line
// segment defined using two points.
func (bb BB) IntersectsSegment(a, b Vect) bool {
  return !math.IsInf(bb.SegmentQuery(a, b), 1)
}

// Merge returns a bounding box that holds both bounding boxes.
func (a BB) Merge(b BB) BB {
  return BB{
    math.Min(a.l, b.l),
    math.Min(a.b, b.b),
    math.Max(a.r, b.r),
    math.Max(a.t, b.t)}
}

// MergedArea merges two bounding boxes and returns the area of the merged bounding box.
func (a BB) MergedArea(b BB) float64 {
  return (math.Max(a.r, b.r) - math.Min(a.l, b.l)) * (math.Max(a.t, b.t) - math.Min(a.b, b.b))
}

// SegmentQuery returns the fraction along the segment query the BB is hit.
// Returns math.Inf(1) if it doesn't hit.
func (bb BB) SegmentQuery(a, b Vect) float64 {
  pinf := math.Inf(1)
  ninf := math.Inf(-1)

  idx := 1.0 / (b.X - a.X)
  idy := 1.0 / (b.Y - a.Y)
  tx1 := ninf
  tx2 := pinf
  ty1 := ninf
  ty2 := pinf

  if bb.l != a.X {
    tx1 = (bb.l - a.X) * idx
  }

  if bb.r != a.X {
    tx2 = (bb.r - a.X) * idx
  }

  if bb.b != a.Y {
    ty1 = (bb.b - a.Y) * idy
  }

  if bb.t != a.Y {
    ty2 = (bb.t - a.Y) * idy
  }

  txmin := math.Min(tx1, tx2)
  txmax := math.Max(tx1, tx2)
  tymin := math.Min(ty1, ty2)
  tymax := math.Max(ty1, ty2)

  if tymin <= txmax && txmin <= tymax {
    min := math.Max(txmin, tymin)
    max := math.Min(txmax, tymax)

    if 0.0 <= max && min <= 1.0 {
      return math.Max(min, 0.0)
    }
  }

  return pinf
}

// String converts a BB to a human-readable string.
func (b BB) String() string {
  return fmt.Sprintf("(BB){l:%g, b:%g, r:%g, t:%g}", b.l, b.b, b.r, b.t)
}

// WrapVect wraps a vector to a bounding box.
func (bb BB) WrapVect(v Vect) Vect {
  ix := math.Abs(bb.r - bb.l)
  modx := math.Mod(v.X-bb.l, ix)
  x := modx

  if modx <= 0.0 {
    x += ix
  }

  iy := math.Abs(bb.t - bb.b)
  mody := math.Mod(v.Y-bb.b, iy)
  y := mody

  if mody <= 0.0 {
    y += iy
  }

  return Vect{X: x + bb.l, Y: y + bb.b}
}

// c converts BB to C.cpBB.
func (b BB) c() C.cpBB {
  return C.cpBB{
    l: C.cpFloat(b.l),
    b: C.cpFloat(b.b),
    r: C.cpFloat(b.r),
    t: C.cpFloat(b.t)}
}

// cpBB converts C.cpBB to BB.
func cpBB(bb C.cpBB) BB {
  return BBNew(float64(bb.l), float64(bb.b), float64(bb.r), float64(bb.t))
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
