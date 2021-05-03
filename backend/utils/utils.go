package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"../config"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo"
	"github.com/leekchan/accounting"
)

type Fecha time.Time

func TimeIn(t time.Time) time.Time {
	loc, err := time.LoadLocation("America/Buenos_Aires")
	if err == nil {
		t = t.In(loc)
	}
	return t
}
func BodyToJson(c echo.Context) (map[string]interface{}, error) {
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	var body map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &body); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", body)
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	return body, nil
}

func GetUserId(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id, _ := strconv.ParseUint(claims["idusuario"].(string), 10, 32)

	return uint(id)
}

func GetUserIdTipo(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id, _ := strconv.ParseUint(claims["idusuariotipo"].(string), 10, 32)

	return uint(id)
}

func ParseInt(str string) uint {
	id, _ := strconv.ParseUint(str, 10, 32)

	return uint(id)
}
func ParseInt2(str string) int {
	id, _ := strconv.Atoi(str)
	return int(id)
}
func GetPath() string {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

const (
	cnGofpdfDir = "."
	cnFontDir   = cnGofpdfDir + "/font"
	ImgDir      = cnGofpdfDir + "/image"
	cnTextDir   = cnGofpdfDir + "/text"
)

func PadLeft(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
func PadRight(str, pad string, lenght int) string {
	for {
		str = pad + str
		if len(str) > lenght {
			return str[len(str)-lenght : len(str)]
		}
	}
}
func Index(vs []uint, t uint) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

var (
	//ErrValorNoAdmitido error para valor no admitidos
	ErrValorNoAdmitido = errors.New("Valor no admitido")
	us                 = []string{"cero", "uno", "dos", "tres", "cuatro", "cinco", "seis", "siete", "ocho", "nueve"}
	ds                 = []string{"X", "y", "veinte", "treinta", "cuarenta", "cincuenta", "sesenta", "setenta", "ochenta", "noventa"}
	des                = []string{"diez", "once", "doce", "trece", "catorce", "quince", "dieciseis", "diecisiete", "dieciocho", "diecinueve"}
	cs                 = []string{"x", "cien", "doscientos", "trescientos", "cuatrocientos", "quinientos", "seiscientos", "setecientos", "ochocientos", "novecientos"}
)

//IntLetra convierte un numero entero a su representacion en palabras
func IntLetra(n int) (s string, err error) {
	var aux string
	sb := strings.Builder{}
	if n < 0 {
		sb.WriteString("menos")
		n *= -1
	}
	millones := quotient(n, 1000000)
	if millones > 999999 {
		return s, ErrValorNoAdmitido
	}
	millares := quotient(n, 1000) % 1000
	centenas := quotient(n, 100) % 10
	decenas := quotient(n, 10) % 10
	unidades := n % 10
	/* fmt.Println("millones", millones)
	fmt.Println("millares", millares)
	fmt.Println("centenas", centenas)
	fmt.Println("decenas", decenas)
	fmt.Println("unidades", unidades)
	return */
	if millones == 1 {
		sb.WriteString(" un millón")
	} else if millones > 1 {
		aux, err = IntLetra(millones)
		sb.WriteString(aux)
		sb.WriteString(" millones")
	}
	if millares == 1 {
		sb.WriteString(" mil")
	} else if millares > 1 {
		aux, err = IntLetra(millares)
		sb.WriteRune(' ')
		sb.WriteString(aux)
		sb.WriteString(" mil")
	}
	if centenas == 1 {
		if n%100 == 0 {
			sb.WriteString(" cien")
		} else {
			sb.WriteString(" ciento")
		}
	} else if centenas > 0 {
		sb.WriteRune(' ')
		sb.WriteString(cs[centenas])
	}
	if decenas == 1 {
		sb.WriteRune(' ')
		sb.WriteString(des[n%10])
		unidades = 0
	} else if decenas == 2 && unidades > 0 {
		sb.WriteString(" veinti")
		sb.WriteString(us[unidades])
		unidades = 0
	} else if decenas > 1 {
		sb.WriteRune(' ')
		sb.WriteString(ds[decenas])
		if unidades > 0 {
			sb.WriteString(" y")
		}
	}
	if unidades > 0 {
		sb.WriteRune(' ')
		sb.WriteString(us[unidades])
	} else if n == 0 {
		sb.WriteString(us[0])
	}
	return strings.TrimSpace(sb.String()), err
}

func quotient(numerator, denominator int) int {
	return numerator / denominator
}

func Numeroaletra(num float64) string {

	entero := int(num)

	e, err := IntLetra(entero)
	if err != nil {
		return "error..."
	}
	decimal := int((num - float64(int(num))) * 100)
	d, errr := IntLetra(decimal)
	if errr != nil {
		return "error..."
	}
	return e + " con " + d
}

func CellFormat(pdf *gofpdf.Fpdf, maxline int, w float64, h float64, txt string, border string, ln int, align string, fill bool, link int, linkStr string) string {

	lines := pdf.SplitLines([]byte(txt), w)

	i := 0
	resto := ""
	for _, line := range lines {
		i = i + 1
		if i <= maxline {
			//		pdf.MultiCell(w, h, string(line),border,align,fill)
			pdf.CellFormat(w, h, string(line), border, 2, align, fill, link, linkStr)
		} else {
			resto = resto + string(line)
		}

	}
	return resto
}
func MultiCell(pdf *gofpdf.Fpdf, w float64, h float64, txt string, border string, align string, fill bool, maxline int) string {

	lines := pdf.SplitLines([]byte(txt), w)

	i := 0
	resto := ""
	for _, line := range lines {
		i = i + 1
		if i <= maxline {
			//		pdf.MultiCell(w, h, string(line),border,align,fill)
			pdf.CellFormat(w, h, string(line), border, 0, align, fill, 0, "")
		} else {
			resto = resto + string(line)
		}

	}
	return resto
}
func Substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func Refresh(f uint) {
	host:= config.UrlSocket
	url := fmt.Sprintf("http://"+host+"/ws/franquicias_%v", f)

	resp, err := http.Post(url, "application/json", nil)
	fmt.Println("'err:'", err)

		if resp != nil {
			defer resp.Body.Close()

		}
}

func ParseDate(t time.Time) time.Time {
	year := t.Year()
	month := t.Month()
	day := t.Day()
	hour := 0
	minute := 0
	second := 0

	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
}

func Money(numero float64) string {
	ac := accounting.Accounting{
		Symbol:    "", //El símbolo
		Precision: 2,   // ¿Cuántos "centavos" queremos? (también llamado precisión)
		Thousand:  ",", //Separador de miles
		Decimal:   ".", //Separador de decimales

	}
	numeroComoDinero := ac.FormatMoney(numero)
	return numeroComoDinero
}