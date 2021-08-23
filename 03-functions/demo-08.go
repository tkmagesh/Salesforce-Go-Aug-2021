package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Deferred from main")
	}()
	f1()
	fmt.Println("main invoked")

}

func f1() {
	f1Deffered := func() {
		fmt.Println("Deferred(1) from f1")
	}
	defer f1Deffered()

	defer func() {
		fmt.Println("Deferred(2) from f1")
	}()
	defer func() {
		fmt.Println("Deferred(3) from f1")
	}()
	fmt.Println("f1 invoked")
	f2()

	//operation

}

func f2() {
	defer func() {
		fmt.Println("Deferred from f2")
	}()
	fmt.Println("f2 invoked")
}

/* func processFile(fileName string) {
	//open the file
	defer func (){
		//close the file
	}()
	//while not eof

	//	- read a line
		//if error return

	//  - parse the line
		//if error return


	//  - process the data
		//if error return



	//return the result
}
*/
