/**
RSA签名算法

（1）参数产生
选择两个满足需要的大素数P和Q，计算n=p*q, fi(n)=(p-1)(q-1)
选择一个整数e，满足1<e<fi(n),且gcd(fi(n), e) = 1。通过d*e=1 mod fi(n)计算出d
以{e,n}为公开秘钥，{d,n}为秘密秘钥

（2）签名过程
假设要签名消息为m,对m签名为s=m^d mod n

（3）验证过程
m = s^e mod n

 */

// @title RSA签名、验证算法
// @Description
// @Author 201983290527 2021.11.01 14:08
// @Update 201983290527 2021.11.01 19:25
package main

import (
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"
)

// @title    generateParams
// @description   产生RSA签名算法所需的参数（e, n, d），并返回素数生成失败时的异常
// @auth      201983290527
// @param     
// @return    e       int64         公开秘钥
//			  n       int64         公开秘钥、秘密秘钥	
//			  d       int64         秘密秘钥
//            er      error         错误信息（没有时为nil）
func generateParams() (e int64, n int64, d int64, er error) {

	var p, q int64
	er = nil
	var err error
	// 生成素数
	p, err = generatePrime(40,70)
	// 素数生成异常处理
	if err != nil{
		err = errors.New("素数p生成失败")
		return
	}
	q, err = generatePrime(80, 120)
	// 素数生成异常处理
	if err != nil{
		err =  errors.New("素数q生成失败")
		return
	}

	// 生成n
	n = p * q

	// 生成e
	for ;;{
		e = chooseInt(1, (p-1)*(q-1))
		if gcd(e, (p-1)*(q-1))==1 {
			break
		}
	}


	// 生成d
	d = findModOne(e, (p-1)*(q-1))

	return

}

// @title    sign
// @description   使用秘密密钥（d,n）对明文m进行加密
// @auth      201983290527
// @param     m       int64         待签名明文
//			  d       int64         秘密秘钥
//            n       int64         秘密秘钥
// @return    sInt   int64         签名后的明文
//			          error         错误信息（没有时为nil） 
func sign(m, d, n int64) (int64, error) {
	//（2）签名过程
	//假设要签名消息为m,对m签名为s=m^d mod n

	s := big.NewInt(n).Mod(pow(m, d), big.NewInt(n)).String()
	fmt.Println("签名后的明文为：", s)
	sInt, err2 := strconv.ParseInt(s, 10, 64)
	//print(s_int, err2)
	if err2 != nil{
		return 0, errors.New("string-->int64转化失败！")
	}
	return sInt, nil
}

// @title    confirm
// @description   使用公开密钥（e,n）对密文s_int进行验证
// @auth      201983290527
// @param     sInt    int64         待验证的明文
//			  e       int64         公开秘钥
//            n       int64         公开秘钥
// @return    m2Int   int64         验证后的明文
func confirm(sInt, e, n int64) int64 {
	//（3）验证过程
	//m = s^e mod n

	m2 := big.NewInt(n).Mod(pow(sInt, e), big.NewInt(n)).String()
	//fmt.Println("验证时明文为：", m2)
	m2Int, err := strconv.ParseInt(m2, 10, 64)
	if err != nil{
		print(err)
	}
	return m2Int
}


func pow(x int64, y int64) *big.Int {
	var i int64
	result := big.NewInt(1)
	for i=0;i<y;i++{
		//print(result.String()+"\n")
		result = result.Mul(result, big.NewInt(x))
	}

	return result
}


// 求解x*e = 1 mod fi(n)
func findModOne(e int64, modNum int64) int64 {
	var i int64 = 1
	var div int64 // 商
	var remainder int64 // 余数
	var middleNum int64
	for ;;i++{
		middleNum = e * i
		if middleNum < modNum {
			remainder = middleNum // 当两数相乘小于模数时，结果为两数相乘结果
		}else if middleNum >= modNum {
			div = middleNum / modNum
			remainder = middleNum - modNum*div // 两数相乘大于模数
		}

		if remainder == 1 {
			return i
		}


	}
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

// 给定范围选取整数
func chooseInt(min int64, max int64) int64{
	rand.Seed(time.Now().UnixNano())
	return  rand.Int63n(max-min)+min
}

// 给定num1,num2求最大公约数
func gcd(num1 int64, num2 int64) int64{
	// 欧几里得算法

	// 确保大的数为num1
	if num1 < num2{
		num1, num2 = num2, num1
	}
	var div int64  // 向下取整后的商
	var remainder int64  // 余数

	for ; ;{

		div = num1 / num2 // 计算商
		remainder = num1 - div*num2  // 计算余数

		//fmt.Println(div)
		//fmt.Println(remainder)

		// 较小数赋给较大数位置，余数赋给较小数位置
		num1 = num2
		num2 = remainder

		// 当余数为0时，跳出循环
		if remainder == 0 {
			break
		}
	}

	return num1
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



func main() {
	//fmt.Println(gcd(27,36))
	////fmt.Println(judge_prime(157))
	//fmt.Println(generate_prime(2,20))
	//fmt.Println(choose_int(50,60))

	var M string
	var m []rune
	var sSign []int64
	var mConfirm []rune

	// 读取待签名明文
	fmt.Println("请输入待签名明文：")
	_, err := fmt.Scanln(&M)
	if err != nil {
		return 
	}

	for i:=0;i<len(M);i++{
		m = append(m, rune(M[i]))
	}
	fmt.Println("输入明文的ASCII码序列", m)

	var es []int64
	var ns []int64
	var ds []int64

	for i:=0;i<len(m);i++{

		// 生成参数
		e, n, d, err:= generateParams()
		// 存储公钥、私钥
		es = append(es, e)
		ns = append(ns, n)
		ds = append(ds, d)

		if err != nil{
			fmt.Println("生成出错")
			return
		}
		//{e,n}为公开秘钥，{d,n}为秘密秘钥
		fmt.Println("公开秘钥：", e, n)
		fmt.Println("秘密秘钥：", e, d)
		//fmt.Println(mod(21, 7))

		// 签名过程
		sInt, err2 := sign(int64(m[i]), d, n)
		if err2 != nil{
			print("签名过程出错")
			print(err2)
			return
		}
		sSign = append(sSign, sInt)
	}

	//fmt.Println(s_sign)

	// 验证过程
	for i:=0;i<len(sSign);i++{
		mConfirm = append(mConfirm, rune(confirm(sSign[i], es[i], ns[i])))
	}
	fmt.Printf("验证后明文为：")
	for i:=0;i<len(mConfirm);i++{
		fmt.Printf("%s", string(mConfirm[i]))
	}



	//// 签名过程
	//s_int, err2 := sign(9, d, n)
	//// 异常处理
	//if err2 != nil{
	//	print("签名过程出错")
	//	print(err2)
	//	return
	//}
	//
	//// 认证过程
	//fmt.Println("认证过程结束明文：",confirm(s_int,e,n))







	//a := new(big.Int)
	//a,_ = a.SetString("123456234242452452042020522452452452245245278910", 10)
	//fmt.Println(a.String())
	//b := big.NewInt(2)
	//b = b.Add(a, b)
	//fmt.Println(b.String())
	//fmt.Println(reflect.TypeOf(b))

	//print(pow(2,3).String())

	//a := big.NewInt(3)
	//b := big.NewInt(2)
	//b = b.Mod(b, a)
	//print(b.String())
}
