package users

import (
	"cursoGo/GoPractica/modelos"
	"fmt"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
