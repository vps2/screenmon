package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Program struct {
	id uint32
}

func NewProgram(vertexSource, fragmentSource []byte) (*Program, error) {
	vertexCode := string(vertexSource)
	vertexShader, err := compileShader(vertexCode, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	fragmentCode := string(fragmentSource)
	fragmentShader, err := compileShader(fragmentCode, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Program{id: program}, nil
}

func (p *Program) Use() {
	gl.UseProgram(p.id)
}

func (p *Program) SetUniform1(name string, value interface{}) error {
	var loc int32
	if loc = gl.GetUniformLocation(p.id, gl.Str(name+"\x00")); loc == -1 {
		return fmt.Errorf("wrong uniform name '%s'", name)
	}

	switch t := value.(type) {
	case float32:
		gl.Uniform1f(loc, value.(float32))
	default:
		return fmt.Errorf("unexpected type %T", t)
	}

	return nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
