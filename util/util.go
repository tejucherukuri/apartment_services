package util

import (
	//"crypto/aes"
	//"crypto/cipher"
	"bytes"
	"encoding/base64"
	"regexp"
	"strconv"
	"time"
	"math"
)

const Yyyymmddform string = "2006-01-02"
const Ddmmyyyyform string = "02-01-2006"
const DDmmyyyyformWithSlashes string = "02/01/2006"
const Ustimeform string = "01/02/2006 15:04:05 MST"
const Indtimeform string = "02/01/2006 15:04:05 MST"
const Usdateform string = "01/02/2006"
const MysqlTimeForm string = "2006-01-02 15:04:05"
const EMAIL_REGEX = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func IsValidEmail(e string) bool {
	exp, _ := regexp.Compile(EMAIL_REGEX)
	return exp.MatchString(e)
}

func CourseColorHexMap() []string {
	return []string{"4862e4", "f3aa4c", "6eb732", "2ea0ad", "af4182", "10b344", "adbb0b", "7651c1", "3abfce", "2334b1"}
}

func Decrypt(text string, key []byte, iv []byte) string {
	b, _ := base64.URLEncoding.DecodeString(text)
	return string(b)
	//TODO Some issue with this, need to revisit
	//encrypted, _ := base64.URLEncoding.DecodeString(text)
	//t := []byte(encrypted)
	//block, _ := aes.NewCipher(key)
	//decrypted := make([]byte, len(t))
	//cipher.NewCFBDecrypter(block, iv).XORKeyStream(decrypted, encrypted)
	//return string(decrypted)
}

func StringInSlice(str string, list []string) bool {
	for _, l := range list {
		if str == l {
			return true
		}
	}
	return false
}

func IntInSlice(num int, list []int) bool {
	for _, l := range list {
		if num == l {
			return true
		}
	}
	return false
}

func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,4}$`)
	return Re.MatchString(email)
}

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func PrettyInt(i int) string {
	s := strconv.Itoa(i)
	r1 := ""
	idx := 0

	// Reverse and interleave the separator.
	for i = len(s) - 1; i >= 0; i-- {
		idx++
		if idx == 4 {
			idx = 1
			r1 = r1 + ","
		}
		r1 = r1 + string(s[i])
	}

	// Reverse back and return.
	r2 := ""
	for i = len(r1) - 1; i >= 0; i-- {
		r2 = r2 + string(r1[i])
	}
	return r2
}

func PrettyFloat(f float64) string {
	s := strconv.FormatFloat(f, 'f', 2, 32)
	r1 := ""
	idx := 0
	for i := len(s) - 4; i >= 0; i-- {
		idx++
		if idx == 4 {
			idx = 1
			r1 = r1 + ","
		}
		r1 = r1 + string(s[i])
	}
	r2 := ""
	for i := len(r1) - 1; i >= 0; i-- {
		r2 = r2 + string(r1[i])
	}
	r2 = r2 + s[len(s)-3:]
	return r2
}

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div > roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
