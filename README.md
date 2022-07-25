# tea_go
Golang tea 加密算法


	data := []byte("{\"code\":2000, \"msg\":\"ok\", \"dat\":[]}")
	encryptData := Encrypt(data)
	fmt.Println("encryptData：", encryptData)

	dectyptData := Decrypt(encryptData)
	fmt.Println("dectyptData：", string(dectyptData))
