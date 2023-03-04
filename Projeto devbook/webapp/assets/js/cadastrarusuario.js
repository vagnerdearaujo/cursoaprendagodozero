//Cria um evento submit para o botão dentro do formulário
$('#formulario-cadastro').on('submit',cadastrarusuario)

function cadastrarusuario(evento) {
    //Não permite o comportamento padrão do submit, ou seja, limpar o formulário automaticamente.
    //ao executar o preventdefault, o formulário permanece preenchido mesmo após o submit.
    evento.preventDefault(); 

    senha = $('#senha').val()
    confirmarSenha = $('#confirmar-senha').val()
    if (senha != confirmarSenha) {
        Swal.fire({
            title: "Cadastramento de Usuário",
            text: "As senhas não coincidem.",
            icon: "error"
        })
        return
    }

    //Executar a requisição para o webapp
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function() {
        Swal.fire({
            title: "Cadastramento de Usuário",
            text: "Usuário cadastrado com sucesso!",
            icon: "success"
        }).then(function(){
            //Se cadastrou com sucesso, loga automaticamente passando as credenciais do usuário.
            $.ajax({
                url: "/login",
                method: "POST",
                data: {
                    email: $('#email').val(),
                    senha: $('#senha').val()
                }
            }).done(function(){
                window.location = "/home";
            }).fail(function(){
                Swal.fire("Falha no login","Erro ao autenticar o usuário","error");
            })
        })

    }).fail(function(erro) {
        console.log(erro)
        Swal.fire({
            title: "Cadastramento de Usuário",
            text: "Erro ao cadastrar usuário: "+erro.responseJSON.erro,
            icon: "error"
        })
    });

    //Habilitar a linha de baixo, permite visualizar no console os dados enviados
    //console.log($('#nome').val()+' - '+$('#email').val()+' - '+$('#nick').val()+' - '+$('#senha').val())
}