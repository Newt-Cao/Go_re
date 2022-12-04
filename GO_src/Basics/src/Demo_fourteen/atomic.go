/*******************************************************
 *	                       _oo0oo_
 *	                      o8888888o
 *	                      88" . "88
 *	                      (| -_- |)
 *	                      0\  =  /0
 *	                    ___/`---'\___
 *	                  .' \\|     |// '.
 *	                 / \\|||  :  |||// \
 *	                / _||||| -:- |||||- \
 *	               |   | \\\  - /// |   |
 *	               | \_|  ''\---/''  |_/ |
 *	               \  .-\__  '-'  ___/-. /
 *	             ___'. .'  /--.--\  `. .'___
 *	          ."" '<  `.___\_<|>_/___.' >' "".
 *	         | | :  `- \`.;`\ _ /`;.`/ - ` : | |
 *	         \  \ `_.   \_ __\ /__ _/   .-` /  /
 *	     =====`-.____`.___ \_____/___.-`___.-'=====
 *	                       `=---='
 *
 *
 *	     ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *	           佛祖保佑     永不宕机     永无BUG
 *
 *	Author       : Mr.Chan
 *	Date         : 2022-08-30 15:15:31
 *	LastEditors  : Mr.Chan
 *	LastEditTime : 2022-08-30 15:47:33
 *	FilePath     : /vivianchan.cn/Basics/src/Demo_fourteen/atomic.go
 *******************************************************/

package main

import (
	"fmt"
	"sync/atomic"
)

// 同一个变量
var i int32 = 100

// 增减
func atomic_Add() {
	atomic.AddInt32(&i, 1)
	fmt.Println(i) // 101
	atomic.AddInt32(&i, -1)
	fmt.Println(i) // 100
}

//读
func atomic_Read() {
	l := atomic.LoadInt32(&i)
	fmt.Println(l) // 100
}

//写
func atomic_Write() {
	atomic.StoreInt32(&i, 200)
	fmt.Println(i) // 200
}

// cas，比较并交换，先和原值比较，相同转换失败，不同是转换，别的线程干扰时也会转换失败
func atomic_CompareSwap() {
	b := atomic.CompareAndSwapInt32(&i, 200, 100)
	if b {
		fmt.Println(i)
	} else {
		fmt.Println("转换失败")
	}
}

func main() {
	atomic_Add()
	atomic_Read()
	atomic_Write()
}
