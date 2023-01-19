package servidor

import (
	"crud/banco"
	"crud/settings"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Mapeia a tabela usuarios do banco para o struct usuario
type usuarios struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario Inclui um novo usuário no banco.
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	/*
		Os dados tratados neste método são oriundos diretamente do elemento body
		da página html.
		Os dados estão em formato json e o método ReadAll irá extrair estes dados
	*/
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Dados incorretos enviados para persistência."))

		//O return em um método sem retorno declarado, força seu encerramento
		return
	}

	var usuario usuarios

	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter os dados para persistência"))
		return
	}

	db, erro := banco.ConectarDB(settings.MySQLSettings())

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	fmt.Println("Conexão realizada com sucesso")

	defer db.Close()

	/*
		Para realizar a inclusão o mais óbvio seria o comando:
		insert into usuarios (nome, email) values ("nome","email")

		o problema desta abordagem é que ela está sujeita ao ataque sql-injection.

		para mitigar este problema existe a preparação do comando em sql.
	*/

	sqlInsert, erro := db.Prepare("insert into usuarios (nome, email) values (?,?)")
	if erro != nil {
		fmt.Println(erro)
		//w.Write([]byte("A preparação do comando de insert falhou"))
		return
	}
	defer sqlInsert.Close()
	//fmt.Println("Comando preparado")

	resultInsert, erro := sqlInsert.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Erro ao tentar incluir o usuário."))
		return
	}
	//fmt.Println("Inclusão realizada com sucesso")

	idUsuario, erro := resultInsert.LastInsertId()
	if erro != nil {
		w.Write([]byte("Erro ao obter o id do usuário inserido."))
		return
	}

	fmt.Println(idUsuario, usuario)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("<br>ID:\t %d</br>", idUsuario)))
	w.Write([]byte(fmt.Sprintf("<br>Nome:\t %v</br>", usuario.Nome)))
	w.Write([]byte(fmt.Sprintf("<br>E-Mail:\t %v</br>", usuario.Email)))
}

// ListarUsuarios Lista todos os usuários do banco
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.ConectarDB(settings.MySQLSettings())

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	defer db.Close()

	tuplas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Erro ao ler a tabela de usuarios"))
		return
	}

	defer tuplas.Close()
	//Cria um slice de usuários para conter todos os usuários recuperados do banco
	var todos_usuarios []usuarios

	//Varre todas as tuplas
	for tuplas.Next() {
		var usuario usuarios
		//O scan busca o dado dentro do recordset e os coloca nos campos
		if erro := tuplas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao recuperar os dados do usuário com scan"))
			return
		}

		//Adiciona no slice de usuários os dados obtidos via scan
		todos_usuarios = append(todos_usuarios, usuario)
	}

	//O retorno obrigatoriamente precisa ser realizado via json, porém o marshall não poderá ser utilizado neste caso
	if erro := json.NewEncoder(w).Encode(todos_usuarios); erro != nil {
		w.Write([]byte("Erro ao codificar os dados em json para devolver o resuttado"))
		return
	}
	//Retorna o status Ok para a página
	w.WriteHeader(http.StatusOK)
}

// BuscarUsuario Recupera os dados de um usuário usando seu ID
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

	/*
		Na declaração da rota foi utilizado /usuario/{id}
		Para recuperar este parâmetro será necessário inspecionar a requisição http (r)
		e o pacote mux consegue extrair os parâmetros passados na requisição
	*/

	parametros := mux.Vars(r)

	//vars cria um map[string]string, portanto o parâmetro estará como string
	/*
		ParseUint recebe
		string, base do número (neste caso decimal), tamanho em bits do retorno
	*/

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ID do usuário tem que ser um número inteiro"))
		return
	}

	db, erro := banco.ConectarDB(settings.MySQLSettings())

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	var usuario usuarios
	tupla, erro := db.Query("select * from usuarios where ID=?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao ler a tabela de usuarios"))
		return
	}
	defer tupla.Close()

	if tupla.Next() {
		if erro := tupla.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao recuperar os dados do usuário com scan"))
			return
		}
	}

	if usuario.ID != uint32(ID) {
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	//O retorno obrigatoriamente precisa ser realizado via json, porém o marshall não poderá ser utilizado neste caso
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Erro ao codificar os dados em json para devolver o resuttado"))
		return
	}
	//Retorna o status Ok para a página
	w.WriteHeader(http.StatusOK)
}

// AtualizarUsuario Atualiza os dados de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario usuarios
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("O ID do usuário deve ser um número inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Não foi possível ler os dados do usuário da requisição"))
		return
	}

	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		w.Write([]byte("Erro ao converter para json o corpo da requisição"))
		return
	}

	//Abre a conexão com o banco
	db, erro := banco.ConectarDB(settings.MySQLSettings())
	if erro != nil {
		w.Write([]byte("Erro ao conectar com o banco de dados"))
		return
	}
	defer db.Close()

	//Para qualquer operação em banco que implique em persistência de dados, deve-se utilizar o statment
	atualizaDados, erro := db.Prepare("update usuarios set nome=?,email=? where id=?")
	if erro != nil {
		w.Write([]byte("Erro ao preparar a atualização dos dados"))
		return
	}
	defer atualizaDados.Close()

	_, erro = atualizaDados.Exec(usuario.Nome, usuario.Email, ID)
	if erro != nil {
		w.Write([]byte("Erro ao executar a atualização de dados"))
		return
	}

	w.WriteHeader(http.StatusOK)

}

// ExcluirUsuario Exclui um usuário pelo ID
func ExcluirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ID do usuário tem que ser um número inteiro"))
		return
	}

	db, erro := banco.ConectarDB(settings.MySQLSettings())

	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}
	defer db.Close()

	var usuario usuarios
	tupla, erro := db.Query("select * from usuarios where ID=?", ID)
	if erro != nil {
		w.Write([]byte("Erro ao ler a tabela de usuarios"))
		return
	}
	defer tupla.Close()

	if tupla.Next() {
		if erro := tupla.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao recuperar os dados do usuário com scan"))
			return
		}
	}

	if usuario.ID != uint32(ID) {
		w.Write([]byte("Usuário não encontrado"))
		return
	}

	excluirDados, erro := db.Prepare("delete from usuarios where id=?")
	if erro != nil {
		w.Write([]byte("Erro ao recuperar os dados do usuário com scan"))
		return
	}
	defer excluirDados.Close()

	if _, erro := excluirDados.Exec(ID); erro != nil {
		w.Write([]byte("Erro ao tentar excluir os dados"))
		return
	}
	w.Write([]byte("Usuário " + usuario.Nome + " excluído com sucesso"))
	//Retorna o status Ok para a página
	w.WriteHeader(http.StatusOK)
}
