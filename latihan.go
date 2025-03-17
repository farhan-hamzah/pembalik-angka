// Pembalik Angka
package main
import "fmt"
func balik(n int, hasil *int){
    var i int
    *hasil = 0
    for n > 0{
        i = n%10
        *hasil = (*hasil*10) + i 
        n = n/10
    }
    fmt.Print(*hasil)
}

func main(){
    var masukan, keluaran int
    fmt.Scan(&masukan)
    balik(masukan, &keluaran)
}
