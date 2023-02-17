package auxiliar

import (
	"fmt"

	"github.com/badoux/checkmail"
)

// Recebe um e-mail e valida
func Validar_email() {
	fmt.Println(checkmail.ValidateFormat("texte//oier.ere"))
}
