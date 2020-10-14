package main
import (
"fmt"
)
func main(){
   p:=fmt.Print
   pl:=fmt.Println
   var sayi1,sayi2,islem float32
   p("Sayı Giriniz:")
   fmt.Scanln(&sayi1)
   p("Sayı Giriniz:")
   fmt.Scanln(&sayi2)
   pl("1)Topla \n2)Çıkart \n3)Çarp \n4)Böl")
   p("İşlem Gir:")
   fmt.Scanln(&islem)
   switch islem {
  case 1:
    p("Sonuç:",sayi1+sayi2 )
  case 2:
    p("Sonuç:",sayi1-sayi2)
  case 3:
    p("Sonuç:",sayi1*sayi2)
  case 4:
    p("Sonuç:",sayi1/sayi2)
   }





}
