package main

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

/**
Alice Bob分别选择一个随机数a, b(0≤ a,b ≤p-2)
目的：计算出共享密钥
 */

var A = make(chan int64, 1)
var g = make(chan int64, 1)
var p = make(chan int64, 1)
var B = make(chan int64, 1)

func Alice()  {
	var prime int64  // 素数
	var a int64  // 随机数a
	var generator int64  // 生成元
	var A_middle int64  // g^a mod p
	var B_middle int64  // g^b mod p
	var k int64  // B^a mod p

	prime,_ = generatePrime(10, 100)
	//fmt.Println("素数为",prime)
	a = chooseInt(0, prime - 2)
	//print("a is ", a)
	//fmt.Println("Alice随机数为：",a)
	generator = findGenerator(prime)
	A_middle, _ = strconv.ParseInt(big.NewInt(prime).Mod(pow(generator, a), big.NewInt(prime)).String(),10, 64)
	//fmt.Println("g^a mod p",A_middle)
	fmt.Println("Alice正在传输g,p,A", A_middle, generator, prime)
	g <- generator
	p <- prime
	A <- A_middle

	fmt.Println("Alice正在等待B")
	B_middle = <- B
	fmt.Println("Alice已取得B", B_middle)

	k,_ = strconv.ParseInt(big.NewInt(prime).Mod(pow(B_middle, a), big.NewInt(prime)).String(),10,64)
	fmt.Println("Alice协商获得密钥", k)
}

func Bob()  {
	var prime int64  // 素数
	var generator int64  // 生成元
	var A_middle int64  // g^a mod p
	var b int64  // 随机数b
	var B_middle int64  // g^b mod p
	var k int64  // A^b mod p

	fmt.Println("Bob正在等待Alice传输a,g,p")
	prime = <- p
	generator = <- g
	A_middle = <- A
	fmt.Println("Bob已取得a,g,p", A_middle, generator, prime)
	b = chooseInt(0, prime - 2)
	//print("b is ", b)
	B_middle,_ = strconv.ParseInt(big.NewInt(prime).Mod(pow(generator, b), big.NewInt(prime)).String(),10,64)

	fmt.Println("Bob向Alice传输B", B_middle)
	B <- B_middle

	k,_ = strconv.ParseInt(big.NewInt(prime).Mod(pow(A_middle, b), big.NewInt(prime)).String(),10,64)

	fmt.Println("Bob协商得到密钥", k)
}

// 给定上下限生成素数
func generatePrime(min int64, max int64) (int64, error) {
	if min==0 {
		return 0, errors.New("范围最小不可为0")
	}
	for i:=min;i<=max;i++{
		if judgePrime(i){
			return i, nil
		}
	}
	return 0,errors.New("范围中不存在")

}

// 判断素数
func judgePrime(num int64) bool {
	if num%2==0 {
		return false
	} else if num%3==0 {
		return false
	}else if num%5 ==0 {
		return false
	}else {
		var i int64 = 2
		for ;i<num;i++{
			if num%i==0 {
				//fmt.Println(i)
				return false
			}
		}
	}
	return true
}

// 给定范围选取整数
func chooseInt(min int64, max int64) int64{
	rand.Seed(time.Now().UnixNano())
	return  rand.Int63n(max-min)+min
}

//寻找生成元
func findGenerator(prime int64) int64 {
	// 从2-p-2遍历g^(p-1) = 1 (mod p)
	var i int64 = 2
	var primeBig *big.Int = big.NewInt(prime)
	//fmt.Println(primeBig)
	for ;i<prime-1;i++{
		//fmt.Println("i", i)
		//fmt.Println("平方", pow(i, prime - 1))
		// 遍历i^(p-1) mod p
		if primeBig.Mod(pow(i, prime - 1), primeBig).String() == "1"{
			return i
		}
	}
	return -1
}

// 大整数指数
func pow(x int64, y int64) *big.Int {
	var i int64
	result := big.NewInt(1)
	for i=0;i<y;i++{
		//print(result.String()+"\n")
		result = result.Mul(result, big.NewInt(x))
	}

	return result
}

func main() {
	//print(pow(2,2).String())
	//print(findGenerator(17))
	go Alice()
	go Bob()
	for  {
		
	}
}
