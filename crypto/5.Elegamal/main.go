package main

import (
	"fmt"
	"math"
)

var p int = 97  // 素数
var g int = 5  // 生成元
var x int = 58 // 随机数

// 生成公钥y=g^x
func generatePubkey() int {
	//fmt.Println(float64(g))
	//fmt.Println(float64(x))
	y :=  int(math.Mod(pow(g,x,p), float64(p)))
	return y
}

// 加密
func encrypt(m rune, y int) (C1 int, C2 int)  {
	// 随机生成k， 且0 < k < p-1
	//k := rand.Intn(p-1) + 1
	k := 36
	m_int := int(m)
	// 求解密文对
	C1 = int(math.Mod(pow(g,k,p),float64(p)))  // C1 = g^k mod p
	C2 = int(math.Mod(pow(y,k,p)*math.Mod(float64(m_int),float64(p)),float64(p)))
	//fmt.Printf("密文对（%d, %d）", C1, C2)
	return C1, C2
}

// 解密
func decrypt(C1 int, C2 int) int{
	//fmt.Printf("%d, %d\n", C1, C2)
	// M = C1/c2^x mod p = ((C1 mod p) / (C2^x mod p))mod p
	//fmt.Println(math.Mod(float64(C2), float64(p)))
	//fmt.Println(math.Mod(math.Pow(float64(C1),float64(x)),float64(p)))
	//fmt.Println(fraction_mod(int(math.Mod(float64(C2), float64(p))),int(math.Mod(math.Pow(float64(C1),float64(x)),float64(p))),p))
	//print("\nC1:",C1)
	//print("\nX:",x)
	//print("\nC1^X:",pow(C1,x,p))
	//print("\np",float64(p))
	//fmt.Println("test")
	//fmt.Printf("%.f",math.Mod(pow(C1,x,p),float64(p)))
	//print("\nC2:",C2)
	//fmt.Println("TEST2:%.f\n",math.Mod(float64(C2), float64(p)))
	m2 := fraction_mod(int(math.Mod(float64(C2), float64(p))),
		int(math.Mod(pow(C1,x,p), float64(p))),
		p)
	return m2
}


// 分数模运算
func fraction_mod(a int, b int, p int) int {
	//print("\nb:",b)
	//print("\na:", a)
	for m:=0; ;m++{
		if math.Mod(float64(b*m), float64(p)) == float64(a){
			//print(m)
			return m
		}
	}
}


// 模次方运算
func pow(x int,y int, p int) float64 {
	var result float64 = 1
	for i:=0;i<y;i++{
		result =  math.Mod(result * float64(x), float64(p))
	}
	result = math.Mod(result,float64(p))
	return result
}

func main() {
	fmt.Println("Elgamal算法")
	//fmt.Println(generatePubkey())
	//fmt.Printf("%.f",generatePubkey())
	y := generatePubkey()
	fmt.Println(y)
	fmt.Println("请输入明文消息：")
	var M string
	fmt.Scanln(&M)
	fmt.Printf("待加密的明文为：%s\n", M)
	fmt.Printf("明文转化为ASCII码后：")
	// 将明文转化为ASCII码并放入数组中
	var m []rune
	for i:=0; i < len(M); i++{
		m = append(m, rune(M[i]))
		fmt.Printf("%d ", m[i])
	}
	fmt.Printf("\n加密公钥为：(%d, %d, %d)\n", y,g,p)

	// 加密
	res := make([][2]int, len(m))
	//fmt.Println(res)
	for i:=0; i<len(m);i++{
		res[i][0], res[i][1] = encrypt(m[i], y)
	}
	fmt.Printf("加密后密文对为：%V\n",res)

	//fmt.Printf("%d, %d",len(m),len(res))
	var m2 []rune
	for i:=0; i<len(res); i++{
		m2 = append(m2, rune(decrypt(res[i][0], res[i][1])))
		//fmt.Printf("%d", m2)
	}

	fmt.Printf("解密后ASCII码为：")
	for i:=0;i<len(m2);i++{
		fmt.Printf("%d", m2[i])
	}
	fmt.Println("\n")
	fmt.Printf("解密后明文为：")
	for i:=0;i<len(m2);i++{
		fmt.Printf("%s", string(m2[i]))
	}

	//z1 := pow(5,58, 97)
	//fmt.Printf("%.f\n",z1)
	//z2 := math.Mod(z1,97)
	//fmt.Println(z2)


}
