package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"
)

// command : go run . sample.txt result.txt
func main() {
	arg := os.Args[1:]
	if len(arg) > 2 || len(arg) < 2 {
		return
	}
	// Nom fichier d'entrée
	file1 := arg[0]
	// Nom fichier de sortie
	file2 := arg[1]
	// fmt.Println("file 1 : " + file1)
	// fmt.Println("file 2 : " + file2)

	// lecture du fichier d'entrée
	data, err := os.ReadFile(file1)
	if err != nil {
		log.Fatal(err)
	}

	// création du fichier de sortie
	file, err := os.Create(file2)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Tranformation de l'apostrophe (4 bytes) utilisé en quote (1 byte)
	data = apostrophe(string(data))

	// fmt.Printf("Contenu d'origine du fichier : ")
	// fmt.Printf(string(data))

	// fmt.Printf("Transform (hex) : ")
	// fmt.Println(transformhex("1a")) // 26
	// fmt.Printf("Transform (bin) : ")
	// fmt.Println(transformbin(10))
	// fmt.Printf("Transform (up) : ")
	// fmt.Println(ToUpper("toto"))
	// fmt.Printf("Transform (low) : ")
	// fmt.Println(ToLower("TOTO"))
	// fmt.Printf("Transform (cap) : ")
	// fmt.Println(Capitalize("bridge"))

	datatemp := []byte(data)
	var datares []byte
	for a := len(datatemp) - 1; a >= 0; a-- {
		temp := ""
		counta := 0
		if datatemp[a] == byte(')') {
			for b := a; datatemp[b] != byte('('); b-- {
				temp = temp + string(datatemp[b])
				counta++
			}
		}
		a = a - counta

		// si (low, <number>), extraction du number
		var numbers string
		var instance string
		number := 0
		for c, x := range temp {
			if x == ',' {
				numbers = temp[1:]
				instance = temp[c+1:]
				for _, y := range numbers {
					if y >= '0' && y <= '9' {
						number = number*10 + int(y-'0')
					}
				}
			}
		}

		countb := 0
		var temp2 []string
		// traitement de (hex)
		if temp == ")xeh" {
			//traitement 1er mot
			for b := a - 2; datatemp[b] != byte(' '); b-- {
				temp2 = append(temp2, string(datatemp[b]))
			}

			var mothex string
			// inversion du mot par décrémentation
			for b := len(temp2) - 1; b >= 0; b-- {
				mothex = mothex + temp2[b]
				countb++
			}

			// application de la fonction du hex
			intmothex := transformhex(mothex)
			// injection dans ma variable final de la transformation
			for intmothex > 0 {
				datares = append(datares, byte(int(intmothex%10+'0')))
				intmothex = intmothex / 10
			}
			countb++
		} else if temp == ")nib" { // traitement de (bin)

			//traitement 1er mot
			for b := a - 2; datatemp[b] != byte(' '); b-- {
				temp2 = append(temp2, string(datatemp[b]))
			}

			var motint string
			// inversion du mot
			for b := len(temp2) - 1; b >= 0; b-- {
				motint = motint + temp2[b]
				countb++
			}
			num := 0
			for _, x := range motint {
				num = num*10 + int(x-'0')
			}

			// application de la fonction du bin
			numint := transformbin(num)

			// int en byte
			for numint > 0 {
				datares = append(datares, byte(numint%10+'0'))
				numint = numint / 10
			}
			countb++
			// Transformation low, up et cap
			// number est la variable qui contiens le nombre de mot a transformer
		} else if instance == "wol" || temp == ")wol" { // traitement de (low)
			nbespace := 0
			if number == 0 {
				number++
			}

			// Récupération du ou des mot(s) a lower
			for b := a - 2; number > nbespace && b >= 0; b-- {
				temp2 = append(temp2, string(datatemp[b]))
				if datatemp[b] == ' ' {
					nbespace++
				}
			}

			var temprev string
			for b := len(temp2) - 1; b >= 0; b-- {
				temprev = temprev + temp2[b]
			}

			// Transformation en (low)
			transformlow := ToLower(temprev)
			// injection de la transformation dans le mots
			for b := len(transformlow) - 1; b >= 0; b-- {
				datares = append(datares, byte(transformlow[b]))
				countb++
			}
			countb++

		} else if instance == "pu" || temp == ")pu" { // traitement de (up)
			nbespace := 0
			if number == 0 {
				number++
			}

			// Récupération du ou des mot(s) a upper
			for b := a - 2; number > nbespace && b >= 0; b-- {
				temp2 = append(temp2, string(datatemp[b]))
				if datatemp[b] == ' ' {
					nbespace++
				}
			}

			var temprev string
			for b := len(temp2) - 1; b >= 0; b-- {
				temprev = temprev + temp2[b]
			}

			// Transformation en (low)
			transformup := ToUpper(temprev)
			// injection de la transformation dans le mots
			for b := len(transformup) - 1; b >= 0; b-- {
				datares = append(datares, byte(transformup[b]))
				countb++
			}
			countb++

		} else if instance == "pac" || temp == ")pac" { // traitement de (cap)
			nbespace := 0
			if number == 0 {
				number++
			}

			// Récupération du ou des mot(s) a Capitalisé
			for b := a - 2; number > nbespace && b >= 0; b-- {
				temp2 = append(temp2, string(datatemp[b]))
				if datatemp[b] == ' ' {
					nbespace++
				}
			}

			var temprev string
			for b := len(temp2) - 1; b >= 0; b-- {
				temprev = temprev + temp2[b]
			}

			// Transformation en (low)
			transformCap := Capitalize(temprev)
			// injection de la transformation dans le mots
			for b := len(transformCap) - 1; b >= 0; b-- {
				datares = append(datares, byte(transformCap[b]))
				countb++
			}
			countb++

		} else {
			datares = append(datares, datatemp[a])
		}
		a = a - countb
	}

	// gestion des espaces avant la ponctuations
	datares = espaceavponc(datares)

	fmt.Println(string(datares))
	// inversion de datares pour le retour final dans le bon sens
	for a, b := 0, len(datares)-1; a < b; a, b = a+1, b-1 {
		datares[a], datares[b] = datares[b], datares[a]
	}
	fmt.Println(string(datares))
	// gestion des espaces aprés la ponctuations
	datares = espaceafterponc(datares)

	// gestion des espaces pour les quotes
	datares = quote(datares)

	// fmt.Printf("Contenu avant voyelle : ")
	// fmt.Println(string(datares))

	// gestion des voyelles a transformer
	datares = voyelle(datares)

	// ecrire dans le fichier de sortie
	erreur := os.WriteFile(file2, datares, 0666)
	if erreur != nil {
		log.Fatal(erreur)
	}

	// fmt.Printf("Contenu final du fichier : ")
	// fmt.Println(string(datares))
}

// transformation occurance (hex) (ex : 42 deviens 66)
func transformhex(s string) int {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Probléme fonction hex")
	}
	return int(decoded[0])
}

// transformation occurance (bin) (ex : 10 deviens 2)
func transformbin(s int) int {
	var numbyteinv []byte
	// injection de l'int dans le tableau
	for x := 0; s > 0; x++ {
		numbyteinv = append(numbyteinv, byte(s%10))
		s = s / 10
	}

	taille := len(numbyteinv)
	res := make([]byte, 8)
	// injection de l'int dans le array de 8 élément (byte = 8 éléments de 0 et de 1)
	for a := 0; a < taille; a++ {
		res[a] = numbyteinv[a]
	}
	// inversion de l'int pour le mettre a l'endroit l'int d'entrée
	res[0], res[1], res[2], res[3], res[4], res[5], res[6], res[7] = res[7], res[6], res[5], res[4], res[3], res[2], res[1], res[0]

	var resfinal byte
	for a := 0; a < 8; a++ {
		// injection du byte dans une variable
		// binaire donc x2 est une révolution
		resfinal = resfinal*2 + res[a]
	}
	return int(resfinal)
}

// transformation occurance (up) (ex : toto deviens TOTO)
func ToUpper(s string) []rune {
	mots := []rune(s)
	for i := 0; i < len(s); i++ {
		if mots[i] >= 'a' && mots[i] <= 'z' {
			mots[i] = mots[i] - 32
		}
	}
	return mots
}

// transformation occurance (low) (ex : TOTO deviens toto)
func ToLower(s string) []rune {
	mots := []rune(s)
	for i := 0; i < len(s); i++ {
		if mots[i] >= 'A' && mots[i] <= 'Z' {
			mots[i] = mots[i] + 32
		}
	}
	return mots
}

// transformation occurance (cap) (ex : bridge deviens Bridge)
func Capitalize(s string) []rune {
	// Transformation de tout les caractéres  majuscule en minuscule
	mots := []rune(s)
	for i := 0; i < len(s); i++ {
		if mots[i] >= 'A' && mots[i] <= 'Z' {
			mots[i] = mots[i] + 32
		}
	}

	// transformation des caractéres minuscule derriére un caractére non alphanumérique en majuscule
	for i := 0; i < len(s)-1; i++ {
		if (mots[i] < '0' || mots[i] > 'z') || (mots[i] > '9' && mots[i] < 'A') || (mots[i] > 'Z' && mots[i] < 'a') {
			if mots[i+1] >= 'a' && mots[i+1] <= 'z' {
				mots[i+1] = mots[i+1] - 32
			}
		}
	}

	// gestion de la 1ére lettre si une minuscule
	if mots[0] >= 'a' && mots[0] <= 'z' {
		mots[0] = mots[0] - 32
	}

	return mots
}

// gestion des espaces avant la ponctuations
func espaceavponc(datares []byte) []byte {
	var datatemp2 []byte
	for a := 0; a < len(datares); a++ {
		if datares[a] == byte('.') || datares[a] == byte(',') || datares[a] == byte('!') || datares[a] == byte('?') || datares[a] == byte(':') || datares[a] == byte(';') {
			// retire l'espace apres
			datatemp2 = append(datatemp2, datares[a])
			if datares[a+1] == byte(' ') {
				a++
			}
		} else {
			datatemp2 = append(datatemp2, datares[a])
		}
	}
	return datatemp2
}

// gestion des espaces aprés la ponctuations
func espaceafterponc(datares []byte) []byte {
	var datatemp3 []byte
	for a := 0; a < len(datares); a++ {
		if (datares[a] == byte('.') || datares[a] == byte(',') || datares[a] == byte('!') || datares[a] == byte('?') || datares[a] == byte(':') || datares[a] == byte(';')) && a < len(datares)-1 {
			datatemp3 = append(datatemp3, datares[a])
			// ajouter un espace apres
			if ((datares[a+1] == byte(' ')) || datares[a+1] == byte('.') || datares[a+1] == byte(',') || datares[a+1] == byte('!') || datares[a+1] == byte('?') || datares[a+1] == byte(':') || datares[a+1] == byte(';')) && a < len(datares)-1 {
				// ne rien faire
			} else {
				datatemp3 = append(datatemp3, byte(' '))
			}
		} else {
			datatemp3 = append(datatemp3, datares[a])
		}
	}
	return datatemp3
}

// Tranformation de l'apostrophe (4 bytes) utilisé en quote (1 byte)
func apostrophe(data string) []byte {
	tempapos := strings.ReplaceAll(string(data), "‘", "")
	var tempaposbyte []byte
	for _, x := range tempapos {
		tempaposbyte = append(tempaposbyte, byte(x))
	}
	return tempaposbyte
}

// gestion des espaces pour les quotes
func quote(datares []byte) []byte {
	var dataquote []byte
	countquote := 0
	for a := 0; a < len(datares); a++ {
		if datares[a] == byte('\'') && countquote == 0 {
			countquote++
			dataquote = append(dataquote, datares[a])
			if datares[a+1] == byte(' ') {
				a++
			}
		} else if datares[a] == byte('\'') && countquote == 1 {
			if datares[a-1] == byte(' ') {
				dataquote = dataquote[0 : a-2]
				dataquote = append(dataquote, datares[a])
			} else {
				dataquote = append(dataquote, datares[a])
			}
			countquote = 0
		} else {
			dataquote = append(dataquote, datares[a])
		}
	}
	return dataquote
}

// gestion des voyelles a transformer
func voyelle(datares []byte) []byte {
	var voyelle []byte
	for x := 0; x < len(datares)-1; x++ {
		if (datares[x] == byte('a') ||
			datares[x] == byte('e') ||
			datares[x] == byte('i') ||
			datares[x] == byte('o') ||
			datares[x] == byte('u')) &&
			datares[x+1] == byte(' ') &&
			datares[x-1] == byte(' ') &&
			x < len(datares)-1 &&
			x > 1 {
			if (datares[x+2] == byte('a') ||
				datares[x+2] == byte('e') ||
				datares[x+2] == byte('i') ||
				datares[x+2] == byte('o') ||
				datares[x+2] == byte('u') ||
				datares[x+2] == byte('h')) &&
				x < len(datares)-2 {
				voyelle = append(voyelle, datares[x])
				voyelle = append(voyelle, byte('n'))
			} else {
				voyelle = append(voyelle, datares[x])
			}
		} else {
			voyelle = append(voyelle, datares[x])
		}
	}
	voyelle = append(voyelle, datares[len(datares)-1])
	return voyelle
}
