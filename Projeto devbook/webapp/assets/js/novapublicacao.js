$('#nova-publicacao').on('submit',novapublicacao)


function novapublicacao(evento) {
    evento.preventDefault(); 

 
    //Executar a requisição para o webapp
    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function() {
        window.location = "/home";
    }).fail(function(erro) {
        console.log(erro)
        alert("Erro ao tentar publicar: "+erro);
    });

}