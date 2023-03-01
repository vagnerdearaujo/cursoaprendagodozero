// # Pesquisa um ID 
$('#nova-publicacao').on('submit',novapublicacao)

//. Pesquisa uma classe
$('.curtir-publicacao').on('click',curtirPublicacao)

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

function curtirPublicacao(evento) {
    evento.preventDefault();
    //Captura o elemento clicado dentro do evento
    const elementoClicado = $(evento.target);
    
    //Obtém a propriedade publicacao-id, declarada como data-pubicacao-id
    const publicacaoId = elementoClicado.data('publicacao-id');
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"        
    }).done(function() {
        const contadorCurtidas = elementoClicado.prev('span');
        const qtdeCurtidas = parseInt(contadorCurtidas.text());
        contadorCurtidas.text(qtdeCurtidas+1)
    }).fail(function(erro) {
        console.log(erro)
        alert('Erro ao tentar curtir');
    })

}