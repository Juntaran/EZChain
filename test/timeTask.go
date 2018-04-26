/** 
  * Author: Juntaran 
  * Email:  Jacinthmail@gmail.com 
  * Date:   2018/4/24 15:13
  */

package main

import (
	"time"
	"fmt"
)

func TimingFixBlockchain(duration time.Duration) {
	// 每隔 duration 维护一次链
	ticker := time.NewTicker(time.Second * time.Duration(duration))
	for _ = range ticker.C {
		// 维护链
		fmt.Println("rua")
	}
}

func main() {
	go TimingFixBlockchain(1)
	ch := make(chan int)
	<- ch
}
