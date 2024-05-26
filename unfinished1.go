package main
import "fmt"

type money struct{
    id int
    day int
    month int
    year int
    inout string
    amount int
}
var idcount int = 1
const size = 2568
var array [size]money
func main() {
    options(array)
}

func options(array [size]money){
    var options int
    fmt.Scan(&options)
    switch options{
        case 1:
            viewData(array)
        case 2:
            insertData(array)
    }
}

func viewData(array [size]money){
    var i int
      for i = 0;i <= size;i++ {
          id := array[i].id
          tgl := array[i].day
          bln := array[i].month
          thn := array[i].year
          jenis := array[i].inout
          jumlah := array[i].amount
          if tgl == 0{
              break
          } else {
              fmt.Printf("%v | %v/%v/%v : %v %v \n",id ,tgl,bln,thn,jenis,jumlah)
          }
  }
  options(array)
}

func insertData(array [size]money){
    var i int = 0
    
    for array[i].amount != 0{
        i++
    }
    fmt.Println("Masukkan Tanggal : ")
    fmt.Scan(&array[i].day)
    fmt.Println("Masukkan Bulan : ")
    fmt.Scan(&array[i].month)
    fmt.Println("Masukkan Tahun : ")
    fmt.Scan(&array[i].year)
    fmt.Println("Jenis : ")
    fmt.Scan(&array[i].inout)
    fmt.Println("Masukkan Jumlah : ")
    fmt.Scan(&array[i].amount)
    array[i].id = idcount
    idcount++
    options(array)
}
