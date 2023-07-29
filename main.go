package main

import (
	"fmt"
	"github.com/GabrielPMagni/GoFileOrganizer/methods"
	"github.com/rudolfoborges/pdf2go"
)

func main() {

	folderScanner := FolderScanner{}
	folderScanner.listFolderFiles("./arquivos/", true)

	pdf, err := pdf2go.New("./arquivos/NFSE_1_79977_1_1.pdf", pdf2go.Config{
		LogLevel: pdf2go.LogLevelError,
	})

	if err != nil {
		panic(err)
	}

	text, err := pdf.Text()
	if err != nil {
		panic(err)
	}

	pages, err := pdf.Pages()

	if err != nil {
		panic(err)
	}

	for _, page := range pages {
		fmt.Println(page.Text())
	}
}

type FolderScanner struct {
	Files []string
}

func (fs *FolderScanner) listFolderFiles(dir string, debug bool) {
	if debug {
		fmt.Println("Listando diretórios...")
	}
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Erro ao ler o diretório:", err)
		return
	}

	for _, item := range items {
		d := filepath.Join(dir, item.Name())
		if item.IsDir() {
			if debug {
				fmt.Println("Pasta encontrada")
			}
			fs.listFolderFiles(d, debug)
		} else {
			if debug {
				fmt.Println("Arquivo Encontrado")
			}
			fs.Files = append(fs.Files, d)
		}
	}

	if len(fs.Files) == 0 {
		fmt.Println("Não encontrados arquivos")
		os.Exit(3)
	}
}
