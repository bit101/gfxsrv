package main

import "gfxsrv/gfxsrv"


func main() {
	gfxsrv.NewServer("../project", "../site", 8083)
}


