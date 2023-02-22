$('#form-login').on('submit',autenticarusuario)

function autenticarusuario(evento) {
    evento.preventDefault(); 

    senha = $('#senha').val()
    email = $('#email').val()
    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            "email":email,
            "senha":senha
        }
    }).done(function(resposta) {
        //console.log(resposta.responseJSON)
        window.location = "/home";
        alert("Usuário autenticado com sucesso.")
    }).fail(function(erro) {
        console.log(erro)
        alert(erro.responseJSON.erro);
    });
}