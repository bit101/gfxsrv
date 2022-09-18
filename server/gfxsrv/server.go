package gfxsrv

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

type GfxServer struct {
	projectPath string
	sitePath string
	port int
}



func NewServer(project, site string, port int) {
	gs := GfxServer{project, site, port}

	http.Handle("/", http.FileServer(http.Dir(gs.sitePath)))

	http.HandleFunc("/gen", gs.generateImage)
	http.HandleFunc("/compile", gs.compile)

	log.Printf("Server started on port %d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

func (gs *GfxServer) generateImage(w http.ResponseWriter, r *http.Request) {
	gs.generate()
}

func (gs *GfxServer) generate() {
	log.Println("Generating image...")
	cmd := exec.Command("go", "run", "main.go")
	cmd.Dir = gs.projectPath
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func (gs *GfxServer) compile(w http.ResponseWriter, r *http.Request) {
	log.Println("Compiling tool...")
	cmd := exec.Command("go", "build", "main.go")
	cmd.Dir = gs.projectPath
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}

	gs.generate()
}
