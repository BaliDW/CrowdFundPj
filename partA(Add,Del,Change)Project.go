package main

import "fmt"

//registering new account to store in users array
func showLoggedInMenu(user pengguna) {
    var choice int
    var menuActive bool = true // Kontrol utama untuk loop menu

    for menuActive {
        fmt.Printf("\n--- Selamat Datang, %s (%s)! ---\n", user.nama, user.tipePengguna)

        if user.tipePengguna == "owner" {
            // Menu untuk Owner
            fmt.Println("1. Buat Proyek Baru")
            fmt.Println("2. Lihat Proyek Saya")
            fmt.Println("3. Lihat Semua Proyek")
            fmt.Println("0. Logout")
        } else if user.tipePengguna == "user" {
            // Menu untuk User Biasa
            fmt.Println("1. Lihat Semua Proyek")
            fmt.Println("2. Berkontribusi ke Proyek")
            fmt.Println("3. Lihat Kontribusi Saya")
            fmt.Println("0. Logout")
        } else {
            // Ini seharusnya tidak terjadi jika validasi tipe pengguna sudah benar saat sign up
            fmt.Println("Tipe pengguna tidak dikenal. Silakan hubungi administrator.")
            menuActive = false // Set 'menuActive' menjadi false untuk keluar dari loop
            // 'break' dihapus di sini, loop akan berakhir secara alami karena 'menuActive' sudah false.
        }

        // Jika menuActive sudah false (karena tipe pengguna tidak dikenal),
        // kita tidak perlu lagi meminta input.
        if !menuActive {
            continue // Langsung lanjut ke evaluasi kondisi loop 'for'
        }

        fmt.Print("Pilih: ")
        fmt.Scanln(&choice) // Gunakan Scanln untuk input pilihan

        switch choice {
        case 1:
            if user.tipePengguna == "owner" {
                fmt.Println("Anda memilih: Buat Proyek Baru (fungsi belum diimplementasi)")
                // Panggil fungsi untuk membuat proyek baru di sini
            } else { // user.tipePengguna == "user"
                fmt.Println("Anda memilih: Lihat Semua Proyek (fungsi belum diimplementasi)")
                // Panggil fungsi untuk melihat semua proyek di sini
            }
        case 2:
            if user.tipePengguna == "owner" {
                fmt.Println("Anda memilih: Lihat Proyek Saya (fungsi belum diimplementasi)")
                // Panggil fungsi untuk melihat proyek milik owner di sini
            } else { // user.tipePengguna == "user"
                fmt.Println("Anda memilih: Berkontribusi ke Proyek (fungsi belum diimplementasi)")
                // Panggil fungsi untuk berkontribusi di sini
            }
        case 3:
             if user.tipePengguna == "owner" {
                fmt.Println("Anda memilih: Lihat Semua Proyek (fungsi belum diimplementasi)")
                // Panggil fungsi untuk melihat semua proyek di sini
            } else { // user.tipePengguna == "user"
                fmt.Println("Anda memilih: Lihat Kontribusi Saya (fungsi belum diimplementasi)")
                // Panggil fungsi untuk melihat kontribusi user di sini
            }
        case 0:
            fmt.Println("Logout berhasil. Kembali ke menu utama.")
            menuActive = false // Set 'menuActive' menjadi false untuk keluar dari loop
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

func createProject(ownerID int) {
    if countProject >= maxProject {
        fmt.Println("Gagal membuat proyek: Kapasitas proyek penuh!")
        return
    }

    var projectName string
    var projectTarget float64

    fmt.Println("\n--- Buat Proyek Baru ---")
    fmt.Print("Masukkan Nama Proyek: ")
    fmt.Scanln(&projectName)

    fmt.Print("Masukkan Target Dana (contoh: 100000.00): ")
    fmt.Scanln(&projectTarget)

    if projectTarget <= 0 {
        fmt.Println("Target dana harus lebih dari 0. Proyek gagal dibuat.")
        return
    }

    var newProject projek
    newProject.ID = nextProjectID
    newProject.nama = projectName
    newProject.target = projectTarget
    newProject.current = 0.0 // Dana awal selalu 0
    newProject.jmlDonatur = 0
    newProject.ownerID = ownerID

    projects[countProject] = newProject
    countProject = countProject + 1
    nextProjectID = nextProjectID + 1

    fmt.Println("Proyek berhasil dibuat!")
    fmt.Printf("ID Proyek: %d, Nama: %s, Target: %.2f\n", newProject.ID, newProject.nama, newProject.target)
}
