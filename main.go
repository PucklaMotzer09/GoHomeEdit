package main

import (
	"github.com/PucklaMotzer09/GoHomeEngine/src/frameworks/GTK"
	"github.com/PucklaMotzer09/GoHomeEngine/src/gohome"
	"github.com/PucklaMotzer09/GoHomeEngine/src/renderers/OpenGL"
)

func main() {
	gohome.MainLop.Run(&framework.GTKFramework{
		UseWholeWindowAsGLArea: false,
		MenuBarFix: true,
	},&renderer.OpenGLRenderer{},1280,720,"GoHomeEdit",&EditScene{})
}
