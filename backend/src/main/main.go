package main

import "fmt"
import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	url := "https://firebasestorage.googleapis.com/v0/b/project-leo-mvp.appspot.com/o/avatars%2Fbb682b66-8f1a-4f33-943f-70fe904705de?alt=media&token=cc577ed6-b7e4-43ac-be31-5ae9c7d1fd30"

	fmt.Println(url)
}
