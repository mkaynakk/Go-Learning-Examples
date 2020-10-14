package main
import(
"fmt"
)
func main(){
  var sayi,toplam int

fmt.Print("Bir Sayı Giriniz:")
fmt.Scanln(&sayi)
for i:=0; i<sayi; i++{
 toplam += i*sayi
}
fmt.Print(toplam)
}
