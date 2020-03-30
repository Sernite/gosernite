package gosernit

import (
	"bufio"
	"fmt"
	"os"
)

var handler HandleFunc

// Params : senrite params
var Params []string

//Done : Exits app with an error
var Done DoneFunc

// SetHandler : sets handler function
func SetHandler(h HandleFunc){
	handler = h;
}

var send SendFunc
var nitMessage NitMsg

func init(){
	reader := bufio.NewReader(os.Stdin)
	send = func(message string){
		fmt.Println(message)
	}

	nitMessage = func(name string,message string)string{
		fmt.Fprintf(os.Stderr,"nit;;;%s;;;%s\n",name,message)
		res,_,_ :=reader.ReadLine()
		return string(res)
	}

}


// Run : Starts script
func Run(){
	// Set Params
	Params = os.Args[1:]
	
	// Run Handler
	handler(send,nitMessage)
}
