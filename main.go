package main

import "github.com/geosteffanov/bfinterpreter/bf"

func main() {
	//reader  := bufio.NewReader(os.Stdin)
	//
	//input, err := reader.ReadString('\n')
	//
	//fmt.Println(input)
	//fmt.Println(err)
	//interpreter := bf.NewInterpreter(",.")
	//bf.Run(interpreter)

	repl := bf.NewRepl()
	repl.Start()
	//writer := bufio.NewWriter(os.Stdout)
	//
	//for {
	//	writer.Write([]byte("wow"))
	//	writer.Flush()
	//	time.Sleep(4 * time.Second)
	//}

}
