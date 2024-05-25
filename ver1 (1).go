// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"
type money struct{
    day int
    month int
    year int
    inout string
    amount int
}
const size = 2568
var array [size]money
func main() {
  i := 0
  while i <= size{
      tgl = array[i].day
      bln = array[i].month
      thn = array[i].year
      jenis = array[i].inout
      jumlah = array[i].amount
      if jumlah == 0{
          break
      } else {
          fmt.Printf("%v/%v/%v : %v %v",tgl,bln,thn,jenis,jumlah)
      }
  }
}