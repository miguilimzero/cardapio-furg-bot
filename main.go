package main

import "fmt"

func makePrintAndUpload(url string, filename string) string {
	success := takeScreenshot(filename, url)

	if !success {
		return ""
	}

	res := uploadMedia(filename)

	return res["media_id_string"].(string)
}

func main() {
	var res map[string]interface{}

	var mediaLago string
	var mediaCC string
	var mediaCCMAR string

	var mediaIds []string

	mediaLago = makePrintAndUpload("https://www.furg.br/estudantes/cardapio-ru/restaurante-universitario-lago", "stubs/ru-lago.png")
	if mediaLago != "" {
		mediaIds = append(mediaIds, mediaLago)
	}

	mediaCC = makePrintAndUpload("https://www.furg.br/estudantes/cardapio-ru/restaurante-universitario-cc", "stubs/ru-cc.png")
	if mediaCC != "" {
		mediaIds = append(mediaIds, mediaCC)
	}

	mediaCCMAR = makePrintAndUpload("https://www.furg.br/estudantes/cardapio-ru/restaurante-universitario-ccmar", "stubs/ru-ccmar.png")
	if mediaCCMAR != "" {
		mediaIds = append(mediaIds, mediaCCMAR)
	}

	if len(mediaIds) == 0 {
		res = sendTweet("Parece que não temos cardápio hoje 😔", mediaIds) // Sending an empty array of mediaIds, since its always needed by the method
	} else {
		res = sendTweet("Cardápio de hoje 🧑‍🍳", mediaIds)
	}
	
	fmt.Println(res)
}
