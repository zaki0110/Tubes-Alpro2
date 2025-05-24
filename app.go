package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Project struct {
	name     string
	category string
	difficulty int
	date     time.Time
}

var Porto []Project


func main() {
	var choice int
	for {
		fmt.Println("\n📊 Manajemen Portofolio Data Science")
		fmt.Println("1. Tambah Portofolio")
		fmt.Println("2. Ubah Project")
		fmt.Println("3. Hapus Project")
		fmt.Println("4. Lihat Semua Portofolio")
		fmt.Println("5. Cari Portofolio")
		fmt.Println("6. Urutan Project")
		fmt.Println("7. Statistik Proyek per Kategori")
		fmt.Println("8. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			tambahPorto()
		case 2:
			ubahPorto()
		case 3:
			hapusPorto()
		case 4:
			showPorto()
		case 5:
			cariPorto()
		case 6:
			urutanPorto()
		case 7:
			statPorto()
		case 8:
			fmt.Println("Terima kasih!")
			return
		case 1234:
			dummy()
		}
	}
}

func dummy()  {
	Porto = append(Porto, Project{"AI Assistant", "AI", 3, time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC)})
	Porto = append(Porto, Project{"Budget Tracker", "Finance", 2, time.Date(2023, 11, 5, 0, 0, 0, 0, time.UTC)})
	Porto = append(Porto, Project{"Chat App", "Communication", 3, time.Date(2023, 8, 20, 0, 0, 0, 0, time.UTC)})
	Porto = append(Porto, Project{"Language Learner", "Education", 2, time.Date(2022, 8, 3, 0, 0, 0, 0, time.UTC)})
	Porto = append(Porto, Project{"Online Portfolio", "Web Development", 1, time.Date(2022, 9, 13, 0, 0, 0, 0, time.UTC)})
}

func tambahPorto() {
	Input := bufio.NewReader(os.Stdin)

	var add Project

	fmt.Println()
	fmt.Print("Nama Proyek: ")
	name, _ := Input.ReadString('\n')
	add.name = strings.TrimSpace(name)

	fmt.Print("Kategori: ")
	category, _ := Input.ReadString('\n')
	add.category = strings.TrimSpace(category)

	fmt.Print("Kesulitan (1=Mudah, 2=Sedang, 3=Sulit): ")
	inputdiff, _ := Input.ReadString('\n')
	inputdiff = strings.TrimSpace(inputdiff)

	diffInt, err := strconv.Atoi(inputdiff)
	if err != nil || diffInt < 1 || diffInt > 3 {
		fmt.Println("Input kesulitan tidak valid! Harus angka 1–3.")
		return
	}
	add.difficulty = diffInt

	add.date = time.Now()

	Porto = append(Porto, add)
	fmt.Printf("Portofolio Berhasil Ditambahkan\n\n")
}

func ubahPorto() {
	Input := bufio.NewReader(os.Stdin)

	var nomor int

	showPorto()
	if len(Porto) == 0 {
		return
	}
	fmt.Print("Pilih nomor proyek yang ingin diedit: ")
	fmt.Scanln(&nomor)

	if nomor < 1 || nomor > len(Porto) {
		fmt.Println("Nomor proyek tidak valid.")
		return
	}

	ubah := &Porto[nomor-1]

	fmt.Printf("\nUBAH DATA\n\n")
	fmt.Println("Nama Projek Sekarang : ", ubah.name)
	fmt.Print("Nama Projek Yang Baru  (Kosongkan Jika Tidak Ingin Mengubah Judul) : ")
	name, _ := Input.ReadString('\n')
	name = strings.TrimSpace(name)
	if name != "" {
		ubah.name = name
	}

	fmt.Println("Kategori Projek Sekarang : ", ubah.category)
	fmt.Print("Kategori Projek Yang Baru  (Kosongkan Jika Tidak Ingin Mengubah Kategori) : ")
	category, _ := Input.ReadString('\n')
	category = strings.TrimSpace(category)
	if category != "" {
		ubah.category = category
	}

	fmt.Println("Difficulty Projek Sekarang : ", ubah.difficulty)
	fmt.Print("Difficulty Projek Yang Baru  (Kosongkan Jika Tidak Ingin Mengubah Difficulty) : ")
	inputdiff, _ := Input.ReadString('\n')
	inputdiff = strings.TrimSpace(inputdiff)
	if inputdiff != "" {
		diffInt, err := strconv.Atoi(inputdiff)
		if err != nil || diffInt < 1 || diffInt > 3 {
			fmt.Println("Input kesulitan tidak valid! Harus angka 1–3.")
			return
		}
		ubah.difficulty = diffInt
	}

	fmt.Printf("Portofolio Berhasil Diperbarui\n\n")
}

func hapusPorto()  {
	var nomor int
	
	if len(Porto) == 0{
		fmt.Println()
		fmt.Println("Belum ada Portofolio")
		return
	}
	showPorto()
	fmt.Print("Masukan Nomor Porto yang ingin dihapus: ")
	fmt.Scan(&nomor)
	nomor--

	if nomor < 0 || nomor > len(Porto)-1 {
		fmt.Println()
		fmt.Println("Input tidak valid!")
		return
	}else{
		var confirm string
		fmt.Print("Konfirmasi untuk menghapus Porto y/n: ")
		fmt.Scan(&confirm)
		if confirm == "y" {
			Porto = append(Porto[:nomor], Porto[nomor+1:]...)
            fmt.Println("\nPortofolio Telah di Hapus")
		}else if confirm == "n" {
			return
		}else{
			fmt.Println("Action tidak sesuai")
		}
	}
}

func showPorto() {
	if len(Porto) == 0 {
		fmt.Printf("\n")
		fmt.Println("Nama: -")
		fmt.Println("Kategori: -")
		fmt.Println("Difficulty: -")
		fmt.Println("Tanggal: -")
		return
	}
	for i, s := range Porto {
		fmt.Printf("\n")
		fmt.Printf("%d.Nama Portofolio: %s\n", i+1, s.name)
		fmt.Println("   Kategori       :",s.category)
		fmt.Println("   Kesulian       :",s.difficulty)
		fmt.Println("   Tanggal        :",s.date.Format("02-01-2006"))
	}
}

func cariPorto()  {
	var pilih int
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Println("Pilihan: ")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search ")
	fmt.Print("Masukan Pilihan: ")
	fmt.Scan(&pilih)
	if pilih < 1 || pilih > 2{
		fmt.Println("Pilihan tidak valid")
		return
	}
	fmt.Print("Masukkan nama Project yang ingin dicari: ")
	sJudul, _ := reader.ReadString('\n')
	sJudul = strings.TrimSpace(strings.ToLower(sJudul))

	switch pilih {
	case 1:
		cari := sequentialSearchPorto(sJudul)
		sque := Porto[cari]

		if cari == -1{
			fmt.Println()
			fmt.Println("Data Tidak Ditemukan")
			fmt.Println("Nama: -")
			fmt.Println("Kategori: -")
			fmt.Println("Difficulty: -")
			fmt.Println("Tanggal: -")
		}else{
			fmt.Println()
			fmt.Println("Data Ditemukan")
			fmt.Println("Nama:", sque.name)
			fmt.Println("Kategori:", sque.category)
			fmt.Println("Difficulty:", sque.difficulty)
			fmt.Println("Tanggal:",sque.date.Format("02-01-2006"))
		}
	case 2:
		cari := binarySearchPorto(sJudul)
		sque := Porto[cari]

		if cari == -1{
			fmt.Println()
			fmt.Println("Data Tidak Ditemukan")
			fmt.Println("Nama: -")
			fmt.Println("Kategori: -")
			fmt.Println("Difficulty: -")
			fmt.Println("Tanggal: -")
		}else{
			fmt.Println()
			fmt.Println("Data Ditemukan:")
			fmt.Println("Nama:", sque.name)
			fmt.Println("Kategori:", sque.category)
			fmt.Println("Difficulty:", sque.difficulty)
			fmt.Println("Tanggal:",sque.date.Format("02-01-2006"))
		}
	}
}

func sequentialSearchPorto(sJudul string) int {
	for i, cari := range Porto {
		if strings.ToLower(cari.name) == strings.ToLower(sJudul){
			return i
		}
	}
	return -1
}

func binarySearchPorto(sJudul string) int  {
	SelectionSort()
    start := 0
    end := len(Porto) - 1
	cari := strings.ToLower(sJudul)
    for start <= end {
        mid := (start + end) / 2
		data := strings.ToLower(Porto[mid].name)
		if data== cari {
            return mid
        } else if data < cari {
            start = mid + 1
        } else if data > cari {
            end = mid - 1
        }
    }
    return -1
}

func urutanPorto()  {

	var opsi int
	fmt.Println()
	fmt.Println("Pilihan Sorting Portofolio")
	fmt.Println("1. Insertion Sort(mengurutkan kesulitan 1-3)")
	fmt.Println("2. Selection Sort(mengurutkan judul a-z)")
	fmt.Print("Masukan pilihan: ")
	fmt.Scan(&opsi)
	
	switch opsi {
	case 1:
		InsertionSort()
	case 2:
		SelectionSort()
	}

}

func InsertionSort()  {
	for i := 1; i < len(Porto); i++ {
				j := i - 1
				temp := Porto[i]
				for j >= 0 && Porto[j].difficulty > temp.difficulty {
					Porto[j+1] = Porto[j]
					j--
				}
				Porto[j+1] = temp
		}
		for i, sort := range Porto {
			fmt.Printf("\n")
			fmt.Printf("%d. Nama Portofolio: %s\n", i+1, sort.name)
			fmt.Println("   Kategori       :",sort.category)
			fmt.Println("   Kesulian       :",sort.difficulty)
			fmt.Println("   Tanggal        :",sort.date.Format("02-01-2006"))

		}
}

func SelectionSort()  {
	for i := 0; i < len(Porto)-1; i++ {
		    minIndex := i
		    for j := i+1; j < len(Porto); j++ {
		        if strings.ToLower(Porto[j].name) < strings.ToLower(Porto[minIndex].name) {
		            minIndex = j
		        }
		    }
			
		    if minIndex != i {
		        Porto[i], Porto[minIndex] = Porto[minIndex], Porto[i]
		    }
		}
		for i, sort := range Porto {
			fmt.Printf("\n")
			fmt.Printf("%d. Nama Portofolio: %s\n", i+1, sort.name)
			fmt.Println("   Kategori       :",sort.category)
			fmt.Println("   Kesulian       :",sort.difficulty)
			fmt.Println("   Tanggal        :",sort.date.Format("02-01-2006"))

		}
}

func statPorto()  {
	categoryCount := make(map[string]int)

    for _, project := range Porto {
        categoryCount[project.category]++
    }
	fmt.Println()
	fmt.Println("Portofolio per Kategori")
	for i, categori := range categoryCount{
		fmt.Printf("- %s: %d proyek\n", i, categori)
	}
    
}