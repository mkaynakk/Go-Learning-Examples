package main
import(
"fmt"
"math/rand"
"time"
)
func main(){
  p:=fmt.Print
  pl:=fmt.Println
  var buldu bool=false
  var tahmin,sayi,sayac int
  rand.Seed(time.Now().UnixNano())
  sayi = rand.Intn(100)
  p("Tahmin sayisi giriniz:")
  fmt.Scanln(&tahmin)
  sayac=1
  for (buldu == false){
  if (tahmin == sayi){
  p("Tebrikler ",sayac,". Tahminde Buldunuz.")
  buldu=true
} else if(tahmin<sayi){
  pl("Sayıyı Büyültün")
  p("Tahmin sayisi giriniz:")
  fmt.Scanln(&tahmin)
} else{
  pl("Sayıyı Küçültün")
  p("Tahmin sayisi giriniz:")
  fmt.Scanln(&tahmin)
}
  sayac = sayac+1
}
}
