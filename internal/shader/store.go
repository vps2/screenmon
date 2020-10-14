package shader

import "fmt"

type Type int

const (
	Vertex Type = iota
	Fragment
)

//go:generate go-bindata -pkg $GOPACKAGE -o assets.go -prefix "../../assets" ../../assets/shader.vert ../../assets/shader.frag

//Store - это хранилище исходников шейдеров
type Store struct {
	vertSource []byte
	fragSource []byte
}

//New  создаёт новое хранилище исходников шейдеров или возвращает ошибку.
func NewStore() *Store {
	vertSource, err := Asset("shader.vert")
	if err != nil {
		panic(fmt.Errorf("the vertex shader is not found: %w", err))
	}
	vertSource = append(vertSource, '\x00')

	fragSource, err := Asset("shader.frag")
	if err != nil {
		panic(fmt.Errorf("the fragment shader is not found: %w", err))
	}
	fragSource = append(fragSource, '\x00')

	return &Store{
		vertSource: vertSource,
		fragSource: fragSource,
	}
}

//GetShader возвращает исходник шейдера, заданного в параметре типа.
func (s *Store) GetShader(t Type) []byte {
	switch t {
	case Vertex:
		return s.vertSource
	case Fragment:
		return s.fragSource
	default:
		return make([]byte, 0)
	}
}
