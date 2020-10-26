package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// FILE HANDLING IN GO.
func main() {

	create()
	open()
	openFile()
	renameNdMove()
}

/*
  1. CREATE
  func create(name string)( *file, error)
*/
func create() {
	newFile, err := os.Create("vik.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(newFile)
	newFile.Close()
}

/*
2. OPEN and CLOSE file
func open(name string) (*file error)
*/
func open() {
	file, err := os.Open("vik.txt")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(file)
	file.Close()
}

/*
3. OPENFILE
func openFile(name string, flog int, perm FileMode) (*File error)
// Use these attributes individually or combined
    // with an OR for second arg of OpenFile()
    // e.g. os.O_CREATE|os.O_APPEND
    // or os.O_CREATE|os.O_TRUNC|os.O_WRONLY

    // os.O_RDONLY // Read only
    // os.O_WRONLY // Write only
    // os.O_RDWR // Read and write
    // os.O_APPEND // Append to end of file
    // os.O_CREATE // Create is none exist
    // os.O_TRUNC // Truncate file when opening
*/
func openFile() {
	file, err := os.OpenFile("vik.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	file.Close()
	log.Println(file)
}

/*
4.GETFILEINFO	
*/

func getFileInfo(){
	fileInfo, err := os.Stat("vik.txt")
	if err != nil {
		log.Fatalln(err)
	}
fmt.Println(fileInfo)
}

/*
5.Rename and move  a file
*/

func renameNdMove(){
	err:= os.Rename("vik.txt","silk.txt")
	if err != nil {
		log.Fatalln(err)
	}

}
/*
6. Delete a file
*/

func del(){
	err:= os.Remove("silk.txt")
	if err != nil {
		log.Fatalln(err)
	}
}

/*
7. Check if file exists
*/
func check(){
	fileInfo, err:= os.Stat("silk.txt")
	if err != nil {
		if os.IsNotExist(err){

			log.Fatalln("file does not exists.")
		}
	}
	log.Println(fileInfo)
}

/*
8. Check read and write permisiions
*/
func checkReadWritePem(){
	//write permission
	file, err:= os.OpenFile("text.txt", os.O_WRONLY, 0666)
	if err != nil {
		if os.IsPermission(err){
			log.Println("Error: write permission denied.")
		}
	}
	file.Close()

	// read permission
	file, err := os.OpenFile("silk.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsPermission(err){
			log.Fatalln("Error : read permission denied.")
		}
	}
	file.Close()
}

/*
9.Change permission
*/
func chngPerm(){
	// using linux style
	err :=os.Chmod("silk.txt", 0777)
	if err != nil {
		log.Println(err)
	}
	// change ownership
	err= os.Chown("silk.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatalln(err)
	}
	// change timestamps
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
    lastAccessTime := twoDaysFromNow
    lastModifyTime := twoDaysFromNow
    err = os.Chtimes("test.txt", lastAccessTime, lastModifyTime)
    if err != nil {@
        log.Println(err)
    }
}

/*
10. Copy a file
*/
func copy(){
	 // Open original file
	 originalFile, err := os.Open("test.txt")
	 if err != nil {
		 log.Fatal(err)
	 }
	 defer originalFile.Close()
 
	 // Create new file
	 newFile, err := os.Create("test_copy.txt")
	 if err != nil {
		 log.Fatal(err)
	 }
	 defer newFile.Close()
 
	 // Copy the bytes to destination from source
	 bytesWritten, err := io.Copy(newFile, originalFile)
	 if err != nil {
		 log.Fatal(err)
	 }
	 log.Printf("Copied %d bytes.", bytesWritten)
	 
	 // Commit the file contents
	 // Flushes memory to disk
	 err = newFile.Sync()
	 if err != nil {
		 log.Fatal(err)
	 }
}