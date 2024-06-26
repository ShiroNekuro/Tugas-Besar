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
	fmt.Println("4. Sort by Amount")
	fmt.Println("5. Sort by Date")
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
		case 4:
			sortAmount(array)
	 	case 5:
	 		sortDate(array)
	}
}

func viewData(array [size]money){
    var i int
    i = 0
    fmt.Println("====== Data ======")
    for i < idcount - 1{
        id := array[i].id
        tgl := array[i].day
        bln := array[i].month
        thn := array[i].year
        jenis := array[i].inout
    	jumlah := array[i].amount
        fmt.Printf("%v | %v/%v/%v : %v %v \n",id ,tgl,bln,thn,jenis,jumlah)
		i++
    }
    options(array)
}

func insertData(array [size]money){
    var i int = 0
    for array[i].id != 0{
		//mencari index yang kosong untuk memasukkan data
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
    for i <= idcount{
		//untuk mencari index yang akan diedit
        if x == array[i].id {
            index = i
			break
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

//Sorting menggunakan insertion sort
func sortAmount(array [size]money){
    var sortby, i, j int
	fmt.Println("======= Sort By Amount =======")
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")
    fmt.Scan(&sortby)
	switch sortby{
		case 1:
			//ascending
			for i = 0; i < idcount - 1; i++ {
				for j = i; j > 0 && array[j-1].amount > array[j].amount; j-- {
					//swap
					array[j], array[j-1] = array[j-1], array[j]
				}
			}
			options(array)
		case 2:
			//descending
			for i = 0; i < idcount - 1; i++ {
				for j = i; j > 0 && array[j-1].amount < array[j].amount; j-- {
					//swap
					array[j], array[j-1] = array[j-1], array[j]
				}
			}
			options(array)
	}
}

//Sorting menggunakan Selection sort
func sortDate(array [size]money){
    var sortby, i, j, minmax int
	fmt.Println("======= Sort By Date =======")
	fmt.Println("1. Oldest")
	fmt.Println("2. Newest")
    fmt.Scan(&sortby)
	switch sortby{
		case 1:
		//Ascending
			for i = 0; i <= idcount - 1; i++{
				minmax = i
				for j = i + 1; j <= idcount - 1;j++{
					if array[i].year > array[j].year {
						minmax = j
					} else if array[i].year == array[j].year {
						if array[i].month > array[j].month {
							minmax = j
						}else if array[i].month == array[j].month {
							if array[i].day > array[j].day {
								minmax = j
							}
						}
					}	
				}
			array[i], array[minmax] = array[minmax], array[i]
			}
			options(array)

		case 2:
		//Descending
			for i = 0; i <= idcount - 1; i++{
				minmax = i
				for j = i + 1; j <= idcount - 1;j++{
					if array[i].year < array[j].year {
						minmax = j
					} else if array[i].year == array[j].year {
						if array[i].month < array[j].month {
							minmax = j
						}else if array[i].month == array[j].month {
							if array[i].day < array[j].day {
								minmax = j
							}
						}
					}	
				}
			array[i], array[minmax] = array[minmax], array[i]
			}
			options(array)
		}
}
