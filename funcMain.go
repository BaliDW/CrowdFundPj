
package main

import "fmt"

const maxProject int = 1000
var countProject int = 0
type arrProject [maxProject]projek
var projects arrProject
var nextProjectID int = 1
type projek struct {
	ID         int
	nama       string
	target     float64
	current    float64
	jmlDonatur int
	ownerID    int
}

const maxKontribusi int = 20000
var countKontribusi int = 0
type contributions [maxKontribusi]kontribusi
var nextKontribusiID int = 1
type kontribusi struct {
	ID         int //IDtransaksi
	projectID  int
	penggunaID string
	jumlah     float64 //jumlah dana 
}

const maxPengguna int = 10000 // Batas maksimum pengguna
var countPengguna int = 0     // Jumlah pengguna yang saat ini valid
type arrPengguna [maxPengguna]pengguna
var users arrPengguna
var nextPenggunaID int = 1
type pengguna struct {
	ID           int
	tipePengguna string
	password     byte
	nama         string
}

func signUp() {
if countPengguna >= maxPengguna {
        fmt.Println("Gagal mendaftar: Kapasitas pengguna sedang penuh!")
        return // PENTING: Keluar dari fungsi jika kapasitas penuh
    }

    // Variabel lokal untuk menyimpan input
    var newNama string
    var newPassword byte // <--- Perbaikan: Password harus string
    var newType string
    var isTypeValid bool = false // Nama variabel lebih sesuai dengan isian tipe pengguna

    // ID otomatis dihitung dan diberitahukan, tidak diminta dari pengguna
    var assignedID int = nextPenggunaID
    nextPenggunaID = nextPenggunaID + 1 // <--- Perbaikan: Gunakan = untuk increment

    fmt.Println("\n--- Pendaftaran Pengguna Baru ---")
    fmt.Printf("Anda akan terdaftar dengan ID: %d\n", assignedID) // Menggunakan ID otomatis

    fmt.Print("Masukkan Nama Anda: ")
    fmt.Scanln(&newNama) // <--- Pastikan semua menggunakan Scanln

    fmt.Print("Masukkan Password: ")
    fmt.Scanln(&newPassword) // <--- Pastikan semua menggunakan Scanln, dan newPassword adalah string

    for !isTypeValid { // Loop validasi tipe pengguna
        fmt.Print("Daftar sebagai (owner/user): ")
        fmt.Scanln(&newType) // <--- Pastikan semua menggunakan Scanln
		fmt.Println()
        if newType == "owner" || newType == "user" {
            isTypeValid = true
        } else {
            fmt.Println("Tipe pengguna tidak valid. Mohon masukkan 'owner' atau 'user' (huruf kecil).")
        }
    }

    // Membuat struct pengguna baru
    var newUser pengguna

    newUser.ID = assignedID // Menggunakan ID yang otomatis digenerate
    newUser.tipePengguna = newType
    newUser.password = newPassword // <--- Perbaikan: Gunakan field yang benar untuk password
    newUser.nama = newNama

    // Menyimpan pengguna baru ke array global users
    users[countPengguna] = newUser 
    countPengguna++
    fmt.Println("Pendaftaran berhasil! Silakan Login dengan ID dan Password Anda.")
    fmt.Printf("ID Anda adalah: %d\n", assignedID) // Menggunakan ID otomatis
}

func findPenggunaByID(idToFind int)(pengguna,bool){
	var found bool = false 
    var foundUser pengguna 
	var i int 

	for i = 0; i < countPengguna && !found; i++ {
		if users[i].ID == idToFind {
       		foundUser = users[i]
			found = true
		}
	}
	return foundUser, found
}

func loginPengguna(id int, password byte) (pengguna, bool) {
    var user pengguna
    var found bool
    user, found = findPenggunaByID(id) // Langkah 1: Mencari pengguna berdasarkan ID

    if !found { // Jika pengguna tidak ditemukan
        var zeroPengguna pengguna
        var resultBool bool = false
        return zeroPengguna, resultBool // Mengembalikan struct kosong dan false
    }

    // Jika pengguna ditemukan, lanjutkan ke langkah 2: Membandingkan password
    if user.password == password { // Membandingkan password yang disimpan dengan input
        var resultBool bool = true
        return user, resultBool // Mengembalikan data pengguna dan true (berhasil login)
    } else { // Jika password tidak cocok
        var zeroPengguna pengguna
        var resultBool bool = false
        return zeroPengguna, resultBool // Mengembalikan struct kosong dan false
    }
}

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
                fmt.Println("Anda memilih: Buat Proyek Baru ")
                // Panggil fungsi untuk membuat proyek baru di sini
				createProject(user.ID)
            } else { // user.tipePengguna == "user"
                fmt.Println("Anda memilih: Lihat Semua Proyek ")
                // Panggil fungsi untuk melihat semua proyek di sini
				viewAllProjects()
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
                fmt.Println("Anda memilih: Lihat Semua Proyek")
                // Panggil fungsi untuk melihat semua proyek di sini
				viewAllProjects()
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

// Fungsi untuk membuat proyek baru (hanya untuk owner)
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

// Fungsi untuk melihat daftar semua proyek
func viewAllProjects() {
    fmt.Println("\n--- Daftar Semua Proyek ---")
    if countProject == 0 {
        fmt.Println("Belum ada proyek yang terdaftar.")
        return
    }

    fmt.Printf("%-5s %-20s %-15s %-15s %-15s %-10s\n", "ID", "Nama Proyek", "Target Dana", "Dana Terkumpul", "Jml Donatur", "Owner name")
    fmt.Println("--------------------------------------------------------------------------------")
    for i := 0; i < countProject; i++ {
        p := projects[i]
		ownerUser, found := findPenggunaByID(p.ownerID)
        var ownerName string
        if found {
            ownerName = ownerUser.nama
        } else {
            ownerName = "Owner Tidak Ditemukan" // Atau pesan lain jika ID Owner tidak valid
        }
        fmt.Printf("%-5d %-20s %-15.2f %-15.2f %-15d %-10s\n",
            p.ID, p.nama, p.target, p.current, p.jmlDonatur, ownerName)
    }
    fmt.Println("--------------------------------------------------------------------------------")
}

func main() {
	var loggedInUser pengguna // Ganti nama 'login' menjadi 'loggedInUser' agar lebih jelas
	var isLoggedIn bool = false // Inisialisasi status login
	var isEND bool = false      // Inisialisasi status selesai program
	var initialChoice int
	
	// Loop akan terus berjalan selama pengguna belum login DAN belum memilih untuk keluar
	for !isLoggedIn && !isEND {
		fmt.Println("\n--- Selamat Datang ---")
		fmt.Println("1. Login")
		fmt.Println("2. Sign Up (Daftar Baru)")
		fmt.Println("0. Keluar Program")
		fmt.Print("Pilih: ")
		fmt.Scanln(&initialChoice) // Pastikan menggunakan Scanln

		// Tidak perlu reset menuAwal di sini, karena loop condition yang mengontrol.

		switch initialChoice {
		case 1: // Pilih Login
			fmt.Println("\nSilakan login dengan akun yang telah diregister.")
			var inputID int      // Gunakan variabel sementara untuk input
			var inputPassword byte

			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scanln(&inputID) // Pastikan menggunakan Scanln
			
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&inputPassword) // Pastikan menggunakan Scanln
			
			// Panggil fungsi loginPengguna dan tangkap hasilnya ke loggedInUser dan isLoggedIn
			loggedInUser, isLoggedIn = loginPengguna(inputID, inputPassword)

			if isLoggedIn {
				fmt.Printf("\nSelamat datang, %s!\n", loggedInUser.nama)
				fmt.Printf("Anda login sebagai: %s\n", loggedInUser.tipePengguna)
				// Jika berhasil login, loop utama akan berakhir karena 'isLoggedIn' menjadi true.
				// Di sini Anda bisa melanjutkan ke menu utama setelah login.

				// Panggil menu khusus untuk pengguna yang sudah login (owner/user)
				showLoggedInMenu(loggedInUser) 
				
				// Setelah 'showLoggedInMenu' kembali (misal, pengguna memilih 'Logout'),
				// reset status login ke 'false' agar loop utama kembali menampilkan menu Login/Sign Up.
				isLoggedIn = false 
			} else {
				fmt.Println("Login gagal: ID atau Password salah.")
				// Loop akan berlanjut, pengguna bisa mencoba lagi atau memilih opsi lain.
			}

		case 2: // Pilih Sign Up
			signUp() // Panggil prosedur registrasi
			fmt.Println("\nPendaftaran berhasil! Silakan pilih '1. Login' dari menu untuk masuk.")
			// Loop akan berlanjut, kembali menampilkan menu utama.

		case 0: // Pilih Keluar Program
			fmt.Println("Terima kasih. Program selesai.")
			isEND = true // Set isEND menjadi true untuk keluar dari loop
		default: // Pilihan tidak valid
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}

	// Setelah loop selesai, program akan sampai di sini.
	if isLoggedIn {
		// Jika loop berakhir karena berhasil login, Anda bisa melanjutkan ke menu berikutnya di sini.
		fmt.Println("\nAnda telah berhasil login. (Lanjutkan ke menu user/owner)")
		// Contoh: Jika ada fungsi menu setelah login, panggil di sini
		// showUserMenu(loggedInUser)
	} else {
		// Jika loop berakhir karena pengguna memilih keluar (isEND = true)
		fmt.Println("\nProgram selesai.") 
	}
}

