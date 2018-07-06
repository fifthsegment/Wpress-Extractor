package main

import (
	"fmt"
	"./inc"
	"os"
)

func main() {
	fmt.Printf("Wpress Extracter.\n")

	if ( len(os.Args) == 2 ){
		pathTofile := os.Args[1]
		fmt.Println(pathTofile);
		archiver, _ := wpress.NewReader(pathTofile)
		_ , err := archiver.Extract();
		if (err!=nil){
			fmt.Println("Error = ");
			fmt.Println(err);
		}else{
			fmt.Println("All done!");
		}

		
		// fmt.Println("total files = ", i, " files read = ", x);
	}else{
		fmt.Println("Inorder to run the extractor please provide the path to the .wpress file as the first argument.");
	}
	
	// wpress.Init(archiver);


}
