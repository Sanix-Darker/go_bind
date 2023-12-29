package go_bind

import "fmt"

type TEstData struct {
	Data string
	Len  int
}

//export HelloThere
func HelloThere(ex TEstData) TEstData {
	ex.Data = "Hello there from GO code '" + ex.Data + "';"
	ex.Len = len(ex.Data)
	return ex
}

func main() {
	fmt.Printf("Hello world !\n")
}
