package conf

import "os"

var (
	p  = ":3000"
	bp = "files/"
	fn = "address.csv"
	au = "https://appreactivacioneconomica.tlajomulco.gob.mx/api/georeferencia/queryGeoreferencia"
)

func ConfigInit() {
	os.Setenv("p", p)
	os.Setenv("bp", bp)
	os.Setenv("fn", fn)
	os.Setenv("au", au)
}
