package main

import (
	"fmt"
)

// --- Struktur Data Global ---

const maxProject int = 1000
const maxKontribusi int = 20000
const maxPengguna int = 10000

var (
	countProject     int = 0
	projects         [maxProject]Projek
	nextProjectID    int = 1
	countKontribusi  int = 0
	contributions    [maxKontribusi]Kontribusi
	nextKontribusiID int = 1
	countPengguna    int = 0
	users            [maxPengguna]Pengguna
	nextPenggunaID   int = 1
)

type Projek struct {
	ID         int
	Nama       string
	Target     float64
	Current    float64
	JmlDonatur int
	OwnerID    int
	Category   string
}

type Kontribusi struct {
	ID         int
	ProjectID  int
	PenggunaID int
	Jumlah     float64
}

type Pengguna struct {
	ID           int
	TipePengguna string
	Password     string
	Nama         string
}

// --- Fungsi Bantu (Helper Functions) ---

func clearScreen() {
	// Karena tidak boleh import "os", kita tidak bisa membersihkan layar
	// secara harfiah. Sebagai gantinya, kita cetak banyak baris baru.
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n<<< Layar dibersihkan (simulasi) >>>\n\n")
}

func pauseExecution() {
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	var dummy string
	fmt.Scanln(&dummy)
}

// --- Fungsi Aplikasi: Manajemen Akun ---

func signUp() {
	if countPengguna >= maxPengguna {
		fmt.Println("Kapasitas pengguna penuh!")
		return
	}

	var newNama string
	var newPassword string
	var newType string
	var isTypeValid bool = false

	fmt.Println("\n--- Pendaftaran Pengguna Baru ---")
	fmt.Printf("ID Pengguna Anda (otomatis): %d\n", nextPenggunaID)

	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&newNama)

	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newPassword)

	var attempt int = 0
	for !isTypeValid && attempt < 3 { // Batasi percobaan input tipe pengguna
		fmt.Print("Daftar sebagai (owner/user): ")
		fmt.Scanln(&newType)
		if newType == "owner" || newType == "user" {
			isTypeValid = true
		} else {
			fmt.Println("Tipe pengguna tidak valid. Mohon masukkan 'owner' atau 'user'.")
			attempt++
		}
	}

	if !isTypeValid {
		fmt.Println("Gagal mendaftar: Tipe pengguna tidak valid setelah beberapa percobaan.")
		return
	}

	users[countPengguna] = Pengguna{
		ID:           nextPenggunaID,
		TipePengguna: newType,
		Password:     newPassword,
		Nama:         newNama,
	}
	countPengguna++
	nextPenggunaID++

	fmt.Println("Pendaftaran berhasil! Silakan Login dengan ID dan Password Anda.")
	fmt.Printf("ID Anda adalah: %d\n", users[countPengguna-1].ID)
}

func findPenggunaByID(idToFind int) (Pengguna, bool) {
	var found bool = false
	var foundUser Pengguna
	var i int = 0

	for i < countPengguna && !found {
		if users[i].ID == idToFind {
			foundUser = users[i]
			found = true
		}
		i++
	}
	return foundUser, found
}

func loginPengguna(id int, password string) (Pengguna, bool) {
	user, found := findPenggunaByID(id)

	if !found {
		var zeroPengguna Pengguna
		var resultBool bool = false
		return zeroPengguna, resultBool
	}

	var resultBool bool = false
	if user.Password == password {
		resultBool = true
	}
	return user, resultBool
}

// --- Fungsi Aplikasi: Manajemen Proyek ---

func findProjekByID(idToFind int) (Projek, bool) {
	var found bool = false
	var foundProjek Projek
	var i int = 0

	for i < countProject && !found {
		if projects[i].ID == idToFind {
			foundProjek = projects[i]
			found = true
		}
		i++
	}
	return foundProjek, found
}

func createProject(ownerID int) {
	if countProject >= maxProject {
		fmt.Println("Kapasitas proyek penuh!")
		return
	}
	var projectName string
	var projectTarget float64
	var projectCategory string

	fmt.Println("\n--- Buat Proyek Baru ---")
	fmt.Print("Masukkan Nama Proyek: ")
	fmt.Scanln(&projectName)

	fmt.Print("Masukkan Kategori Proyek (misal: Teknologi, Seni, Sosial): ")
	fmt.Scanln(&projectCategory)

	fmt.Print("Masukkan Target Dana (contoh: 100000.00): ")
	fmt.Scanln(&projectTarget)

	if projectTarget <= 0 {
		fmt.Println("Target dana harus lebih dari 0. Proyek gagal dibuat.")
		return
	}
	newProjek := Projek{
		ID:         nextProjectID,
		Nama:       projectName,
		Target:     projectTarget,
		Current:    0.0,
		JmlDonatur: 0,
		OwnerID:    ownerID,
		Category:   projectCategory,
	}

	projects[countProject] = newProjek
	countProject++
	nextProjectID++
	fmt.Println("Proyek berhasil dibuat!")
	fmt.Printf("ID Proyek: %d, Nama: '%s', Kategori: '%s', Target: %.2f\n", newProjek.ID, newProjek.Nama, newProjek.Category, newProjek.Target)
}

func editProject(ownerID int) {
	fmt.Println("\n--- Ubah Proyek Anda ---")
	viewMyProjects(ownerID)
	if countProject == 0 {
		fmt.Println("Tidak ada proyek untuk diubah.")
		return
	}

	var projectIDToEdit int
	fmt.Print("Masukkan ID Proyek yang ingin diubah (0 untuk batal): ")
	fmt.Scanln(&projectIDToEdit)

	if projectIDToEdit == 0 {
		fmt.Println("Pengubahan dibatalkan.")
		return
	}

	currentProject, projectFound := findProjekByID(projectIDToEdit)

	if !projectFound {
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
		return
	}

	if currentProject.OwnerID != ownerID {
		fmt.Println("Anda tidak memiliki izin untuk mengubah proyek ini. Anda bukan pemiliknya.")
		return
	}

	fmt.Printf("Mengubah proyek '%s' (ID: %d).\n", currentProject.Nama, currentProject.ID)

	var newNama, newKategori string
	var newTarget float64

	fmt.Printf("Nama Proyek baru (satu kata, '-' jika tidak ubah, sebelumnya: %s): ", currentProject.Nama)
	fmt.Scanln(&newNama)
	fmt.Printf("Kategori Proyek baru (satu kata, '-' jika tidak ubah, sebelumnya: %s): ", currentProject.Category)
	fmt.Scanln(&newKategori)
	fmt.Printf("Target Dana baru (angka, 0 jika tidak ubah, sebelumnya: %.2f): ", currentProject.Target)
	fmt.Scanln(&newTarget)

	var projectIdx int = -1
	var i int = 0
	for i < countProject && projectIdx == -1 {
		if projects[i].ID == projectIDToEdit {
			projectIdx = i
		}
		i++
	}

	if projectIdx != -1 {
		if newNama != "-" && len(newNama) > 0 {
			projects[projectIdx].Nama = newNama
		}
		if newKategori != "-" && len(newKategori) > 0 {
			projects[projectIdx].Category = newKategori
		}

		if newTarget < 0 {
			fmt.Println("Target dana baru tidak valid (tidak boleh negatif). Target tidak diubah.")
		} else if newTarget > 0 {
			projects[projectIdx].Target = newTarget
		}
		fmt.Println("Proyek berhasil diubah.")
	} else {
		fmt.Println("Internal Error: Proyek tidak ditemukan di array untuk diubah.")
	}
}

func deleteProject(ownerID int) {
	fmt.Println("\n--- Hapus Proyek Anda ---")
	viewMyProjects(ownerID)
	if countProject == 0 {
		fmt.Println("Tidak ada proyek untuk dihapus.")
		return
	}

	var projectIDToDelete int
	fmt.Print("Masukkan ID Proyek yang ingin dihapus (0 untuk batal): ")
	fmt.Scanln(&projectIDToDelete)

	if projectIDToDelete == 0 {
		fmt.Println("Penghapusan dibatalkan.")
		return
	}

	currentProject, projectFound := findProjekByID(projectIDToDelete)

	if !projectFound {
		fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
		return
	}

	if currentProject.OwnerID != ownerID {
		fmt.Println("Anda tidak memiliki izin untuk menghapus proyek ini. Anda bukan pemiliknya.")
		return
	}

	var confirm string
	fmt.Printf("Anda yakin ingin menghapus proyek '%s' (ID: %d)? (y/n): ", currentProject.Nama, currentProject.ID)
	fmt.Scanln(&confirm)

	if confirm == "y" {
		var namaProyekDihapus string = currentProject.Nama
		var projectIdx int = -1
		var i int = 0
		for i < countProject && projectIdx == -1 {
			if projects[i].ID == projectIDToDelete {
				projectIdx = i
			}
			i++
		}

		if projectIdx != -1 {
			var j int = projectIdx
			for j < countProject-1 {
				projects[j] = projects[j+1]
				j++
			}
			countProject--
			// Membersihkan elemen terakhir array yang mungkin tersisa
			if countProject >= 0 {
				var zeroProjek Projek // Membuat struct kosong
				projects[countProject] = zeroProjek
			}
			fmt.Printf("Proyek '%s' berhasil dihapus.\n", namaProyekDihapus)
		} else {
			fmt.Println("Internal Error: Proyek tidak ditemukan di array untuk dihapus.")
		}
	} else {
		fmt.Println("Penghapusan proyek dibatalkan.")
	}
}

func contributeToProject(userID int) {
	fmt.Println("\n--- Berkontribusi ke Proyek ---")
	if countProject == 0 {
		fmt.Println("Belum ada proyek yang bisa didanai.")
		return
	}
	viewAllProjects()

	var projectIDToDonate int
	fmt.Print("Masukkan ID Proyek yang ingin didanai: ")
	fmt.Scanln(&projectIDToDonate)

	targetProjek, projekFound := findProjekByID(projectIDToDonate)

	if !projekFound {
		fmt.Println("Gagal berkontribusi: Proyek dengan ID tersebut tidak ditemukan.")
		return
	}

	var amount float64
	fmt.Printf("Anda akan berkontribusi untuk proyek: %s\n", targetProjek.Nama)
	fmt.Print("Masukkan jumlah kontribusi: ")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Jumlah kontribusi harus lebih dari 0.")
		return
	}
	if countKontribusi >= maxKontribusi {
		fmt.Println("Kapasitas data kontribusi penuh.")
		return
	}

	var projekIndex int = -1
	var i int = 0
	for i < countProject && projekIndex == -1 {
		if projects[i].ID == projectIDToDonate {
			projekIndex = i
		}
		i++
	}

	if projekIndex != -1 {
		projects[projekIndex].Current += amount
		projects[projekIndex].JmlDonatur++

		newKontribusi := Kontribusi{
			ID:         nextKontribusiID,
			ProjectID:  projectIDToDonate,
			PenggunaID: userID,
			Jumlah:     amount,
		}

		contributions[countKontribusi] = newKontribusi
		countKontribusi++
		nextKontribusiID++

		fmt.Println("Terima kasih! Kontribusi Anda telah dicatat.")
		fmt.Printf("Dana terkumpul untuk proyek '%s' sekarang: %.2f dari %.2f\n", projects[projekIndex].Nama, projects[projekIndex].Current, projects[projekIndex].Target)
		if projects[projekIndex].Current >= projects[projekIndex].Target {
			fmt.Printf("🎉 Selamat! Proyek '%s' telah mencapai target pendanaan!\n", projects[projekIndex].Nama)
		}
	} else {
		fmt.Println("Internal Error: Proyek tidak ditemukan di array untuk kontribusi.")
	}
}

// --- Fungsi Aplikasi: Tampilan Proyek & Kontribusi ---

func printProjectHeader() {
	fmt.Printf("%-5s %-18s %-12s %-12s %-10s %-12s %-12s\n", "ID", "Nama Proyek", "Target", "Terkumpul", "Donatur", "Owner Name", "Kategori")
	fmt.Println("-----------------------------------------------------------------------------------------------")
}

func printProjectDetail(p Projek) {
	owner, _ := findPenggunaByID(p.OwnerID)
	ownerName := owner.Nama
	if len(ownerName) == 0 {
		ownerName = "N/A"
	}
	fmt.Printf("%-5d %-18s %-12.2f %-12.2f %-10d %-12s %-12s\n",
		p.ID, p.Nama, p.Target, p.Current, p.JmlDonatur, ownerName, p.Category)
}

func viewAllProjects() {
	fmt.Println("\n--- Daftar Semua Proyek ---")
	if countProject == 0 {
		fmt.Println("Belum ada proyek yang terdaftar.")
		return
	}
	printProjectHeader()
	var i int = 0
	for i < countProject {
		p := projects[i]
		ownerUser, found := findPenggunaByID(p.OwnerID)
		var ownerName string
		if found {
			ownerName = ownerUser.Nama
		} else {
			ownerName = "Owner Tidak Ditemukan"
		}
		fmt.Printf("%-5d %-20s %-15.2f %-15.2f %-15d %-20s %-20s\n",
			p.ID, p.Nama, p.Target, p.Current, p.JmlDonatur, ownerName, p.Category)
		i++
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	if countProject > 10 { // Jika ada batasan tampilan, seperti 10 proyek pertama
		fmt.Printf("Menampilkan %d dari total %d proyek. (Ada lebih banyak proyek yang tidak ditampilkan).\n", 10, countProject)
	}
}

func viewMyProjects(ownerID int) {
	fmt.Println("\n--- Proyek yang Anda Buat ---")
	var foundProjects bool = false
	printProjectHeader()
	var i int = 0
	for i < countProject {
		p := projects[i]
		if p.OwnerID == ownerID {
			foundProjects = true
			ownerUser, found := findPenggunaByID(p.OwnerID)
			var ownerName string
			if found {
				ownerName = ownerUser.Nama
			} else {
				ownerName = "Owner Tidak Ditemukan"
			}
			fmt.Printf("%-5d %-20s %-15.2f %-15.2f %-15d %-20s %-20s\n",
				p.ID, p.Nama, p.Target, p.Current, p.JmlDonatur, ownerName, p.Category)
		}
		i++
	}
	fmt.Println("--------------------------------------------------------------------------------------------")
	if !foundProjects {
		fmt.Println("Anda belum memiliki proyek yang terdaftar.")
	}
}

func viewMyContributions(userID int) {
	fmt.Println("\n--- Kontribusi Saya ---")
	var foundAny bool = false
	var headerPrinted bool = false

	var i int = 0
	for i < countKontribusi {
		k := contributions[i]
		if k.PenggunaID == userID {
			if !headerPrinted {
				fmt.Printf("%-5s %-10s %-20s %-15s\n", "ID", "ID Proyek", "Nama Proyek", "Jumlah")
				fmt.Println("----------------------------------------------------------------")
				headerPrinted = true
			}
			foundAny = true
			var namaProjek string = "N/A"
			var projectFoundForContrib bool = false
			var idxP int = 0
			for idxP < countProject && !projectFoundForContrib {
				if projects[idxP].ID == k.ProjectID {
					namaProjek = projects[idxP].Nama
					projectFoundForContrib = true
				}
				idxP++
			}
			fmt.Printf("%-5d %-10d %-20s %-15.2f\n", k.ID, k.ProjectID, namaProjek, k.Jumlah)
		}
		i++
	}
	if headerPrinted {
		fmt.Println("----------------------------------------------------------------")
	}
	if !foundAny {
		fmt.Println("Anda belum melakukan kontribusi.")
	}
}

// --- Fungsi Aplikasi: Pencarian Proyek ---

func sequentialSearchProject(query string, searchType string) {
	fmt.Printf("\n--- Hasil Pencarian Sequential untuk '%s' berdasarkan '%s' ---\n", query, searchType)

	var foundAny bool = false
	var headerPrinted bool = false

	var i int = 0
	for i < countProject {
		p := projects[i]
		var isMatch bool = false

		if searchType == "nama" {
			if p.Nama == query {
				isMatch = true
			}
		} else if searchType == "kategori" {
			if p.Category == query {
				isMatch = true
			}
		} else {
			fmt.Println("Tipe pencarian tidak valid. Gunakan 'nama' atau 'kategori'.")
			return
		}

		if isMatch {
			if !headerPrinted {
				printProjectHeader()
				headerPrinted = true
			}
			foundAny = true

			ownerUser, found := findPenggunaByID(p.OwnerID)
			var ownerName string
			if found {
				ownerName = ownerUser.Nama
			} else {
				ownerName = "Owner Tidak Ditemukan"
			}
			fmt.Printf("%-5d %-20s %-15.2f %-15.2f %-15d %-20s %-20s\n",
				p.ID, p.Nama, p.Target, p.Current, p.JmlDonatur, ownerName, p.Category)
		}
		i++
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	if !foundAny {
		fmt.Println("Tidak ada proyek yang cocok dengan kriteria pencarian.")
	}
}

func binarySearchProject(query string, searchType string) {
	fmt.Printf("\n--- Hasil Pencarian Binary untuk '%s' berdasarkan '%s' ---\n", query, searchType)

	if countProject == 0 {
		fmt.Println("Tidak ada proyek untuk dicari.")
		return
	}

	fmt.Println("CATATAN: Pencarian ini mengasumsikan daftar proyek SUDAH TERURUT berdasarkan '" + searchType + "'.")
	fmt.Println("------------------------------------------------------------------------------------------------------------")

	selectionSortProjek(searchType) // Panggil pengurutan sebelum binary search
	fmt.Println("Data proyek telah diurutkan sebelum pencarian biner.")

	var low int = 0
	var high int = countProject - 1
	var mid int
	var foundProjek Projek
	var foundMatch bool = false

	for low <= high && !foundMatch {
		mid = low + (high-low)/2
		p := projects[mid]
		var compareValue string

		if searchType == "nama" {
			compareValue = p.Nama
		} else if searchType == "kategori" {
			compareValue = p.Category
		} else {
			fmt.Println("Tipe pencarian tidak valid. Gunakan 'nama' atau 'kategori'.")
			return
		}

		if compareValue == query {
			foundProjek = p
			foundMatch = true
		} else if compareValue < query {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if foundMatch {
		printProjectHeader()
		printProjectDetail(foundProjek)
		fmt.Println("-----------------------------------------------------------------------------------------------")
	} else {
		fmt.Println("Proyek tidak ditemukan (atau data tidak terurut sesuai).")
	}
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}

// --- Pengurutan Proyek ---

func swapProjek(idx1 int, idx2 int) {
	tempProjek := projects[idx1]
	projects[idx1] = projects[idx2]
	projects[idx2] = tempProjek
}

func selectionSortProjek(sortType string) { // Variabel 'ascending' dihapus
	n := countProject

	var i int = 0
	for i < n-1 {
		var minIndex int = i

		var j int = i + 1
		for j < n {
			var isLess bool = false

			if sortType == "nama" {
				if projects[j].Nama < projects[minIndex].Nama {
					isLess = true
				}
			} else if sortType == "kategori" {
				if projects[j].Category < projects[minIndex].Category {
					isLess = true
				}
			} else {
				fmt.Println("Tipe pengurutan tidak valid. Gunakan 'nama' atau 'kategori'.")
				return
			}

			if isLess {
				minIndex = j
			}
			j++
		}

		if minIndex != i {
			swapProjek(i, minIndex)
		}
		i++
	}
	fmt.Printf("Proyek berhasil diurutkan berdasarkan %s.\n", sortType)
}

func insertionSortProjects(sortBy string, ascending bool) {
	if countProject <= 1 {
		fmt.Println("Tidak cukup data untuk diurutkan atau sudah terurut.")
		viewAllProjects()
		return
	}

	var orderStr string
	if ascending {
		orderStr = "Ascending"
	} else {
		orderStr = "Descending"
	}
	fmt.Printf("Mengurutkan proyek berdasarkan '%s' (%s) menggunakan Insertion Sort...\n", sortBy, orderStr)

	var i, j int
	i = 1
	for i < countProject {
		key := projects[i]
		j = i - 1

		var keepMoving bool = true
		for j >= 0 && keepMoving {
			var valKey, valJ float64
			pJ := projects[j]

			if sortBy == "dana" {
				valKey = key.Current
				valJ = pJ.Current
			} else if sortBy == "donatur" {
				valKey = float64(key.JmlDonatur)
				valJ = float64(pJ.JmlDonatur)
			} else {
				fmt.Println("Tipe pengurutan tidak valid. Gunakan 'dana' atau 'donatur'.")
				keepMoving = false
			}

			if keepMoving {
				var shouldMove bool = false
				if ascending {
					if valJ > valKey {
						shouldMove = true
					}
				} else {
					if valJ < valKey {
						shouldMove = true
					}
				}

				if shouldMove {
					projects[j+1] = projects[j]
					j--
				} else {
					keepMoving = false
				}
			}
		}
		projects[j+1] = key
		i++
	}
	fmt.Println("Pengurutan selesai.")
	viewAllProjects()
}

// --- Tampilkan Proyek Tercapai Target ---

func viewFundedProjects() {
	fmt.Println("\n--- Daftar Proyek yang Telah Mencapai Target Pendanaan ---")
	var foundAny bool = false
	var headerPrinted bool = false
	var i int = 0
	for i < countProject {
		p := projects[i]
		if p.Current >= p.Target {
			if !headerPrinted {
				printProjectHeader()
				headerPrinted = true
			}
			printProjectDetail(p)
			foundAny = true
		}
		i++
	}
	if headerPrinted {
		fmt.Println("-----------------------------------------------------------------------------------------------")
	}
	if !foundAny {
		fmt.Println("Belum ada proyek yang mencapai target pendanaan.")
	}
}

// --- Menu Utama Aplikasi ---

func showLoggedInMenu(user Pengguna) {
	var choice int
	var stayInMenu bool = true

	for stayInMenu {
		clearScreen()
		fmt.Printf("--- Selamat Datang, %s (%s)! ---\n", user.Nama, user.TipePengguna)

		viewAllProjects()

		fmt.Println("\n--- MENU ---")
		if user.TipePengguna == "owner" {
			fmt.Println("1. Buat Proyek Baru")
			fmt.Println("2. Lihat Proyek Saya")
			fmt.Println("3. Ubah Proyek Saya")
			fmt.Println("4. Hapus Proyek Saya")
			fmt.Println("5. Cari Proyek (Sequential)")
			fmt.Println("6. Cari Proyek (Binary)")
			fmt.Println("7. Urutkan Proyek (Selection Sort - Nama/Kategori)") // Update deskripsi
			fmt.Println("8. Urutkan Proyek (Insertion Sort - Dana/Donatur)") // Update deskripsi
			fmt.Println("9. Lihat Proyek Capai Target")
			fmt.Println("0. Logout")
		} else if user.TipePengguna == "user" {
			fmt.Println("1. Berkontribusi ke Proyek")
			fmt.Println("2. Lihat Kontribusi Saya")
			fmt.Println("3. Cari Proyek (Sequential)")
			fmt.Println("4. Cari Proyek (Binary)")
			fmt.Println("5. Urutkan Proyek (Selection Sort - Nama/Kategori)") // Update deskripsi
			fmt.Println("6. Urutkan Proyek (Insertion Sort - Dana/Donatur)") // Update deskripsi
			fmt.Println("7. Lihat Proyek Capai Target")
			fmt.Println("0. Logout")
		} else {
			fmt.Println("Tipe pengguna tidak dikenal. Silakan hubungi administrator.")
			stayInMenu = false
		}

		if stayInMenu {
			fmt.Print("Pilih Opsi: ")
			fmt.Scanln(&choice)

			// Logic untuk Owner
			if user.TipePengguna == "owner" {
				if choice == 1 {
					createProject(user.ID)
				} else if choice == 2 {
					viewMyProjects(user.ID)
				} else if choice == 3 {
					editProject(user.ID)
				} else if choice == 4 {
					deleteProject(user.ID)
				} else if choice == 5 {
					sequentialSearchProject(getSearchChoiceQuery(), getSearchChoiceType())
				} else if choice == 6 {
					binarySearchProject(getSearchChoiceQuery(), getSearchChoiceType())
				} else if choice == 7 {
					// Selection Sort hanya berdasarkan nama/kategori
					var sortType string
					fmt.Print("Urutkan berdasarkan (nama/kategori): ")
					fmt.Scanln(&sortType)
					selectionSortProjek(sortType)
				} else if choice == 8 {
					sortBy, ascending, valid := getSortChoice()
					if valid {
						insertionSortProjects(sortBy, ascending)
					}
				} else if choice == 9 {
					viewFundedProjects()
				} else if choice == 0 {
					fmt.Println("Logout berhasil. Kembali ke menu utama.")
					stayInMenu = false
				} else {
					fmt.Println("Pilihan menu tidak tersedia. Silakan coba lagi.")
				}
			} else if user.TipePengguna == "user" { // Logic untuk User
				if choice == 1 {
					contributeToProject(user.ID)
				} else if choice == 2 {
					viewMyContributions(user.ID)
				} else if choice == 3 {
					sequentialSearchProject(getSearchChoiceQuery(), getSearchChoiceType())
				} else if choice == 4 {
					binarySearchProject(getSearchChoiceQuery(), getSearchChoiceType())
				} else if choice == 5 {
					// Selection Sort hanya berdasarkan nama/kategori
					var sortType string
					fmt.Print("Urutkan berdasarkan (nama/kategori): ")
					fmt.Scanln(&sortType)
					selectionSortProjek(sortType)
				} else if choice == 6 {
					sortBy, ascending, valid := getSortChoice()
					if valid {
						insertionSortProjects(sortBy, ascending)
					}
				} else if choice == 7 {
					viewFundedProjects()
				} else if choice == 0 {
					fmt.Println("Logout berhasil. Kembali ke menu utama.")
					stayInMenu = false
				} else {
					fmt.Println("Pilihan menu tidak tersedia. Silakan coba lagi.")
				}
			}
			if stayInMenu {
				pauseExecution()
			}
		}
	}
}

// Helper untuk menu
func getSortChoice() (sortBy string, ascending bool, valid bool) {
	var sortTypeChoice int
	var orderChoice int
	fmt.Println("Urutkan berdasarkan: 1. Total Dana Terkumpul 2. Jumlah Donatur")
	fmt.Print("Pilih tipe urutan (1/2): ")
	fmt.Scanln(&sortTypeChoice)
	fmt.Println("Mode urutan: 1. Ascending (Menaik) 2. Descending (Menurun)")
	fmt.Print("Pilih mode urutan (1/2): ")
	fmt.Scanln(&orderChoice)

	valid = true
	if sortTypeChoice == 1 {
		sortBy = "dana"
	} else if sortTypeChoice == 2 {
		sortBy = "donatur"
	} else {
		fmt.Println("Pilihan tipe urutan tidak valid.")
		valid = false
	}

	if valid {
		if orderChoice == 1 {
			ascending = true
		} else if orderChoice == 2 {
			ascending = false
		} else {
			fmt.Println("Pilihan mode urutan tidak valid.")
			valid = false
		}
	}
	return
}

func getSearchChoiceQuery() string {
	var query string
	fmt.Print("Masukkan kata kunci (satu kata, case-sensitive): ")
	fmt.Scanln(&query)
	return query
}

func getSearchChoiceType() string {
	var searchType string
	fmt.Print("Cari Proyek berdasarkan (nama/kategori): ")
	fmt.Scanln(&searchType)
	return searchType
}

// --- Fungsi main (Alur Utama Program) ---
func main() {
	var loggedInUser Pengguna
	var isLoggedIn bool = false
	var stayInApp bool = true

	for stayInApp {
		clearScreen()
		fmt.Println("--- Aplikasi Crowdfunding Sederhana ---")
		fmt.Println("1. Login")
		fmt.Println("2. Sign Up (Daftar Baru)")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Print("Pilih Opsi: ")

		var initialChoice int
		fmt.Scanln(&initialChoice)

		if initialChoice == 1 {
			var inputID int
			var inputPassword string
			fmt.Println("\n--- Login Pengguna ---")
			fmt.Print("Masukkan ID Pengguna: ")
			fmt.Scanln(&inputID)
			fmt.Print("Masukkan Password (satu kata): ")
			fmt.Scanln(&inputPassword)

			loggedInUser, isLoggedIn = loginPengguna(inputID, inputPassword)
			if isLoggedIn {
				fmt.Println("Login berhasil!")
				showLoggedInMenu(loggedInUser)
			} else {
				fmt.Println("Login gagal: ID atau Password salah.")
			}

		} else if initialChoice == 2 {
			signUp()
		} else if initialChoice == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi. Sampai jumpa!")
			stayInApp = false
		} else {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		if !isLoggedIn && stayInApp && initialChoice != 0 {
			pauseExecution()
		} else if stayInApp && initialChoice == 1 && !isLoggedIn {
			pauseExecution()
		}
	}
}