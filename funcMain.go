package main

import "fmt"

// --- Bagian Deklarasi Global (Pastikan Ini Sama Persis di File Anda) ---
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
type arrKontribusi [maxKontribusi]kontribusi 
var contributions arrKontribusi 
var nextKontribusiID int = 1
type kontribusi struct {
	ID         int 
	projectID  int
	penggunaID int     
	jumlah     float64 
}

const maxPengguna int = 10000 
var countPengguna int = 0 
type arrPengguna [maxPengguna]pengguna
var users arrPengguna 

var nextPenggunaID int = 1
type pengguna struct {
	ID           int
	tipePengguna string
	password     string 
	nama         string
}
// --- Akhir Bagian Deklarasi Global ---


// --- Fungsi signUp (tidak berubah) ---
func signUp() {
	if countPengguna >= maxPengguna {
		fmt.Println("Gagal mendaftar: Kapasitas pengguna sedang penuh!")
		return
	}
	var newNama string
	var newPassword string 
	var newType string
	var isTypeValid bool = false
	var assignedID int = nextPenggunaID
	nextPenggunaID = nextPenggunaID + 1 
	fmt.Println("\n--- Pendaftaran Pengguna Baru ---")
	fmt.Printf("Anda akan terdaftar dengan ID: %d\n", assignedID)
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&newNama)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newPassword)
	for !isTypeValid {
		fmt.Print("Daftar sebagai (owner/user): ")
		fmt.Scanln(&newType)
		if newType == "owner" || newType == "user" {
			isTypeValid = true
		} else {
				fmt.Println("Tipe pengguna tidak valid. Mohon masukkan 'owner' atau 'user' (huruf kecil).")
		}
	}
	var newUser pengguna
	newUser.ID = assignedID
	newUser.tipePengguna = newType
	newUser.password = newPassword
	newUser.nama = newNama
	users[countPengguna] = newUser
	countPengguna = countPengguna + 1 
	fmt.Println("Pendaftaran berhasil! Silakan Login dengan ID dan Password Anda.")
	fmt.Printf("ID Anda adalah: %d\n", assignedID)
}

// --- Fungsi findPenggunaByID (tidak berubah) ---
func findPenggunaByID(idToFind int)(pengguna,bool){
	var found bool = false 
	var foundUser pengguna 
	var i int 
	for i = 0; i < countPengguna && !found; i = i + 1 { 
		if users[i].ID == idToFind { 
			foundUser = users[i]
			found = true
		}
	}
	return foundUser, found
}

// --- Fungsi loginPengguna (tidak berubah) ---
func loginPengguna(id int, password string) (pengguna, bool) { 
	var user pengguna
	var found bool
	user, found = findPenggunaByID(id)
	if !found {
		var zeroPengguna pengguna
		var resultBool bool = false
		return zeroPengguna, resultBool
	}
	if user.password == password { 
		var resultBool bool = true
		return user, resultBool
	} else {
		var zeroPengguna pengguna
		var resultBool bool = false
		return zeroPengguna, resultBool
	}
}

// --- Fungsi createProject (tidak berubah) ---
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
	newProject.current = 0.0
	newProject.jmlDonatur = 0
	newProject.ownerID = ownerID
	projects[countProject] = newProject
	countProject = countProject + 1
	nextProjectID = nextProjectID + 1
	fmt.Println("Proyek berhasil dibuat!")
	fmt.Printf("ID Proyek: %d, Nama: %s, Target: %.2f\n", newProject.ID, newProject.nama, newProject.target)
}

// --- Fungsi viewAllProjects (tidak berubah) ---
func viewAllProjects() {
	fmt.Println("\n--- Daftar Semua Proyek ---")
	if countProject == 0 {
		fmt.Println("Belum ada proyek yang terdaftar.")
		return
	}
	fmt.Printf("%-5s %-20s %-15s %-15s %-15s %-20s\n", "ID", "Nama Proyek", "Target Dana", "Dana Terkumpul", "Jml Donatur", "Nama Owner")
	fmt.Println("--------------------------------------------------------------------------------------------") 
	var i int = 0
	for i = 0; i < countProject; i = i + 1 { 
		p := projects[i]
		ownerUser, found := findPenggunaByID(p.ownerID)
		var ownerName string
		if found {
			ownerName = ownerUser.nama
		} else {
			ownerName = "Owner Tidak Ditemukan" 
		}
		fmt.Printf("%-5d %-20s %-15.2f %-15.2f %-15d %-20s\n",
			p.ID, p.nama, p.target, p.current, p.jmlDonatur, ownerName)
	}
	fmt.Println("--------------------------------------------------------------------------------------------") 
}

// --- Fungsi viewMyProjects (tidak berubah) ---
func viewMyProjects(ownerID int) {
	fmt.Println("\n--- Proyek yang Anda Buat ---")
	var foundProjects bool = false 
	fmt.Printf("%-5s %-20s %-15s %-15s %-15s %-20s\n", "ID", "Nama Proyek", "Target Dana", "Dana Terkumpul", "Jml Donatur", "Nama Owner")
	fmt.Println("--------------------------------------------------------------------------------------------") 
	var i int = 0
	for i = 0; i < countProject; i = i + 1 { 
		p := projects[i]
		if p.ownerID == ownerID {
			foundProjects = true 
			ownerUser, found := findPenggunaByID(p.ownerID)
			var ownerName string
			if found {
				ownerName = ownerUser.nama
			} else {
				ownerName = "Owner Tidak Ditemukan" 
			}
			fmt.Printf("%-5d %-20s %-15.2f %-15.2f %-15d %-20s\n",
				p.ID, p.nama, p.target, p.current, p.jmlDonatur, ownerName)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------------------") 
	if !foundProjects {
		fmt.Println("Anda belum memiliki proyek yang terdaftar.")
	}
}

// --- Fungsi showLoggedInMenu (Diperbaiki) ---
func showLoggedInMenu(user pengguna) {
    var choice int
    var menuActive bool = true 

    for menuActive {
        fmt.Printf("\n--- Selamat Datang, %s (%s)! ---\n", user.nama, user.tipePengguna)

        if user.tipePengguna == "owner" {
            // Menu untuk Owner (Opsi 'Lihat Semua Proyek' dihapus karena sudah otomatis tampil saat login)
            fmt.Println("1. Buat Proyek Baru")
            fmt.Println("2. Lihat Proyek Saya")
            // Opsi 3 (Lihat Semua Proyek) tidak ditampilkan lagi di sini
            fmt.Println("0. Logout")
        } else if user.tipePengguna == "user" {
            // Menu untuk User Biasa (Opsi 'Lihat Semua Proyek' dihapus, nomor opsi disesuaikan)
            fmt.Println("1. Berkontribusi ke Proyek")
            fmt.Println("2. Lihat Kontribusi Saya")
            // Opsi 3 (Lihat Semua Proyek) tidak ditampilkan lagi di sini
            fmt.Println("0. Logout")
        } else {
            fmt.Println("Tipe pengguna tidak dikenal. Silakan hubungi administrator.")
            menuActive = false 
        }

        if menuActive { 
            fmt.Print("Pilih: ")
            fmt.Scanln(&choice) 

            switch choice {
            case 1:
                if user.tipePengguna == "owner" {
                    createProject(user.ID) 
                } else { // user.tipePengguna == "user"
                    // Opsi 1 untuk user sekarang adalah 'Berkontribusi ke Proyek'
                    fmt.Println("Anda memilih: Berkontribusi ke Proyek (fungsi belum diimplementasi)")
                }
            case 2:
                if user.tipePengguna == "owner" {
                    viewMyProjects(user.ID) 
                } else { // user.tipePengguna == "user"
                    // Opsi 2 untuk user sekarang adalah 'Lihat Kontribusi Saya'
                    fmt.Println("Anda memilih: Lihat Kontribusi Saya (fungsi belum diimplementasi)")
                }
            case 0:
                fmt.Println("Logout berhasil. Kembali ke menu utama.")
                menuActive = false 
            default:
                fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
            }
        }
    }
}

// --- Fungsi main (tidak berubah) ---
func main() {
    var loggedInUser pengguna 
    var isLoggedIn bool = false 
    var isEND bool = false
    var initialChoice int

    for !isLoggedIn && !isEND {
        fmt.Println("\n--- Selamat Datang ---")
        fmt.Println("1. Login")
        fmt.Println("2. Sign Up (Daftar Baru)")
        fmt.Println("0. Keluar Program")
        fmt.Print("Pilih: ")
        fmt.Scanln(&initialChoice) 

        switch initialChoice {
        case 1: // Pilih Login
            fmt.Println("\nSilakan login dengan akun yang telah diregister.")
            var inputID int 
            var inputPassword string

            fmt.Print("Masukkan ID Pengguna: ")
            fmt.Scanln(&inputID) 

            fmt.Print("Masukkan Password: ")
            fmt.Scanln(&inputPassword) 

            loggedInUser, isLoggedIn = loginPengguna(inputID, inputPassword)

            if isLoggedIn {
            fmt.Printf("\nSelamat datang, %s!\n", loggedInUser.nama)
            fmt.Printf("Anda login sebagai: %s\n", loggedInUser.tipePengguna)
            
            viewAllProjects() // Panggil viewAllProjects di sini
            
            showLoggedInMenu(loggedInUser) 
            
            isLoggedIn = false 
            } else {
                fmt.Println("Login gagal: ID atau Password salah.")
            }

        case 2: // Pilih Sign Up
            signUp() 
            fmt.Println("\nPendaftaran berhasil! Silakan pilih '1. Login' dari menu untuk masuk.")

        case 0: // Pilih Keluar Program
            fmt.Println("Terima kasih. Program selesai.")
            isEND = true 
        default: 
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }

    if isLoggedIn {
         fmt.Println("\nAnda telah berhasil login. (Program akan berakhir atau lanjut ke fungsionalitas lain)")
    } else {
        fmt.Println("\nProgram selesai.") 
    }
}