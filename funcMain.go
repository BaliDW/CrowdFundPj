package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

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

func ClearScreen() {
	clearScreenCustom(os.Stdout)
}

func clearScreenCustom(stdout *os.File) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Debug: Perintah clear screen gagal. OS: '%s', Error: %v\n", runtime.GOOS, err)
		var i int = 0
		for i < 50 {
			fmt.Println()
			i++
		}
		fmt.Println("Fallback newline digunakan karena clear screen gagal (perintah sistem tidak berhasil).")
	}
}

func pauseExecution() {
	fmt.Print("\nTekan Enter untuk melanjutkan...")
	var dummy string
	fmt.Scanln(&dummy)
}

func ResetStyle() {
}

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

func loginPengguna(id int, password string) (Pengguna, bool) {
	user, found := findPenggunaByID(id)
	var success bool = false
	if found {
		success = (user.Password == password)
	}
	return user, success
}

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

func createProject(ownerID int) {
	if countProject >= maxProject {
		fmt.Println("Kapasitas proyek penuh!")
		return
	}
	var projectName, projectCategory string
	var projectTarget float64

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
	newProjek := Projek{ID: nextProjectID, Nama: projectName, Target: projectTarget, Current: 0.0, JmlDonatur: 0, OwnerID: ownerID, Category: projectCategory}
	projects[countProject] = newProjek
	countProject++
	nextProjectID++
	fmt.Println("Proyek berhasil dibuat!")
	fmt.Printf("ID Proyek: %d, Nama: '%s', Kategori: '%s', Target: %.2f\n", newProjek.ID, newProjek.Nama, newProjek.Category, newProjek.Target)
}

func editProject(ownerID int) {
	fmt.Println("\n--- Ubah Proyek Anda ---")
	viewMyProjects(ownerID)
	if countProject == 0 { return }

	var projectIDToEdit int
	fmt.Print("Masukkan ID Proyek yang ingin diubah (0 untuk batal): ")
	fmt.Scanln(&projectIDToEdit)

	if projectIDToEdit == 0 { fmt.Println("Pengubahan dibatalkan."); return }

	currentProject, projectFound := findProjekByID(projectIDToEdit)
	if !projectFound { fmt.Println("Proyek dengan ID tersebut tidak ditemukan."); return }
	if currentProject.OwnerID != ownerID { fmt.Println("Anda bukan pemilik proyek ini."); return }

	projectIdx := findProjectArrayIndex(projectIDToEdit)
	if projectIdx == -1 { fmt.Println("Internal error: Proyek tidak ditemukan di array."); return }

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
	if newNama != "-" && len(newNama) > 0 { projects[projectIdx].Nama = newNama; changed = true }
	if newKategori != "-" && len(newKategori) > 0 { projects[projectIdx].Category = newKategori; changed = true }
	if newTarget > 0 { projects[projectIdx].Target = newTarget; changed = true
	} else if newTarget < 0 { fmt.Println("Target dana baru tidak valid (negatif). Target tidak diubah.") }

	if changed { fmt.Println("Proyek berhasil diubah.")
	} else { fmt.Println("Tidak ada perubahan yang dilakukan.") }
}

func deleteProject(ownerID int) {
	fmt.Println("\n--- Hapus Proyek Anda ---")
	viewMyProjects(ownerID)
	if countProject == 0 { return }

	var projectIDToDelete int
	fmt.Print("Masukkan ID Proyek yang ingin dihapus (0 untuk batal): ")
	fmt.Scanln(&projectIDToDelete)

	if projectIDToDelete == 0 { fmt.Println("Penghapusan dibatalkan."); return }

	currentProject, projectFound := findProjekByID(projectIDToDelete)
	if !projectFound { fmt.Println("Proyek dengan ID tersebut tidak ditemukan."); return }
	if currentProject.OwnerID != ownerID { fmt.Println("Anda bukan pemilik proyek ini."); return }

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
			fmt.Println("Internal Error: Proyek tidak ditemukan di array untuk dihapus.")
		}
	} else {
		fmt.Println("Penghapusan proyek dibatalkan.")
	}
}

func contributeToProject(userID int) {
	fmt.Println("\n--- Berkontribusi ke Proyek ---")
	if countProject == 0 { fmt.Println("Belum ada proyek yang bisa didanai."); return }
	viewAllProjects()

	var projectIDToDonate int
	fmt.Print("Masukkan ID Proyek yang ingin didanai: ")
	fmt.Scanln(&projectIDToDonate)

	targetProjek, projekFound := findProjekByID(projectIDToDonate)
	if !projekFound { fmt.Println("Gagal berkontribusi: Proyek dengan ID tersebut tidak ditemukan."); return }

	var amount float64
	fmt.Printf("Anda akan berkontribusi untuk proyek: %s\n", targetProjek.Nama)
	fmt.Print("Masukkan jumlah kontribusi: ")
	fmt.Scanln(&amount)

	if amount <= 0 { fmt.Println("Jumlah kontribusi harus lebih dari 0."); return }
	if countKontribusi >= maxKontribusi { fmt.Println("Kapasitas data kontribusi penuh."); return }

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
		fmt.Println("Internal Error: Proyek tidak ditemukan di array untuk kontribusi.")
	}
}

func printProjectHeader() {
	fmt.Printf("%-5s %-28s %-12s %-12s %-10s %-15s %-15s\n", "ID", "Nama Proyek", "Target", "Terkumpul", "Donatur", "Owner Name", "Kategori")
	fmt.Println("-------------------------------------------------------------------------------------------------------")
}

func printProjectDetail(p Projek) {
	owner, _ := findPenggunaByID(p.OwnerID)
	ownerName := owner.Nama
	if len(ownerName) == 0 { ownerName = "N/A" }
	fmt.Printf("%-5d %-28s %-12.2f %-12.2f %-10d %-15s %-15s\n",
		p.ID, p.Nama, p.Target, p.Current, p.JmlDonatur, ownerName, p.Category)
}

func viewAllProjects() {
	fmt.Println("\n--- Daftar Semua Proyek ---")
	if countProject == 0 { fmt.Println("Belum ada proyek yang terdaftar."); return }
	printProjectHeader()
	var i int = 0
	for i < countProject {
		printProjectDetail(projects[i])
		i++
	}
}

func viewMyProjects(ownerID int) {
	fmt.Println("\n--- Proyek yang Anda Buat ---")
	var ownerProjectCount int = 0
	var i int = 0
	for i < countProject {
		if projects[i].OwnerID == ownerID {
			ownerProjectCount++
		}
		i++
	}

	if ownerProjectCount > 0 {
		printProjectHeader()
		i = 0
		for i < countProject {
			if projects[i].OwnerID == ownerID {
				printProjectDetail(projects[i])
			}
			i++
		}
	} else {
		fmt.Println("Anda belum memiliki proyek yang terdaftar.")
	}
}

func searchProjectByNameSequential() {
	var query string
	fmt.Print("Masukkan Nama Proyek yang dicari: ")
	fmt.Scanln(&query)
	fmt.Printf("\n--- Hasil Pencarian Nama : '%s' ---\n", query)

	var foundAny bool = false
	var headerPrinted bool = false
	var i int = 0
	for i < countProject {
		if projects[i].Nama == query {
			if !headerPrinted { printProjectHeader(); headerPrinted = true }
			printProjectDetail(projects[i])
			foundAny = true
		}
		i++
	}
	if !foundAny { fmt.Println("Tidak ada proyek yang cocok dengan nama tersebut.") }
}

func searchProjectByCategoryBinary() {
	var query string
	fmt.Print("Masukkan Kategori Proyek yang dicari: ")
	fmt.Scanln(&query)
	fmt.Printf("\n--- Hasil Pencarian Kategori : '%s' ---\n", query)

	if countProject == 0 { fmt.Println("Tidak ada proyek untuk dicari."); return }

	internalSortByCategoryAscending()
	fmt.Println("Data proyek telah diurutkan berdasarkan Kategori (Asc) untuk pencarian biner.")

	var low, high, mid int = 0, countProject - 1, 0
	var foundProjek Projek
	var foundMatch bool = false

	for low <= high && !foundMatch {
		mid = low + (high-low)/2
		if projects[mid].Category == query {
			foundProjek = projects[mid]
			foundMatch = true
		} else if projects[mid].Category < query {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if foundMatch { printProjectHeader(); printProjectDetail(foundProjek)
	} else { fmt.Println("Proyek dengan kategori tersebut tidak ditemukan.") }
}

func swapProjek(idx1 int, idx2 int) {
	projects[idx1], projects[idx2] = projects[idx2], projects[idx1]
}

func internalSortByCategoryAscending() {
	n := countProject
	if n <= 1 { return }
	var i int = 0
	for i < n-1 {
		minIndex := i
		var j int = i + 1
		for j < n {
			if projects[j].Category < projects[minIndex].Category {
				minIndex = j
			}
			j++
		}
		if minIndex != i { swapProjek(i, minIndex) }
		i++
	}
}

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

func selectionSortByDana(ascending bool) {
	n := countProject
	if n <= 1 { fmt.Println("Tidak cukup data untuk diurutkan."); return }
	
	var i int = 0
	for i < n-1 {
		bestIndex := i
		var j int = i + 1
		for j < n {
			performSwap := false
			if ascending {
				if projects[j].Current < projects[bestIndex].Current { performSwap = true }
			} else {
				if projects[j].Current > projects[bestIndex].Current { performSwap = true }
			}
			if performSwap { bestIndex = j }
			j++
		}
		if bestIndex != i { swapProjek(i, bestIndex) }
		i++
	}
	orderStr := "Ascending"; if !ascending { orderStr = "Descending" }
	fmt.Printf("Proyek diurutkan berdasarkan Dana Terkumpul (%s) dengan Selection Sort.\n", orderStr)
	viewAllProjects()
}

func insertionSortByDonatur(ascending bool) {
	n := countProject
	if n <= 1 { fmt.Println("Tidak cukup data untuk diurutkan."); return }

	var i int = 1
	for i < n {
		key := projects[i]
		j := i - 1
		keepMoving := true
		for j >= 0 && keepMoving {
			shouldMove := false
			if ascending {
				if projects[j].JmlDonatur > key.JmlDonatur { shouldMove = true }
			} else {
				if projects[j].JmlDonatur < key.JmlDonatur { shouldMove = true }
			}
			if shouldMove { projects[j+1] = projects[j]; j--
			} else { keepMoving = false }
		}
		projects[j+1] = key
		i++
	}
	orderStr := "Ascending"; if !ascending { orderStr = "Descending" }
	fmt.Printf("Proyek diurutkan berdasarkan Jumlah Donatur (%s) dengan Insertion Sort.\n", orderStr)
	viewAllProjects()
}

func viewFundedProjects() {
	fmt.Println("\n--- Daftar Proyek yang Telah Mencapai Target Pendanaan ---")
	var foundAny bool = false
	var headerPrinted bool = false
	var i int = 0
	for i < countProject {
		p := projects[i]
		if p.Current >= p.Target {
			if !headerPrinted { printProjectHeader(); headerPrinted = true }
			printProjectDetail(p)
			foundAny = true
		}
		i++
	}
	if !foundAny { fmt.Println("Belum ada proyek yang mencapai target pendanaan.") }
}

func showLoggedInMenu(user Pengguna) {
	var choice int
	var stayInMenu bool = true
	for stayInMenu {
		ClearScreen()
		fmt.Printf("--- Selamat Datang, %s (%s)! ---\n", user.Nama, user.TipePengguna)

		fmt.Println("\n--- MENU ---")
		if user.TipePengguna == "owner" {
			fmt.Println("1. Buat Proyek Baru")
			fmt.Println("2. Lihat Proyek Saya")
			fmt.Println("3. Ubah Proyek Saya")
			fmt.Println("4. Hapus Proyek Saya")
			fmt.Println("5. Lihat Semua Proyek")
			fmt.Println("6. Cari Proyek berdasarkan Nama ")
			fmt.Println("7. Cari Proyek berdasarkan Kategori ")
			fmt.Println("8. Urutkan Proyek berdasarkan Dana Terkumpul ")
			fmt.Println("9. Urutkan Proyek berdasarkan Jumlah Donatur ")
			fmt.Println("10. Lihat Proyek Capai Target")
			fmt.Println("0. Logout")
		} else if user.TipePengguna == "user" {
			fmt.Println("1. Lihat Semua Proyek")
			fmt.Println("2. Berkontribusi ke Proyek")
			fmt.Println("3. Cari Proyek berdasarkan Nama ")
			fmt.Println("4. Cari Proyek berdasarkan Kategori ")
			fmt.Println("5. Urutkan Proyek berdasarkan Dana Terkumpul ")
			fmt.Println("6. Urutkan Proyek berdasarkan Jumlah Donatur ")
			fmt.Println("7. Lihat Proyek Capai Target")
			fmt.Println("0. Logout")
		} else { fmt.Println("Tipe pengguna tidak dikenal."); stayInMenu = false }

		if stayInMenu {
			fmt.Print("Pilih Opsi: ")
			fmt.Scanln(&choice)
			actionHandled := true 

			if user.TipePengguna == "owner" {
				if choice == 1 { createProject(user.ID) } else
				if choice == 2 { viewMyProjects(user.ID) } else
				if choice == 3 { editProject(user.ID) } else
				if choice == 4 { deleteProject(user.ID) } else
				if choice == 5 { viewAllProjects() } else
				if choice == 6 { searchProjectByNameSequential() } else
				if choice == 7 { searchProjectByCategoryBinary() } else
				if choice == 8 { 
					asc, valid := promptAscDescOrder()
					if valid { selectionSortByDana(asc) }
				} else
				if choice == 9 { 
					asc, valid := promptAscDescOrder()
					if valid { insertionSortByDonatur(asc) }
				} else
				if choice == 10 { viewFundedProjects() } else
				if choice == 0 { fmt.Println("Logout berhasil."); stayInMenu = false; actionHandled = false } else
				{ fmt.Println("Pilihan menu tidak tersedia.") }
			} else if user.TipePengguna == "user" {
				if choice == 1 { viewAllProjects() } else
				if choice == 2 { contributeToProject(user.ID) } else
				if choice == 3 { searchProjectByNameSequential() } else
				if choice == 4 { searchProjectByCategoryBinary() } else
				if choice == 5 { 
					asc, valid := promptAscDescOrder()
					if valid { selectionSortByDana(asc) }
				} else
				if choice == 6 { 
					asc, valid := promptAscDescOrder()
					if valid { insertionSortByDonatur(asc) }
				} else
				if choice == 7 { viewFundedProjects() } else
				if choice == 0 { fmt.Println("Logout berhasil."); stayInMenu = false; actionHandled = false } else
				{ fmt.Println("Pilihan menu tidak tersedia.") }
			}
			if actionHandled { pauseExecution() }
		}
	}
}

func loadDummyData() {
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "owner", Password: "owner1", Nama: "Andi Kreator"}
	countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "owner", Password: "owner2", Nama: "Citra Inovasi"}
	countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "user", Password: "user1", Nama: "Doni Peduli"}
	countPengguna++; nextPenggunaID++
	users[countPengguna] = Pengguna{ID: nextPenggunaID, TipePengguna: "user", Password: "user2", Nama: "Elisa Baikhati"}
	countPengguna++; nextPenggunaID++
	

	projects[countProject] = Projek{ID: nextProjectID, Nama: "Game Edukasi Anak", Target: 2500000, Current: 0, JmlDonatur: 0, OwnerID: 1, Category: "Teknologi",}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Rumah Singgah Hewan", Target: 5000000, Current: 0, JmlDonatur: 0, OwnerID: 1, Category: "Sosial",}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Film Pendek Dokumenter", Target: 3000000, Current: 0, JmlDonatur: 0, OwnerID: 2, Category: "Seni",}; countProject++; nextProjectID++
	projects[countProject] = Projek{ID: nextProjectID, Nama: "Pelatihan UMKM Digital", Target: 4000000, Current: 0, JmlDonatur: 0, OwnerID: 2, Category: "Edukasi",}; countProject++; nextProjectID++
	

	addDummyContributionRecord := func(userID, projectID int, amount float64) {
		if countKontribusi >= maxKontribusi { return }
		projectIdx := findProjectArrayIndex(projectID)
		if projectIdx != -1 {
			projects[projectIdx].Current += amount
			projects[projectIdx].JmlDonatur++
			contributions[countKontribusi] = Kontribusi{ID: nextKontribusiID, ProjectID: projectID, PenggunaID: userID, Jumlah: amount,}; countKontribusi++; nextKontribusiID++
		}
	}
	addDummyContributionRecord(3, 1, 1000000)
	addDummyContributionRecord(3, 3, 500000)
	addDummyContributionRecord(4, 1, 1500000)
	addDummyContributionRecord(4, 2, 2000000)
	addDummyContributionRecord(4, 4, 1000000)
	addDummyContributionRecord(3, 2, 500000)
	
}

func main() {
	loadDummyData()

	var loggedInUser Pengguna
	var isLoggedIn bool = false
	var stayInApp bool = true

	for stayInApp {
		ClearScreen()
		if !isLoggedIn {
			fmt.Println("--- Aplikasi Crowdfunding ---")
			fmt.Println("1. Login")
			fmt.Println("2. Sign Up")
			fmt.Println("0. Keluar Aplikasi")
			fmt.Print("Pilih Opsi: ")
			var initialChoice int
			fmt.Scanln(&initialChoice)

			actionHandled := true
			if initialChoice == 1 {
				var inputID int; var inputPassword string
				fmt.Println("\n--- Login Pengguna ---")
				fmt.Print("Masukkan ID Pengguna: "); fmt.Scanln(&inputID)
				fmt.Print("Masukkan Password: "); fmt.Scanln(&inputPassword)
				user, success := loginPengguna(inputID, inputPassword)
				if success {
					loggedInUser = user; isLoggedIn = true; fmt.Println("Login berhasil!")
					actionHandled = false
				} else { fmt.Println("Login gagal.") }
			} else if initialChoice == 2 { signUp() } else
			if initialChoice == 0 {
				fmt.Println("Terima kasih!")
				ResetStyle()
				stayInApp = false; actionHandled = false
			} else
			if initialChoice != 1 && initialChoice != 2 && initialChoice != 0 {
				fmt.Println("Pilihan tidak valid.")
			}
			if actionHandled { pauseExecution() }
		}

		if isLoggedIn {
			showLoggedInMenu(loggedInUser)
			isLoggedIn = false
			loggedInUser = Pengguna{}
			ResetStyle()
		}
	}
	ResetStyle()
}