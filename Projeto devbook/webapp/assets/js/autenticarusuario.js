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
        window.location = "/home";
    }).fail(function(erro) {
        console.log(erro)
        msgErro = erro.responseJSON.erro;
        Swal.fire({
            title: "Erro",
            text: msgErro,
            icon: "error"
        })
    });
}