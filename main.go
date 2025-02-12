package main

import (
	"log"
	"math"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Minecraft Clone"
)

var (
	yaw, pitch       float64
	xPos, yPos, zPos float64
	firstMouse       = true
	lastX            = float64(windowWidth) / 2.0
	lastY            = float64(windowHeight) / 2.0
	cameraSpeed      = 0.05
	cameraFront      = mgl32.Vec3{0, 0, -1}
	cameraPos        = mgl32.Vec3{0, 0, 3}
	cameraUp         = mgl32.Vec3{0, 1, 0}
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// Set OpenGL version
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.True)

	// Create a window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window:", err)
	}
	window.MakeContextCurrent()

	// Capture mouse
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	// Set callbacks
	window.SetCursorPosCallback(mouseCallback)
	window.SetKeyCallback(keyCallback)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize gl:", err)
	}

	gl.Enable(gl.DEPTH_TEST)

	// Main loop
	for !window.ShouldClose() {
		processInput(window)

		// Clear the screen
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Render a cube
		renderCube()

		// Swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func renderCube() {
	vertices := []float32{
		// positions          // colors
		-0.5, -0.5, -0.5, 1, 0, 0,
		 0.5, -0.5, -0.5, 0, 1, 0,
		 0.5,  0.5, -0.5, 0, 0, 1,
		-0.5,  0.5, -0.5, 1, 1, 0,
		-0.5, -0.5,  0.5, 1, 0, 1,
		 0.5, -0.5,  0.5, 0, 1, 1,
		 0.5,  0.5,  0.5, 1, 1, 1,
		-0.5,  0.5,  0.5, 0.5, 0.5, 0.5,
	}

	indices := []uint32{
		0, 1, 3, 1, 2, 3, // front
		4, 5, 7, 5, 6, 7, // back
		0, 1, 4, 1, 5, 4, // bottom
		2, 3, 6, 3, 7, 6, // top
		0, 3, 4, 3, 7, 4, // left
		1, 2, 5, 2, 6, 5, // right
	}

	var VAO, VBO, EBO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.GenBuffers(1, &VBO)
	gl.GenBuffers(1, &EBO)

	gl.BindVertexArray(VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Position attribute
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	// Color attribute
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.BindVertexArray(VAO)
	gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_INT, gl.PtrOffset(0))

	// Cleanup
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &VBO)
	gl.DeleteBuffers(1, &EBO)
	gl.DeleteVertexArrays(1, &VAO)
}

func mouseCallback(window *glfw.Window, xpos, ypos float64) {
	if firstMouse {
		lastX = xpos
		lastY = ypos
		firstMouse = false
	}

	xOffset := xpos - lastX
	yOffset := lastY - ypos
	lastX = xpos
	lastY = ypos

	sensitivity := 0.1
	xOffset *= sensitivity
	yOffset *= sensitivity

	yaw += xOffset
	pitch += yOffset

	if pitch > 89.0 {
		pitch = 89.0
	}
	if pitch < -89.0 {
		pitch = -89.0
	}

	direction := mgl32.Vec3{
		float32(math.Cos(mgl32.DegToRad(float32(yaw))) * math.Cos(mgl32.DegToRad(float32(pitch)))),
		float32(math.Sin(mgl32.DegToRad(float32(pitch)))),
		float32(math.Sin(mgl32.DegToRad(float32(yaw))) * math.Cos(mgl32.DegToRad(float32(pitch)))),
	}
	cameraFront = direction.Normalize()
}

func keyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if key == glfw.KeyEscape && action == glfw.Press {
		window.SetShouldClose(true)
	}
}

func processInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyW) == glfw.Press {
		cameraPos = cameraPos.Add(cameraFront.Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyS) == glfw.Press {
		cameraPos = cameraPos.Sub(cameraFront.Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyA) == glfw.Press {
		cameraPos = cameraPos.Sub(cameraFront.Cross(cameraUp).Normalize().Mul(cameraSpeed))
	}
	if window.GetKey(glfw.KeyD) == glfw.Press {
		cameraPos = cameraPos.Add(cameraFront.Cross(cameraUp).Normalize().Mul(cameraSpeed))
	}
}
