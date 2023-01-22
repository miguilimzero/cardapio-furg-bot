package main

import "fmt"

func makePrintAndUpload(url string, filename string) string {
	takeScreenshot(filename, url)

	res := uploadMedia(filename)

	return res["media_id_string"].(string)
}

func main() {
	var mediaIds []string

	mediaIds = append(mediaIds, makePrintAndUpload("https://www.furg.br/estudantes/cardapio-ru/restaurante-universitario-lago", "stubs/ru-lago.png"))
	mediaIds = append(mediaIds, makePrintAndUpload("https://www.furg.br/estudantes/cardapio-ru/restaurante-universitario-cc", "stubs/ru-cc.png"))
	mediaIds = append(mediaIds, makePrintAndUpload("https://www.furg.br/estudantes/cardapio-ru/restaurante-universitario-ccmar", "stubs/ru-ccmar.png"))

	res2 := sendTweet("Cardápio de hoje 🧑‍🍳", mediaIds)

	fmt.Println(res2)
}
