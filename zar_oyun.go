package main
import(
"fmt"
"math/rand"
"time"
)
func main(){
  rand.Seed(time.Now().UnixNano())
  p:=fmt.Print
  pl:=fmt.Println
  s:=fmt.Scanln
  var oyuncu1,oyuncu2 string
  var skor1,skor2,beraberlik int
  p("İsim Giriniz:")
  s(&oyuncu1)
  p("İsim Giriniz:")
  s(&oyuncu2)
  for i:=0; i<30; i++{
    zar1:=rand.Intn(6)+1
    zar2:=rand.Intn(6)+1
  pl(oyuncu1,":\n",zar1)
  pl(oyuncu2,":\n",zar2)
  pl("<--------------------->")
    if (zar1>zar2){
 skor1++

}else if(zar1<zar2){
    skor2++

  }else{
    beraberlik++
  }
}
  pl(oyuncu1," Skoru:",skor1)
  pl(oyuncu2," Skoru:",skor2)
  pl("Beraberlik:", beraberlik)
  pl("<------------------------------->")
if (skor1>skor2){
    p(oyuncu1," kazandı.\nSkoru:",skor1)
  }else if(skor1<skor2){
      p(oyuncu2," kazandı.\nSkoru:",skor2)
  }else{
    p("Berabere Bitti")
  }
}
