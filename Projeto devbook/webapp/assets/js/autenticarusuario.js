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
    }).done(function() {
        window.location = "/home";
        alert("Usuário autenticado com sucesso.")
    }).fail(function(erro) {
        console.log(erro)
        alert("Usuário ou senha inválidos");
    });
}