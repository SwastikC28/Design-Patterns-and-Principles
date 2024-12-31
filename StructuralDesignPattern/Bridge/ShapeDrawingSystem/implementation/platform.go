package implementation

import "fmt"

type Platform interface {
	RenderShape(shape string)
}

// ==============================================================

type WindowsPlatform struct{}

func NewWindowsPlatform() Platform {
	return &WindowsPlatform{}
}

func (platform *WindowsPlatform) RenderShape(shape string) {
	str := fmt.Sprintf("Rendering shape %s on Windows", shape)
	fmt.Println(str)
}

// ==============================================================

type MacOSPlatform struct{}

func NewMacOSPlatform() Platform {
	return &MacOSPlatform{}
}

func (platform *MacOSPlatform) RenderShape(shape string) {
	str := fmt.Sprintf("Rendering shape %s on MacOS", shape)
	fmt.Println(str)
}

// ==============================================================

type LinuxPlatform struct{}

func NewLinuxPlatform() Platform {
	return &LinuxPlatform{}
}

func (platform *LinuxPlatform) RenderShape(shape string) {
	str := fmt.Sprintf("Rendering shape %s on Linux", shape)
	fmt.Println(str)
}
