package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// Konstanta untuk kapasitas maksimum struktur data.
const maxProject int = 1000
const maxKontribusi int = 20000
const maxPengguna int = 10000

// Konstanta global untuk lebar kolom tabel proyek.
const (
	idColWidth       = 5
	nameColWidth     = 28
	targetColWidth   = 12
	currentColWidth  = 12
	donaturColWidth  = 10
	ownerColWidth    = 15
	categoryColWidth = 15
	totalTableWidth  = idColWidth + nameColWidth + targetColWidth + currentColWidth + donaturColWidth + ownerColWidth + categoryColWidth + (7 - 1)
)

// Variabel global untuk menyimpan data aplikasi dan mengelola ID.
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

// Projek merepresentasikan proyek penggalangan dana.
type Projek struct {
	ID         int
	Nama       string
	Target     float64
	Current    float64
	JmlDonatur int
	OwnerID    int
	Category   string
}

// Kontribusi merepresentasikan satu kali kontribusi yang dilakukan ke sebuah proyek.
type Kontribusi struct {
	ID         int
	ProjectID  int
	PenggunaID int
	Jumlah     float64
}

// Pengguna merepresentasikan pengguna aplikasi penggalangan dana.
type Pengguna struct {
	ID           int
	TipePengguna string
	Password     string
	Nama         string
}

// CLS membersihkan layar terminal.
func CLS() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("CLS untuk ", runtime.GOOS, " tidak diimplementasikan")
		return
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// pauseExecution menghentikan eksekusi program dan menunggu input dari pengguna (menekan Enter).
func pauseExecution() {
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	var dummy string
	fmt.Scanln(&dummy)
}

// findProjectArrayIndex mencari indeks (posisi dalam array) dari sebuah proyek berdasarkan ID Proyeknya.
func findProjectArrayIndex(projectID int) int {
	idx := -1
	var i = 0
	for i < countProject && idx == -1 {
		if projects[i].ID == projectID {
			idx = i
		}
		i++
	}
	return idx
}

// signUp memungkinkan pengguna baru untuk mendaftar ke aplikasi dengan nama, kata sandi, dan tipe pengguna.
func signUp() {
	if countPengguna >= maxPengguna {
		fmt.Println("Kapasitas pengguna penuh!")
		return
	}
	var newNama, newPassword, newType string
	var isTypeValid bool = false

	fmt.Println("\n--- Pendaftaran Pengguna Baru ---")
	fmt.Printf("ID Pengguna Anda (otomatis): %d\n", nextPenggunaID)
	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&newNama)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&newPassword)

	var attempt int = 0
	for !isTypeValid && attempt < 3 {
		fmt.Print("Daftar sebagai (owner/user): ")
		fmt.Scanln(&newType)
		if newType == "owner" || newType == "user" {
			isTypeValid = true
		} else {
			fmt.Println("Tipe pengguna tidak valid. Mohon masukkan 'owner' atau 'user'.")
			attempt++
		}
	}

	if isTypeValid {
		users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: newType, Password: newPassword, Nama: newNama}
		currentID := nextPenggunaID
		countPengguna++
		nextPenggunaID++
		fmt.Println("Pendaftaran berhasil! Silakan Login dengan ID dan Password Anda.")
		fmt.Printf("ID Anda adalah: %d\n", currentID)
	} else {
		fmt.Println("Gagal mendaftar: Tipe pengguna tidak valid setelah beberapa percobaan.")
	}
}

// findPenggunaByID mencari pengguna berdasarkan ID-nya.
func findPenggunaByID(idToFind int) (Pengguna, bool) {
	var foundUser Pengguna
	var found bool = false
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

// loginPengguna mencoba melakukan proses login untuk pengguna dengan ID dan kata sandi yang diberikan.
func loginPengguna(id int, password string) (Pengguna, bool) {
	user, found := findPenggunaByID(id)
	var success bool = false
	if found {
		success = (user.Password == password)
	}
	return user, success
}

// findProjekByID mencari proyek berdasarkan ID-nya.
func findProjekByID(idToFind int) (Projek, bool) {
	var foundProjek Projek
	var found bool = false
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

// createProject memungkinkan pengguna dengan tipe "owner" untuk membuat proyek crowdfunding baru.
func createProject(ownerID int) {
	if countProject >= maxProject {
		fmt.Println("Kapasitas proyek penuh!")
		return
	}
	var projectName, projectCategory string
	var projectTarget float64

	fmt.Println("\n--- Buat Proyek Baru ---")
	fmt.Print("Masukkan Nama Proyek (gunakan '_' untuk spasi): ")
	fmt.Scanln(&projectName)
	fmt.Print("Masukkan Kategori Proyek (misal: Teknologi, Seni, Sosial): ")
	fmt.Scanln(&projectCategory)
	fmt.Print("Masukkan Target Dana (contoh: 100000.00): ")
	fmt.Scanln(&projectTarget)

	if projectTarget <= 0 {
		fmt.Println("Target dana harus lebih dari 0. Proyek gagal dibuat.")
		return
	}
	newProjek := Projek{ID: nextProjectID, Nama: projectName, Target: projectTarget, Current: 0.0, JmlDonatur: 0, OwnerID: ownerID, Category: projectCategory}
	projects[countProject] = newProjek
	countProject++
	nextProjectID++
	fmt.Println("Proyek berhasil dibuat!")
	fmt.Printf("ID Proyek: %d, Nama: '%s', Kategori: '%s', Target: %.2f\n", newProjek.ID, newProjek.Nama, newProjek.Category, newProjek.Target)
}

// editProject memungkinkan pengguna 'owner' untuk mengubah detail proyek yang mereka miliki.
func editProject(ownerID int) {
	fmt.Println("\n--- Ubah Proyek Anda ---")
	viewMyProjects(ownerID)

	hasProjects := false
	for i := 0; i < countProject; i++ {
		if projects[i].OwnerID == ownerID {
			hasProjects = true
			// Tidak perlu break, flag hasProjects sudah cukup
		}
	}
	if !hasProjects {
		// Pesan "Anda belum memiliki proyek" sudah ditangani oleh viewMyProjects
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
		fmt.Println("Anda bukan pemilik proyek ini.")
		return
	}

	projectIdx := findProjectArrayIndex(projectIDToEdit)
	if projectIdx == -1 {
		fmt.Println("Kesalahan internal: Proyek tidak ditemukan di array.")
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

	var changed bool = false
	if newNama != "-" && len(newNama) > 0 {
		projects[projectIdx].Nama = newNama
		changed = true
	}
	if newKategori != "-" && len(newKategori) > 0 {
		projects[projectIdx].Category = newKategori
		changed = true
	}
	if newTarget > 0 {
		projects[projectIdx].Target = newTarget
		changed = true
	} else if newTarget < 0 {
		fmt.Println("Target dana baru tidak valid (negatif). Target tidak diubah.")
	}

	if changed {
		fmt.Println("Proyek berhasil diubah.")
	} else {
		fmt.Println("Tidak ada perubahan yang dilakukan.")
	}
}

// deleteProject memungkinkan pengguna 'owner' untuk menghapus proyek yang mereka miliki.
func deleteProject(ownerID int) {
	fmt.Println("\n--- Hapus Proyek Anda ---")
	viewMyProjects(ownerID)

	hasProjects := false
	for i := 0; i < countProject; i++ {
		if projects[i].OwnerID == ownerID {
			hasProjects = true
			// Tidak perlu break
		}
	}
	if !hasProjects {
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
		fmt.Println("Anda bukan pemilik proyek ini.")
		return
	}

	var confirm string
	fmt.Printf("Anda yakin ingin menghapus proyek '%s' (ID: %d)? (y/n): ", currentProject.Nama, currentProject.ID)
	fmt.Scanln(&confirm)

	if confirm == "y" {
		projectIdx := findProjectArrayIndex(projectIDToDelete)
		if projectIdx != -1 {
			deletedProjectName := projects[projectIdx].Nama
			var j int = projectIdx
			for j < countProject-1 {
				projects[j] = projects[j+1]
				j++
			}
			countProject--
			var zeroProjek Projek
			if countProject >= 0 && countProject < maxProject {
				projects[countProject] = zeroProjek
			}
			fmt.Printf("Proyek '%s' berhasil dihapus.\n", deletedProjectName)
		} else {
			fmt.Println("Kesalahan Internal: Proyek tidak ditemukan di array untuk dihapus.")
		}
	} else {
		fmt.Println("Penghapusan dibatalkan oleh pengguna.")
	}
}

// contributeToProject memungkinkan pengguna dengan tipe "user" untuk memberikan kontribusi dana ke sebuah proyek.
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

	projectIdx := findProjectArrayIndex(projectIDToDonate)
	if projectIdx != -1 {
		projects[projectIdx].Current += amount
		projects[projectIdx].JmlDonatur++
		contributions[countKontribusi] = Kontribusi{ID: nextKontribusiID, ProjectID: projectIDToDonate, PenggunaID: userID, Jumlah: amount}
		countKontribusi++
		nextKontribusiID++
		fmt.Println("Terima kasih! Kontribusi Anda telah dicatat.")
		fmt.Printf("Dana terkumpul untuk proyek '%s' sekarang: %.2f dari %.2f\n", projects[projectIdx].Nama, projects[projectIdx].Current, projects[projectIdx].Target)
		if projects[projectIdx].Current >= projects[projectIdx].Target {
			fmt.Printf("🎉 Selamat! Proyek '%s' telah mencapai target pendanaan!\n", projects[projectIdx].Nama)
		}
	} else {
		fmt.Println("Kesalahan Internal: Proyek tidak ditemukan di array untuk kontribusi.")
	}
}

// printProjectHeader mencetak baris header yang diformat untuk tabel daftar proyek.
func printProjectHeader() {
	fmt.Printf("╔%s╦%s╦%s╦%s╦%s╦%s╦%s╗\n",
		strings.Repeat("═", idColWidth), strings.Repeat("═", nameColWidth),
		strings.Repeat("═", targetColWidth), strings.Repeat("═", currentColWidth),
		strings.Repeat("═", donaturColWidth), strings.Repeat("═", ownerColWidth),
		strings.Repeat("═", categoryColWidth))
	fmt.Printf("║%-*s║%-*s║%-*s║%-*s║%-*s║%-*s║%-*s║\n",
		idColWidth, "ID", nameColWidth, "Nama Proyek", targetColWidth, "Target",
		currentColWidth, "Terkumpul", donaturColWidth, "Donatur", ownerColWidth, "Nama Owner",
		categoryColWidth, "Kategori")
	fmt.Printf("╠%s╬%s╬%s╬%s╬%s╬%s╬%s╣\n",
		strings.Repeat("═", idColWidth), strings.Repeat("═", nameColWidth),
		strings.Repeat("═", targetColWidth), strings.Repeat("═", currentColWidth),
		strings.Repeat("═", donaturColWidth), strings.Repeat("═", ownerColWidth),
		strings.Repeat("═", categoryColWidth))
}

// printProjectDetail mencetak detail satu proyek dalam baris yang diformat untuk tabel.
func printProjectDetail(p Projek) {
	owner, _ := findPenggunaByID(p.OwnerID)
	ownerName := owner.Nama
	if len(ownerName) == 0 {
		ownerName = "N/A"
	}
	fmt.Printf("║%*d║%-*s║%*.2f║%*.2f║%*d║%-*s║%-*s║\n",
		idColWidth, p.ID, nameColWidth, p.Nama, targetColWidth, p.Target,
		currentColWidth, p.Current, donaturColWidth, p.JmlDonatur, ownerColWidth, ownerName,
		categoryColWidth, p.Category)
}

// printProjectFooter mencetak baris footer yang diformat untuk tabel daftar proyek.
func printProjectFooter() {
	fmt.Printf("╚%s╩%s╩%s╩%s╩%s╩%s╩%s╝\n",
		strings.Repeat("═", idColWidth), strings.Repeat("═", nameColWidth),
		strings.Repeat("═", targetColWidth), strings.Repeat("═", currentColWidth),
		strings.Repeat("═", donaturColWidth), strings.Repeat("═", ownerColWidth),
		strings.Repeat("═", categoryColWidth))
}

// viewAllProjects menampilkan daftar semua proyek penggalangan dana yang terdaftar.
func viewAllProjects() {
	fmt.Println("\n--- Daftar Semua Proyek ---")
	if countProject == 0 {
		fmt.Println("Belum ada proyek yang terdaftar.")
		return
	}
	printProjectHeader()
	for i := 0; i < countProject; i++ {
		printProjectDetail(projects[i])
	}
	printProjectFooter()
}

// viewMyProjects menampilkan daftar proyek yang dibuat oleh pemilik (owner) tertentu.
func viewMyProjects(ownerID int) {
	fmt.Println("\n--- Proyek yang Anda Buat ---")
	var ownerProjectCount int = 0
	for i := 0; i < countProject; i++ {
		if projects[i].OwnerID == ownerID {
			ownerProjectCount++
		}
	}
	if ownerProjectCount > 0 {
		printProjectHeader()
		for i := 0; i < countProject; i++ {
			if projects[i].OwnerID == ownerID {
				printProjectDetail(projects[i])
			}
		}
		printProjectFooter()
	} else {
		fmt.Println("Anda belum memiliki proyek yang terdaftar.")
	}
}

// searchProjectByNameSequential mencari proyek berdasarkan kata kunci nama proyek
// secara sekuensial dan tidak case-sensitive.
func searchProjectByNameSequential() {
	var query string
	fmt.Print("Masukkan Kata Kunci Nama Proyek yang dicari: ")
	fmt.Scanln(&query)
	fmt.Printf("\n--- Hasil Pencarian Nama dengan Kata Kunci : '%s' ---\n", query)
	lowerQuery := strings.ToLower(query)
	var foundAny bool = false
	var headerPrinted bool = false
	for i := 0; i < countProject; i++ {
		lowerProjectName := strings.ToLower(projects[i].Nama)
		if strings.Contains(lowerProjectName, lowerQuery) {
			if !headerPrinted {
				printProjectHeader()
				headerPrinted = true
			}
			printProjectDetail(projects[i])
			foundAny = true
		}
	}
	if foundAny {
		printProjectFooter()
	} else {
		fmt.Println("Tidak ada proyek yang cocok dengan kata kunci tersebut.")
	}
}

// searchProjectByCategory mencari proyek berdasarkan kata kunci kategori
// secara sekuensial dan tidak case-sensitive.
func searchProjectByCategory() {
	var query string
	fmt.Print("Masukkan Kata Kunci Kategori Proyek yang dicari: ")
	fmt.Scanln(&query)
	fmt.Printf("\n--- Hasil Pencarian Kategori dengan Kata Kunci : '%s' ---\n", query)
	if countProject == 0 {
		fmt.Println("Tidak ada proyek untuk dicari.")
		return
	}
	lowerQuery := strings.ToLower(query)
	var foundAny bool = false
	var headerPrinted bool = false
	for i := 0; i < countProject; i++ {
		lowerProjectCategory := strings.ToLower(projects[i].Category)
		if strings.Contains(lowerProjectCategory, lowerQuery) {
			if !headerPrinted {
				printProjectHeader()
				headerPrinted = true
			}
			printProjectDetail(projects[i])
			foundAny = true
		}
	}
	if foundAny {
		printProjectFooter()
	} else {
		fmt.Println("Tidak ada proyek yang cocok dengan kata kunci kategori tersebut.")
	}
}

// swapProjek menukar posisi dua elemen Projek dalam array global 'projects'.
func swapProjek(idx1 int, idx2 int) {
	projects[idx1], projects[idx2] = projects[idx2], projects[idx1]
}

// promptAscDescOrder meminta pengguna untuk memilih urutan pengurutan (Ascending/Menaik atau Descending/Menurun).
func promptAscDescOrder() (ascending bool, validInput bool) {
	var orderChoice int
	validInput = true
	fmt.Println("Mode urutan:")
	fmt.Println("1. Ascending (Menaik)")
	fmt.Println("2. Descending (Menurun)")
	fmt.Print("Pilih mode (1-2): ")
	fmt.Scanln(&orderChoice)
	if orderChoice == 1 {
		ascending = true
	} else if orderChoice == 2 {
		ascending = false
	} else {
		fmt.Println("Mode urutan tidak valid.")
		validInput = false
	}
	return
}

// selectionSortByDana mengurutkan array 'projects' berdasarkan jumlah dana yang terkumpul ('Current') menggunakan Selection Sort.
func selectionSortByDana(ascending bool) {
	n := countProject
	if n <= 1 {
		if n == 0 { fmt.Println("Belum ada proyek yang terdaftar.")
		} else { fmt.Println("Hanya ada satu proyek, tidak perlu diurutkan.") }
		viewAllProjects()
		return
	}
	for i := 0; i < n-1; i++ {
		bestIndex := i
		for j := i + 1; j < n; j++ {
			performSwap := false
			if ascending {
				if projects[j].Current < projects[bestIndex].Current { performSwap = true }
			} else {
				if projects[j].Current > projects[bestIndex].Current { performSwap = true }
			}
			if performSwap { bestIndex = j }
		}
		if bestIndex != i { swapProjek(i, bestIndex) }
	}
	orderStr := "Ascending"; if !ascending { orderStr = "Descending" }
	fmt.Printf("Proyek diurutkan berdasarkan Dana Terkumpul (%s) dengan Selection Sort.\n", orderStr)
	viewAllProjects()
}

// insertionSortByDonatur mengurutkan array 'projects' berdasarkan jumlah donatur ('JmlDonatur') menggunakan Insertion Sort.
func insertionSortByDonatur(ascending bool) {
	n := countProject
	if n <= 1 {
		if n == 0 { fmt.Println("Belum ada proyek yang terdaftar.")
		} else { fmt.Println("Hanya ada satu proyek, tidak perlu diurutkan.") }
		viewAllProjects()
		return
	}
	for i := 1; i < n; i++ {
		key := projects[i]
		j := i - 1
		keepMoving := true // Flag untuk mengontrol loop dalam
		for j >= 0 && keepMoving {
			shouldMove := false
			if ascending {
				if projects[j].JmlDonatur > key.JmlDonatur { shouldMove = true }
			} else {
				if projects[j].JmlDonatur < key.JmlDonatur { shouldMove = true }
			}
			if shouldMove {
				projects[j+1] = projects[j]
				j--
			} else {
				keepMoving = false // Menghentikan loop dalam (alternatif 'break')
			}
		}
		projects[j+1] = key
	}
	orderStr := "Ascending"; if !ascending { orderStr = "Descending" }
	fmt.Printf("Proyek diurutkan berdasarkan Jumlah Donatur (%s) dengan Insertion Sort.\n", orderStr)
	viewAllProjects()
}

// viewFundedProjects menampilkan daftar proyek yang telah mencapai atau melampaui target pendanaan mereka.
func viewFundedProjects() {
	fmt.Println("\n--- Daftar Proyek yang Telah Mencapai Target Pendanaan ---")
	var foundAny bool = false
	var headerPrinted bool = false
	for i := 0; i < countProject; i++ {
		p := projects[i]
		if p.Current >= p.Target {
			if !headerPrinted { printProjectHeader(); headerPrinted = true }
			printProjectDetail(p)
			foundAny = true
		}
	}
	if foundAny { printProjectFooter()
	} else { fmt.Println("Belum ada proyek yang mencapai target pendanaan.") }
}

// showLoggedInMenu menampilkan menu utama aplikasi setelah pengguna berhasil login.
func showLoggedInMenu(user Pengguna) {
	var choice int
	var stayInMenu bool = true
	for stayInMenu {
		CLS()
		const totalWidth = totalTableWidth + 2
		const hLine, vLine, cLine, columnSeparator = '═', '║', '─', '│'
		const itemLeftPad = 1
		fmt.Printf("╔%s╗\n", strings.Repeat(string(hLine), totalWidth-2))
		title := "--- APLIKASI CROWDFUNDING ---"
		paddingTitle := (totalWidth - 2 - len(title)) / 2
		if paddingTitle < 0 { paddingTitle = 0 }
		remainingPaddingTitle := totalWidth - 2 - len(title) - paddingTitle
		if remainingPaddingTitle < 0 { remainingPaddingTitle = 0 }
		fmt.Printf("%c%s%s%s%c\n", vLine, strings.Repeat(" ", paddingTitle), title, strings.Repeat(" ", remainingPaddingTitle), vLine)
		fmt.Printf("╚%s╝\n", strings.Repeat(string(hLine), totalWidth-2))
		fmt.Printf("┌%s┐\n", strings.Repeat(string(cLine), totalWidth-2))
		userInfoContent := fmt.Sprintf("👤 Selamat Datang, %s (ID: %d | Tipe: %s)", user.Nama, user.ID, user.TipePengguna)
		if len(userInfoContent) > totalWidth-4 { userInfoContent = userInfoContent[:totalWidth-7] + "..." }
		fmt.Printf("│ %-*s │\n", totalWidth-4, userInfoContent)
		fmt.Printf("└%s┘\n", strings.Repeat(string(cLine), totalWidth-2))
		menuHeader := " MENU UTAMA "; menuHeaderPadding := (totalWidth - 2 - len(menuHeader)) / 2
		if menuHeaderPadding < 0 { menuHeaderPadding = 0 }
		remainingMenuHeaderPadding := totalWidth - 2 - len(menuHeader) - menuHeaderPadding
		if remainingMenuHeaderPadding < 0 { remainingMenuHeaderPadding = 0 }
		fmt.Printf("\n┌%s%s%s┐\n", strings.Repeat(string(cLine), menuHeaderPadding), menuHeader, strings.Repeat(string(cLine), remainingMenuHeaderPadding))
		var menuOptions []string; var maxOptNum int
		if user.TipePengguna == "owner" {
			menuOptions = []string{
				"Buat Proyek Baru", "Lihat Proyek Saya", "Ubah Proyek Saya", "Hapus Proyek Saya",
				"Lihat Semua Proyek", "Cari Proyek berdasarkan Nama", "Cari Proyek berdasarkan Kategori",
				"Urutkan Proyek berdasarkan Dana Terkumpul", "Urutkan Proyek berdasarkan Jumlah Donatur",
				"Lihat Proyek Capai Target",
			}; maxOptNum = 10
		} else if user.TipePengguna == "user" {
			menuOptions = []string{
				"Lihat Semua Proyek", "Berkontribusi ke Proyek", "Cari Proyek berdasarkan Nama",
				"Cari Proyek berdasarkan Kategori", "Urutkan Proyek berdasarkan Dana Terkumpul",
				"Urutkan Proyek berdasarkan Jumlah Donatur", "Lihat Proyek Capai Target",
			}; maxOptNum = 7
		}
		const innerMenuWidth = totalWidth - 2
		menuItemTextWidthCol1 := (innerMenuWidth - 1 - (2 * itemLeftPad)) / 2
		menuItemTextWidthCol2 := innerMenuWidth - 1 - (2 * itemLeftPad) - menuItemTextWidthCol1
		fmt.Printf("%c%s%c\n", vLine, strings.Repeat(" ", innerMenuWidth), vLine)
		optCounter := 1
		for i := 0; i < len(menuOptions); i += 2 {
			formattedOpt1 := fmt.Sprintf("[%d] %s", optCounter, menuOptions[i]); optCounter++
			formattedOpt2 := ""; if i+1 < len(menuOptions) { formattedOpt2 = fmt.Sprintf("[%d] %s", optCounter, menuOptions[i+1]); optCounter++ }
			if len(formattedOpt1) > menuItemTextWidthCol1 { formattedOpt1 = formattedOpt1[:menuItemTextWidthCol1-3] + "..." }
			if len(formattedOpt2) > menuItemTextWidthCol2 && formattedOpt2 != "" { formattedOpt2 = formattedOpt2[:menuItemTextWidthCol2-3] + "..." }
			fmt.Printf("%c%s%-*s%c%s%-*s%c\n", vLine, strings.Repeat(" ", itemLeftPad), menuItemTextWidthCol1, formattedOpt1,
				columnSeparator, strings.Repeat(" ", itemLeftPad), menuItemTextWidthCol2, formattedOpt2, vLine)
		}
		fmt.Printf("%c%s%c\n", vLine, strings.Repeat(" ", innerMenuWidth), vLine)
		logoutText := fmt.Sprintf("[%d] Logout", 0)
		fmt.Printf("%c%s%-*s%c\n", vLine, strings.Repeat(" ", itemLeftPad), innerMenuWidth-itemLeftPad, logoutText, vLine)
		fmt.Printf("%c%s%c\n", vLine, strings.Repeat(" ", innerMenuWidth), vLine)
		fmt.Printf("└%s┘\n", strings.Repeat(string(cLine), innerMenuWidth))
		if stayInMenu {
			fmt.Printf(">> Pilih Opsi (0-%d): ", maxOptNum); fmt.Scanln(&choice)
			actionHandled := true
			if user.TipePengguna == "owner" {
				switch choice {
				case 1: createProject(user.ID); case 2: viewMyProjects(user.ID); case 3: editProject(user.ID)
				case 4: deleteProject(user.ID); case 5: viewAllProjects(); case 6: searchProjectByNameSequential()
				case 7: searchProjectByCategory(); case 8: asc, valid := promptAscDescOrder(); if valid { selectionSortByDana(asc) }
				case 9: asc, valid := promptAscDescOrder(); if valid { insertionSortByDonatur(asc) }
				case 10: viewFundedProjects(); case 0: fmt.Println("Logout berhasil."); stayInMenu = false; actionHandled = false
				default: fmt.Println("Pilihan menu tidak tersedia.")
				}
			} else if user.TipePengguna == "user" {
				switch choice {
				case 1: viewAllProjects(); case 2: contributeToProject(user.ID); case 3: searchProjectByNameSequential()
				case 4: searchProjectByCategory(); case 5: asc, valid := promptAscDescOrder(); if valid { selectionSortByDana(asc) }
				case 6: asc, valid := promptAscDescOrder(); if valid { insertionSortByDonatur(asc) }
				case 7: viewFundedProjects(); case 0: fmt.Println("Logout berhasil."); stayInMenu = false; actionHandled = false
				default: fmt.Println("Pilihan menu tidak tersedia.")
				}
			}
			if actionHandled { pauseExecution() }
		}
	}
}

// loadDummyData menginisialisasi aplikasi dengan data pengguna dan proyek dummy.
func loadDummyData() {
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "owner", Password: "owner1", Nama: "Andi_Kreator"}; countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "owner", Password: "owner2", Nama: "Citra_Inovasi"}; countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "user", Password: "user1", Nama: "Doni_Peduli"}; countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "user", Password: "user2", Nama: "Elisa_Baikhati"}; countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "owner", Password: "owner", Nama: "owner"}; countPengguna++; nextPenggunaID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Game_Edukasi_Anak", Target: 2500000, OwnerID: 1, Category: "Teknologi"}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Rumah_Singgah_Hewan", Target: 5000000, OwnerID: 1, Category: "Sosial"}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Film_Pendek_Dokumenter", Target: 3000000, OwnerID: 2, Category: "Seni"}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Pelatihan_UMKM_Digital", Target: 4000000, OwnerID: 2, Category: "Edukasi"}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Penggalangan_Dana_Darurat", Target: 999999999.00, OwnerID: 1, Category: "Sosial"}; countProject++; nextProjectID++
	addDummyContributionRecord := func(userID, projectID int, amount float64) {
		if countKontribusi >= maxKontribusi { return }
		projectIdx := findProjectArrayIndex(projectID)
		if projectIdx != -1 {
			projects[projectIdx].Current += amount; projects[projectIdx].JmlDonatur++
			contributions[countKontribusi] = Kontribusi{ID: nextKontribusiID, ProjectID: projectID, PenggunaID: userID, Jumlah: amount}
			countKontribusi++; nextKontribusiID++
		}
	}
	addDummyContributionRecord(3, 1, 1000000); addDummyContributionRecord(4, 1, 1500000)
	addDummyContributionRecord(3, 2, 2000000); addDummyContributionRecord(4, 2, 500000)
	addDummyContributionRecord(3, 3, 500000); addDummyContributionRecord(4, 4, 1000000)
}

// main adalah titik masuk utama aplikasi.
func main() {
	loadDummyData()
	var loggedInUser Pengguna
	var isLoggedIn bool = false
	var stayInApp bool = true
	for stayInApp {
		CLS()
		if !isLoggedIn {
			const loginMenuWidth = totalTableWidth + 2
			fmt.Printf("╔%s╗\n", strings.Repeat("═", loginMenuWidth-2))
			loginTitle := "--- APLIKASI CROWDFUNDING ---"
			loginPaddingTitle := (loginMenuWidth - 2 - len(loginTitle)) / 2
			if loginPaddingTitle < 0 { loginPaddingTitle = 0 }
			remainingLoginPaddingTitle := loginMenuWidth - 2 - len(loginTitle) - loginPaddingTitle
			if remainingLoginPaddingTitle < 0 { remainingLoginPaddingTitle = 0 }
			fmt.Printf("║%s%s%s║\n", strings.Repeat(" ", loginPaddingTitle), loginTitle, strings.Repeat(" ", remainingLoginPaddingTitle))
			fmt.Printf("╚%s╝\n", strings.Repeat("═", loginMenuWidth-2))
			fmt.Printf("┌%s┐\n", strings.Repeat("─", loginMenuWidth-2))
			fmt.Printf("│ %-*s │\n", loginMenuWidth-4, "1. Login")
			fmt.Printf("│ %-*s │\n", loginMenuWidth-4, "2. Sign Up")
			fmt.Printf("│ %-*s │\n", loginMenuWidth-4, "0. Keluar Aplikasi")
			fmt.Printf("└%s┘\n", strings.Repeat("─", loginMenuWidth-2))
			fmt.Print(">> Pilih Opsi: ")
			var initialChoice int; fmt.Scanln(&initialChoice)
			var shouldPause bool = true
			switch initialChoice {
			case 1:
				var inputID int; var inputPassword string
				fmt.Println("\n--- Login Pengguna ---")
				fmt.Print("Masukkan ID Pengguna: "); fmt.Scanln(&inputID)
				fmt.Print("Masukkan Password: "); fmt.Scanln(&inputPassword)
				user, success := loginPengguna(inputID, inputPassword)
				if success { loggedInUser = user; isLoggedIn = true; fmt.Println("Login berhasil!"); shouldPause = false
				} else { fmt.Println("Login gagal.") }
			case 2: signUp()
			case 0: fmt.Println("Terima kasih!"); stayInApp = false; shouldPause = false
			default: fmt.Println("Pilihan tidak valid.")
			}
			if shouldPause && stayInApp { pauseExecution() }
		}
		if isLoggedIn {
			showLoggedInMenu(loggedInUser)
			isLoggedIn = false; loggedInUser = Pengguna{}
		}
	}
}