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
func main() {
    var array [size]money
    options(array)
}

func options(array [size]money){
    var options int
    fmt.Println("==================================")
    fmt.Println("1. Tampilkan Data")
    fmt.Println("2. Masukkan Data")
    fmt.Println("3. Edit Data")
    fmt.Println("==================================")
    fmt.Print("Masukkan Opsi : ")
    fmt.Scan(&options)
    switch options{
        case 1:
            viewData(array)
        case 2:
            insertData(array)
        case 3:
            editData(array)
    }
}

func viewData(array [size]money){
    var i int
    i = 0
    fmt.Println("====== Data ======")
    for i <= size && array[i].id != 0{
        id := array[i].id
        tgl := array[i].day
        bln := array[i].month
        thn := array[i].year
        jenis := array[i].inout
        jumlah := array[i].amount
        fmt.Printf("%v | %v/%v/%v : %v %v \n",id ,tgl,bln,thn,jenis,jumlah)
      }
      options(array)
}

func insertData(array [size]money){
    var i int = 0
    
    for array[i].id != 0{
        i++
    }
    fmt.Println("==== Masukkan Data ====")
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

func editData(array [size]money){
    var x, index, i int
    i = 0
    index = -1
    fmt.Println("====== Edit Data ======")
    fmt.Print("Masukkan ID Data yang diedit : ")
    fmt.Scan(&x)
    for i <= idcount && index == -1{
        if x == array[i].id {
            index = i
        }
        i++
    }
    fmt.Println("Edit Tanggal : ")
    fmt.Scan(&array[index].day)
    fmt.Println("Edit Bulan : ")
    fmt.Scan(&array[index].month)
    fmt.Println("Edit Tahun : ")
    fmt.Scan(&array[index].year)
    fmt.Println("Jenis : ")
    fmt.Scan(&array[index].inout)
    fmt.Println("Edit Jumlah : ")
    fmt.Scan(&array[index].amount)
    options(array)
}
