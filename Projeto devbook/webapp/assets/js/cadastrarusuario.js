//Cria um evento submit para o botão dentro do formulário
$('#formulario-cadastro').on('submit',cadastrarusuario)

function cadastrarusuario(evento) {
    //Não permite o comportamento padrão do submit, ou seja, limpar o formulário automaticamente.
    //ao executar o preventdefault, o formulário permanece preenchido mesmo após o submit.
    evento.preventDefault(); 

    senha = $('#senha').val()
    confirmarSenha = $('#confirmar-senha').val()
    if (senha != confirmarSenha) {
        alert("As senhas não coincidem.")
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
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro) {
        console.log(erro)
        alert("Erro ao cadastrar usuário: "+erro);
    });

    //Habilitar a linha de baixo, permite visualizar no console os dados enviados
    //console.log($('#nome').val()+' - '+$('#email').val()+' - '+$('#nick').val()+' - '+$('#senha').val())
}