package main

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

// This demo containts few code directly from Chipmunk Physics

import (
  "github.com/jteeuwen/glfw"
  "github.com/banthar/gl"
  . "github.com/ftrvxmtrx/gochipmunk/chipmunk"
  "math"
  "math/rand"
  "time"
)

var (
  Running = true
  Width = 640
  Height = 480
  space Space
)

const (
  NotGrabable = Layers(1)
)

func main() {
  rand.Seed(time.Now().UnixNano())

  glfw.Init()
  defer glfw.Terminate()
  glfw.OpenWindow(Width, Height, 8, 8, 8, 8, 0, 8, glfw.Windowed)
  defer glfw.CloseWindow()

  glfw.SetSwapInterval(1)
  glfw.SetWindowTitle("Chipmunk demo")
  glfw.SetWindowSizeCallback(onResize)
  glfw.SetKeyCallback(onKey)

  initGL()
  initScene()

  Running = true

  for Running && glfw.WindowParam(glfw.Opened) == 1 {
    drawScene()
    space.Step(1.0/20.0/3.0)
    space.Step(1.0/20.0/3.0)
    space.Step(1.0/20.0/3.0)
  }
}

func initScene() {
  s := SpaceNew()
  space = s
	s.SetIterations(30)
	s.SetGravity(VectNew(0.0, -100.0))
	s.SetSleepTimeThreshold(0.5)
	s.SetCollisionSlop(0.5)

  static := s.StaticBody()

  shape := s.AddShape(SegmentShapeNew(static, VectNew(-320.0, -240.0), VectNew(-320.0, 240.0), 0.0))
  shape.SetElasticity(1.0)
  shape.SetFriction(1.0)
  shape.SetLayers(NotGrabable)

  shape = s.AddShape(SegmentShapeNew(static, VectNew(320.0, -240.0), VectNew(320.0, 240.0), 0.0))
  shape.SetElasticity(1.0)
  shape.SetFriction(1.0)
  shape.SetLayers(NotGrabable)

  shape = s.AddShape(SegmentShapeNew(static, VectNew(-320.0, -240.0), VectNew(320.0, -240.0), 0.0))
  shape.SetElasticity(1.0)
  shape.SetFriction(1.0)
  shape.SetLayers(NotGrabable)

  for i := 0; i < 16; i++ {
    for j := 0; j < 16; j++ {
      var body Body

      if rand.Float64() < 0.5 {
        w, h := 15.0 + 15.0*rand.Float64(), 15.0 + 15.0*rand.Float64()
        body = s.AddBody(BodyNew(1.0, MomentForBox(1.0, w, h)))
        shape = s.AddShape(BoxShapeNew(body, w, h))
      } else {
        r := 1.0 + 14.0*rand.Float64()
        body = s.AddBody(BodyNew(1.0, MomentForCircle(1.0, r, r, Origin())))
        shape = s.AddShape(CircleShapeNew(body, r, Origin()))
      }

      r1, r2 := 8.0*(0.5 - rand.Float64()), 8.0*(0.5 - rand.Float64())
      body.SetPosition(VectNew(r1 + float64(j)*32.0 - 250.0, r2 + 500.0 - float64(i)*32.0))
      shape.SetElasticity(0.5)
      shape.SetFriction(0.8)
    }
  }
}

func drawScene() {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.MatrixMode(gl.MODELVIEW);
	gl.LoadIdentity()
  gl.Color3f(1, 1, 1)
  drawShapes(space)

  glfw.SwapBuffers()
}

var (
  circleVerts = []float32{
    0.0000, 1.0000,
    0.2588, 0.9659,
    0.5000, 0.8660,
    0.7071, 0.7071,
    0.8660, 0.5000,
    0.9659, 0.2588,
    1.0000, 0.0000,
    0.9659, -0.2588,
    0.8660, -0.5000,
    0.7071, -0.7071,
    0.5000, -0.8660,
    0.2588, -0.9659,
    0.0000, -1.0000,
    -0.2588, -0.9659,
    -0.5000, -0.8660,
    -0.7071, -0.7071,
    -0.8660, -0.5000,
    -0.9659, -0.2588,
    -1.0000, -0.0000,
    -0.9659, 0.2588,
    -0.8660, 0.5000,
    -0.7071, 0.7071,
    -0.5000, 0.8660,
    -0.2588, 0.9659,
    0.0000, 1.0000,
    0.0, 0.0,
  } // For an extra line to see the rotation.
)

func drawCircle(s CircleShape) {
  center := s.Center()
  radius := s.Radius()
  angle := s.Body().Angle()

  gl.VertexPointer(2, 0, circleVerts)
  gl.PushMatrix()
  gl.Translated(center.X, center.Y, 0.0)
  gl.Rotated(angle*180.0/math.Pi, 0.0, 0.0, 1.0)
  gl.Scaled(radius, radius, 1.0)
  gl.DrawArrays(gl.LINE_STRIP, 0, len(circleVerts)/2)
  gl.PopMatrix()
}

func drawSegment(s SegmentShape) {
  verts := []float32{
    float32(s.A().X), float32(s.A().Y),
    float32(s.B().X), float32(s.B().Y),
  }

  gl.VertexPointer(2, 0, verts)
  gl.DrawArrays(gl.LINES, 0, 2)
}

func drawPoly(s PolyShape) {
  verts := s.VertsWorldFloat64()
  gl.VertexPointer(2, 0, verts)
  gl.DrawArrays(gl.LINE_LOOP, 0, len(verts)/2)
}

func drawShapes(s Space) {
  s.EachShape(func(sh Shape) {
    switch sh.(type) {
    case CircleShape:
      drawCircle(sh.(CircleShape))

    case SegmentShape:
      drawSegment(sh.(SegmentShape))

    case PolyShape:
      drawPoly(sh.(PolyShape))
    }
  })
}

func initGL() {
	gl.EnableClientState(gl.VERTEX_ARRAY)
  gl.ShadeModel(gl.SMOOTH)
  gl.ClearColor(0, 0, 0, 0)
  gl.ClearDepth(1)
  gl.DepthFunc(gl.LEQUAL)
  gl.Hint(gl.PERSPECTIVE_CORRECTION_HINT, gl.NICEST)
  gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
}

func onResize(w, h int) {
  if h == 0 {
    h = 1
  }

  gl.Viewport(0, 0, w, h)

	scale := math.Min(float64(w)/640.0, float64(h)/480.0)
	hw := float64(w)*(0.5/scale)
	hh := float64(h)*(0.5/scale)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(-hw, hw, -hh, hh, -1.0, 1.0)
	gl.Translated(0.5, 0.5, 0.0)

  Width = w
  Height = h
}

func onKey(key, state int) {
  switch key {
  case glfw.KeyEsc:
    Running = false
  }
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
