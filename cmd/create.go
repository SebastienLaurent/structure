package cmd

import (
	memorymanager "github.com/SebastienLaurent/structure/repository/manager/memory"
)

// Pour faire taire les warning "is not used"
func use[T any](t T) {}

func Create() {
	manager := memorymanager.New()

	// repo1 and repo2 sont differents, cela fait partie du contrat des managers
	repo1, _ := manager.Get("git@github.com/SebastienLaurent/structure")
	repo2, _ := manager.Get("git@github.com/SebastienLaurent/structure")
	
	fs , _ := repo1.Filesystem()
	
	// Creation d'un fichier
	f, _ := fs.Create("myfile.txt")
	defer f.Close()

	// Ecriture comme sur un fichier normal
	f.Write([]byte("Test"))
	
	// Simlink
	fs.Symlink("mylink.txt","myfile.txt")

	use(repo2)
}