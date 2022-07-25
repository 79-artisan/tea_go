package main

var K = []uint32{1280791999, 1112888951, 1970042213, 1112892528}

func Encrypt(clear []byte) []byte {
	LenIn := len(clear)
	Len := LenIn
	if (LenIn % 8) > 0 {
		Len = (LenIn/8 + 1) << 3
	}
	Loop := Len / 8
	aa := Loop * 2
	v := []uint32{}
	for i := 0; i < aa; i++ {
		v = append(v, 0)
	}
	v = pack(clear, v)
	for i := 0; i < Loop; i++ {
		VV, KK := TeaEncrypt(v, K, i*2)
		v[i*2] = VV
		v[i*2+1] = KK
	}
	return UnPack(v, Loop*8)
}

func pack(src []byte, dest []uint32) []uint32 {
	var shift = 0
	var j = 0
	dest[j] = 0
	srcLen := len(src)
	for i := 0; i < srcLen; i++ {
		dest[j] = dest[j] | ((uint32(src[i]) & 0xff) << shift)
		if shift == 24 {
			shift = 0
			j++
			if j < len(dest) {
				dest[j] = 0
			}
		} else {
			shift += 8
		}
	}
	return dest
}

func TeaEncrypt(v []uint32, K []uint32, idx int) (uint32, uint32) {
	v0 := v[idx]
	v1 := v[idx+1]
	var sum uint32
	k0 := K[0]
	k1 := K[1]
	k2 := K[2]
	k3 := K[3]
	for i := 0; i < 32; i++ {
		sum = sum + 0x9e3779b9
		v0 += ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		v1 += ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
	}
	v[idx] = v0
	v[idx+1] = v1
	return v0, v1
}

func UnPack(src []uint32, length int) []byte {
	var dest []byte
	for i := 0; i < length; i++ {
		dest = append(dest, 0)
	}
	i, count := 0, 4
	for j := 0; j < length; j++ {
		count--
		dest[j] = byte(src[i] >> (24 - (8 * count)) & 0xff)
		if count == 0 {
			count = 4
			i++
		}
	}
	var l = length - 1
	var c = 0
	for {
		if l >= 0 {
			if int(dest[l]) == 0 {
				c++
			} else {
				break
			}
			l--
		} else {
			break
		}
	}

	if length-c > len(dest) {
		return []byte{}
	}

	return dest[0 : length-c]
}

func Decrypt(crypt []byte) []byte {
	var LenIn = len(crypt)
	Len := LenIn
	if (LenIn % 8) > 0 {
		Len = (LenIn/8 + 1) << 3
	}
	var Loop = Len / 8
	aa := Loop * 2
	v := []uint32{}
	for i := 0; i < aa; i++ {
		v = append(v, 0)
	}
	v = pack(crypt, v)
	for i := 0; i < Loop; i++ {
		VV, KK := TeaDecrypt(v, K, i*2)
		v[i*2] = VV
		v[i*2+1] = KK
	}
	return UnPack(v, Loop*8)
}

func TeaDecrypt(v []uint32, K []uint32, idx int) (uint32, uint32) {
	v0 := v[idx]
	v1 := v[idx+1]
	var sum uint32
	sum = 0xC6EF3720
	k0 := K[0]
	k1 := K[1]
	k2 := K[2]
	k3 := K[3]
	for i := 0; i < 32; i++ {
		v1 -= ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
		v0 -= ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		sum = sum - 0x9e3779b9
	}
	v[idx] = v0
	v[idx+1] = v1
	return v0, v1
}
