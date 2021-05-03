package iibb

import (
	"bytes"
	"fmt"
	"io"

	"mime/multipart"
	"net/http"
	"os"
	"time"
	"crypto/tls"
	"crypto/md5"
//	"path/filepath"
//	"net/http/httputil"
	"net/textproto"
	"encoding/hex"
	"strings"
	"encoding/xml"

  "golang.org/x/text/encoding/charmap"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, params map[string]string, paramName, xmlString string) (*http.Request, error) {

	bXml:=strings.NewReader(xmlString)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
  hs,_:=hash_string_md5(xmlString)
  finename := "DFEServicioConsulta_"+ strings.ToUpper(hs) +".XML"
	xmlFile, _ := CreateXML(writer, finename)

	_, err := io.Copy(xmlFile, bXml)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}



	req, err := http.NewRequest("POST", uri, body)

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, err
}



type Comporbante struct {
	XMLName xml.Name `xml:"COMPROBANTE"`
	Contribuyente Contribuyente `xml:"contribuyentes>contribuyente"`
    FechaDesde string  `xml:"fechaDesde"`
	FechaHasta string  `xml:"fechaHasta"`
}

type Contribuyente struct{
	XMLName xml.Name `xml:"contribuyente"`
	AlicuotaPercepcion    string    `xml:"alicuotaPercepcion"`
	AlicuotaRetencion     string    `xml:"alicuotaRetencion"`
}
func monthInterval(y int, m time.Month) (firstDay, lastDay time.Time) {
    firstDay = time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
    lastDay = time.Date(y, m+1, 1, 0, 0, 0, -1, time.UTC)
    return firstDay, lastDay
}
func IIBB(cuit string) (*Comporbante) {
	var (
			y int
			m time.Month
	)

	y, m, _ = time.Now().Date()
	first, last := monthInterval(y, m)


	q := new(Comporbante)
	extraParams := map[string]string{
		"user":"30708619953",
		"password":"072015"}

  xmlString:=`<CONSULTA-ALICUOTA>
<fechaDesde>`+first.Format("20060102")+`</fechaDesde>
<fechaHasta>`+last.Format("20060102")+`</fechaHasta>
<cantidadContribuyentes>1</cantidadContribuyentes>
<contribuyentes class="list">
<contribuyente>
<cuitContribuyente>`+cuit+`</cuitContribuyente>
</contribuyente>
</contribuyentes>
</CONSULTA-ALICUOTA>`

	request, err := newfileUploadRequest("https://dfe.arba.gov.ar/DomicilioElectronico/SeguridadCliente/dfeServicioConsulta.do", extraParams, "file", xmlString)

if err != nil {
		return q
	}

	tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr,Timeout: time.Second * 3}
	//byts, _ := httputil.DumpRequest(request, true)
	//fmt.Println(string(byts))


	resp, err := client.Do(request)

	if err != nil {

		return q
	} else {

		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			return q
		}

    resp.Body.Close()

		if(resp.StatusCode==200){

			parseXML(body.Bytes(), &q)
      return q
		}


	}
	return q
}


func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
    if charset == "ISO-8859-1" {
        // Windows-1252 is a superset of ISO-8859-1, so should do here
        return charmap.Windows1252.NewDecoder().Reader(input), nil
    }
    return nil, fmt.Errorf("Unknown charset: %s", charset)
}
func parseXML(xmlDoc []byte, target interface{}) {
    reader := bytes.NewReader(xmlDoc)
    decoder := xml.NewDecoder(reader)
		decoder.CharsetReader = makeCharsetReader
    // Fixes "xml: encoding \"windows-1252\" declared but Decoder.CharsetReader is nil"
  //  decoder.CharsetReader = charset.NewReader
    if err := decoder.Decode(target); err != nil {
        //log.Fatalf("unable to parse XML '%s':\n%s", err, xmlDoc)
    }
}

func CreateXML(w *multipart.Writer, filename string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", "text/xml")
	return w.CreatePart(h)
}


func hash_file_md5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}


func hash_string_md5(filePath string) (string, error) {
	var returnMD5String string
hash := md5.New()
io.WriteString(hash, filePath)

hashInBytes := hash.Sum(nil)[:16]

//Convert the bytes to a string
returnMD5String = hex.EncodeToString(hashInBytes)

return returnMD5String, nil
}
