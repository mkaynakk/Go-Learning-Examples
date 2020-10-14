package main

import (
"fmt"
"math/rand"
"time"


)
func main() {
    p:=fmt.Print
    rand.Seed(time.Now().UnixNano())
    p("Vize Notunuzu Giriniz:")
    var vize,toplam,final int
    var a bool = true
    fmt.Scanln(&vize)
    for (a == true) {
      final = rand.Intn(101)
      toplam = (vize*40/100) + (final*60/100)
	  if (toplam == 50.0) {
      p("Almanız Gereken Not:",final)
      a = false }
    }
}
