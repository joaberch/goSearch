package treebuilder

import (
	"os"
	"path/filepath"
	"testing"
)

// test made by AI, reviewed by human (not all possibility are tested but the unit testing is not the principal objective right now)
func TestGetFileTree(t *testing.T) {
	// Crée un répertoire temporaire pour le test
	root := t.TempDir()

	// Crée une structure de fichiers fictive
	subDir := filepath.Join(root, "subdir")
	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Erreur lors de la création du dossier : %v", err)
	}

	// Crée un fichier fictif
	file1 := filepath.Join(root, "file1.txt")
	err = os.WriteFile(file1, []byte("contenu"), 0644)
	if err != nil {
		t.Fatalf("Erreur lors de la création du fichier : %v", err)
	}

	// Appelle la fonction à tester
	tree := GetFileTree([]string{root})

	// Vérifie les résultats
	if len(tree.Children) != 1 {
		t.Errorf("Attendu 1 enfant à la racine, obtenu %d", len(tree.Children))
	}

	rootNode := tree.Children[0]
	if rootNode.Name != filepath.Base(root) {
		t.Errorf("Nom incorrect : attendu %s, obtenu %s", filepath.Base(root), rootNode.Name)
	}

	if len(rootNode.Children) != 2 {
		t.Errorf("Attendu 2 enfants dans le dossier racine, obtenu %d", len(rootNode.Children))
	}
}
