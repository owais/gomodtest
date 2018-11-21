package pkgb // import github.com/owais/gomodtest/pkgb

import k "github.com/GuilhermeCaruso/Kair"

func Get() string {
	return k.Now().CustomFormat("dd/MM/YYYY hh:mm:ss")
}
