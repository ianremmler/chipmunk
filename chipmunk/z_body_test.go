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

import (
  "testing"
)

func Test_BodyNewFree(t *testing.T) {
  b := BodyNew(1.0, 1.0)

  if b.c() == nil {
    t.Fatal("nil body")
  }

  b.Free()
}

func Test_BodyInSpace(t *testing.T) {
  s := SpaceNew()
  b := BodyNew(1.0, 1.0)

  if b.Space() != nil || s.Contains(b) {
    t.Fatal("not nil space before Add")
  }

  s.AddBody(b)

  if b.Space() == nil || !s.Contains(b) {
    t.Fatal("nil space after Add")
  }

  s.RemoveBody(b)

  if b.Space() != nil || s.Contains(b) {
    t.Fatal("not nil space after Remove")
  }

  b.Free()
  s.Free()
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
