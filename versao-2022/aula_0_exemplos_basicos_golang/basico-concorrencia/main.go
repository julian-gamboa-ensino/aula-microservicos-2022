package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

func downloadImage(url string, wg *sync.WaitGroup, imageQueue chan<- []byte) {
	defer wg.Done()
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Erro ao baixar imagem de %s: %v\n", url, err)
		return
	}
	defer response.Body.Close()

	imageData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Erro ao ler dados da imagem de %s: %v\n", url, err)
		return
	}

	imageQueue <- imageData
	log.Printf("Imagem de %s baixada com sucesso\n", url)
}

func main() {
	// Configurando o arquivo de log
	logFile, err := os.OpenFile("download_log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Erro ao abrir arquivo de log: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	imageUrls := []string{
		  "https://cv-julian-2022.s3.us-west-2.amazonaws.com/certificados/E02D0099.jpg",
    "https://cv-julian-2022.s3.us-west-2.amazonaws.com/certificados/F9EA66D0.jpg",
    "https://cv-julian-2022.s3.us-west-2.amazonaws.com/certificados/EE3EC6C2.jpg",
	}

	imageQueue := make(chan []byte, len(imageUrls))
	var wg sync.WaitGroup
	downloadedImages := make([][]byte, 0, len(imageUrls))

	for _, url := range imageUrls {
		wg.Add(1)
		go downloadImage(url, &wg, imageQueue)
	}

	go func() {
		wg.Wait()
		close(imageQueue)
	}()

	for imageData := range imageQueue {
		downloadedImages = append(downloadedImages, imageData)
	}

	log.Println("Todas as imagens foram baixadas e processadas com sucesso!")
	fmt.Println("Todas as imagens foram baixadas e processadas com sucesso!")
}
