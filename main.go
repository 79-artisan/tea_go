package main

import "fmt"

func main() {
	data := []byte("{\"code\":2000, \"msg\":\"ok\", \"dat\":[]}")
	encryptData := Encrypt(data)
	fmt.Println("encryptData：", encryptData)

	dectyptData := Decrypt(encryptData)
	fmt.Println("dectyptData：", string(dectyptData))
}
