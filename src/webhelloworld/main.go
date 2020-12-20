package main

import (
	"fmt"
	"net/http"

	greeting "github.com/tenntenn/greeting"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}

//TRY-奇数偶数判定関数
func judge(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

//TRY-複数戻り値の利用
func swap(x, y int) (int, int) {
	return y, x
}

//TRY-ポインタ
func swap2(pn, pm *int) {
	//*ポインたの先の値参照
	//&ポインタを取得
	var inp int
	inp = *pn
	*pn = *pm
	*pm = inp
	return
}

//TRY-レシーバーに変更を与える
/*type Myint int
func (n *Myint) Inc() { *n += 1 }*/

func main() {

	fmt.Pritln(greeting.Do())

	//＊＊＊＊＊＊Gopher道場4＊＊＊＊＊//

	//＊＊＊＊＊＊Gopher道場３＊＊＊＊＊//

	/* TRY-組み込み型（数値）
	var sum int

	sum = 5 + 6 + 3

	var avg float64 = float64(sum / 3)

	if avg > 4.5 {
	  println("good")
	}
	*/

	//TRY-組み込み型（真偽）

	/*
	   var a, b, c bool

	   if a && b || !c {
	     println("true")
	   } else {
	     println("false")
	   }
	*/

	//TRY-スライスの理由
	/*
	   var sum int
	   ns := []int{19, 86, 1, 12}
	   for i := 0; i < cap(ns); i++ {
	     sum += ns[i]
	   }
	   println(sum)
	*/

	//TRY-ユーザ定義型の利用
	/*
	   t := time.Now().UnixNano() //現在時刻
	   rand.Seed(t)
	   s := rand.Intn(101) //0-１０１
	   type kekka struct
	*/

	//TRY-奇数偶数判定関数
	/*
	   for i := 1; i < 101; i++ {
	     if judge(i) {
	       fmt.Printf("%d-偶数\n", i)
	     } else {
	       fmt.Printf("%d-奇数\n", i)
	     }
	   }
	*/

	//TRY-複数戻り値の利用
	/*
	   n, m := swap(10, 20)
	   println(n, m)
	*/

	//TRY-ポインタ
	/*
	   n, m := 10, 20
	   swap2(&n, &m)
	   println(n, m)
	*/

	//TRY-レシーバーに変更を与える
	/*
		var n Myint = 0

		println(n)
		n.Inc()
		println(n)
	*/

	//＊＊＊＊＊Gopher 道場２＊＊＊＊＊//

	/* TRY^奇数と偶数
	for i := 1; i < 101; i++ {
		if i%2 == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}
	*/

	/* TRY-おみくじプログラム

	t := time.Now().UnixNano() //現在時刻
	rand.Seed(t)
	s := rand.Intn(6) //0-5

	switch s + 1 {

	case 6:
		fmt.Printf("大吉")
	case 5, 4:
		fmt.Printf("中吉")
	case 3, 2:
		fmt.Printf("吉")
	default:
		fmt.Printf("凶")
	}

	*/

}
