package screen

import (
	"fmt"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/vps2/screenmon/internal/shader"
)

type Area struct {
	X1 int //left top x
	Y1 int //left top y
	X2 int //right bottom x
	Y2 int //right bottom y
}

type AreaSelector struct {
	display int
}

func NewAreaSelector(displayNumber int) *AreaSelector {
	return &AreaSelector{
		display: displayNumber,
	}
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

var (
	verticies = []float32{
		//x,y,z
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	}
	indices = []uint32{
		0, 1, 3, //Треугольник 1
		1, 2, 3, //Треугольник 2
	}
)

var (
	btnPressed                 bool
	selX1, selY1, selX2, selY2 int //для сохранения экранных координат выделенной области
)

var shadersStore = shader.NewStore()

func (as *AreaSelector) Select() Area {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %w", err))
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Floating, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)

	monitor := glfw.GetMonitors()[as.display]
	mode := monitor.GetVideoMode()

	windowWidth := mode.Width
	windowHeight := mode.Height
	//FIXME если размер экрана выставить равным его разрешению, то экран заполняется непрозрачным цветом!
	window, err := glfw.CreateWindow(windowWidth+1, windowHeight+1, "", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %w", err))
	}
	//
	window.SetMouseButtonCallback(mouseButtonCallback)
	window.SetKeyCallback(keyCallback(windowWidth, windowHeight))
	window.SetCursorPosCallback(cursorPosCallback(windowWidth, windowHeight))
	//
	monPosX, monPosY := monitor.GetPos()
	window.SetPos(monPosX, monPosY)
	//
	cursor := glfw.CreateStandardCursor(glfw.CrosshairCursor)
	defer cursor.Destroy()
	window.SetCursor(cursor)
	//
	window.SetOpacity(0.3)
	//
	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(fmt.Errorf("could not initialize gl: %w", err))
	}

	gl.Viewport(0, 0, int32(windowWidth), int32(windowHeight))

	// Configure the vertex and fragment shaders
	program, err := shader.NewProgram(shadersStore.GetShader(shader.Vertex),
		shadersStore.GetShader(shader.Fragment))
	if err != nil {
		panic(fmt.Errorf("could not configure program: %w", err))
	}

	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(verticies)*4, nil, gl.DYNAMIC_DRAW)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	gl.ClearColor(0.0, 0.5, 1, 1.0)

	program.Use()

	_ = program.SetUniform1("width", float32(windowWidth))
	_ = program.SetUniform1("height", float32(windowHeight))
	_ = program.SetUniform1("border", float32(1))

	for !window.ShouldClose() {
		glfw.PollEvents()

		gl.Clear(gl.COLOR_BUFFER_BIT)

		if btnPressed {
			_ = program.SetUniform1("x1", float32(selX1))
			_ = program.SetUniform1("y1", float32(selY1))
			_ = program.SetUniform1("x2", float32(selX2))
			_ = program.SetUniform1("y2", float32(selY2))

			// TODO попробовать переделать на использование матрицы трансформации, вместо переписывания координат в буфере.
			gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
			gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(verticies)*4, gl.Ptr(verticies))
			gl.BindVertexArray(vao)

			//Отрисовываем объект
			gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))

			gl.BindBuffer(gl.ARRAY_BUFFER, 0)
			gl.BindVertexArray(0)
		}

		window.SwapBuffers()

		time.Sleep(29 * time.Millisecond) //ограничение fps до 35 (1000/fps)  -> ожидание миллисекунд
	}

	gl.DeleteBuffers(1, &ebo)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)

	area := Area{
		X1: selX1,
		Y1: selY1,
		X2: selX2,
		Y2: selY2,
	}

	return area
}

func keyCallback(width, height int) glfw.KeyCallback {
	return func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if key == glfw.KeyEscape && action == glfw.Press {
			selX1 = 0
			selY1 = 0
			selX2 = width
			selY2 = height

			w.SetShouldClose(true)
		}
	}
}

func mouseButtonCallback(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	cursorX, cursorY := w.GetCursorPos()

	if button == glfw.MouseButton1 {
		if action == glfw.Press {
			btnPressed = true

			selX1 = int(cursorX)
			selY1 = int(cursorY)
		} else if action == glfw.Release {
			btnPressed = false

			if int(cursorX) != selX1 && int(cursorY) != selY1 {
				w.SetShouldClose(true)
			} else {
				return
			}

			verticies[0] = 0
			verticies[1] = 0
			//
			verticies[3] = 0
			verticies[4] = 0
			//
			verticies[6] = 0
			verticies[7] = 0
			//
			verticies[9] = 0
			verticies[10] = 0
		}
	}
}

func cursorPosCallback(width, height int) glfw.CursorPosCallback {
	return func(w *glfw.Window, xpos float64, ypos float64) {
		if !btnPressed {
			return
		}

		if xpos < 0 {
			xpos = 0
		} else if int(xpos) > width {
			xpos = float64(width)
		}

		if ypos < 0 {
			ypos = 0
		} else if int(ypos) > height {
			ypos = float64(height)
		}

		selX2 = int(xpos)
		selY2 = int(ypos)

		verticies[0] = float32(selX1)
		verticies[1] = float32(selY1)
		//
		verticies[3] = float32(selX2)
		verticies[4] = float32(selY1)
		//
		verticies[6] = float32(selX2)
		verticies[7] = float32(selY2)
		//
		verticies[9] = float32(selX1)
		verticies[10] = float32(selY2)
	}
}
