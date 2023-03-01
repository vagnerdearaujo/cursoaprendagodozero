// # Pesquisa um ID 
$('#nova-publicacao').on('submit',novapublicacao)

/*
    Por padrão a chamada ao evento submit deveria estar atrelado diretamente à classe
    $('.curtir-publicacao').on('click',curtirPublicacao)

    Neste caso haverá uma mudança dinâmica na classe de forma que o evento deve ser "pescado" diretamente no DOM
*/

//. Assinala os eventos da página pesquisando dentro do DOM
$(document).on('click','.curtir-publicacao',curtirPublicacao)
$(document).on('click','.descurtir-publicacao',descurtirPublicacao)

$('#atualizar-publicacao').on('click',editarPublicacao)


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

    //Se o elemento de curtir for chamado muitas vezes em um curto período de tempo, pode gerar
    // problemas de muitas requisições ao mesmo tempo.
    // Para prevenir este problema o elemento será bloqueado e desbloqueado ao final da chamada.
    elementoClicado.prop('disabled',true)
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"        
    }).done(function() {
        const contadorCurtidas = elementoClicado.prev('span');
        const qtdeCurtidas = parseInt(contadorCurtidas.text());

        //Cria e remove dinamicamente classes ao documento
        elementoClicado.addClass("descurtir-publicacao")
        elementoClicado.addClass("text-danger")
        elementoClicado.removeClass("curtir-publicacao")
        contadorCurtidas.text(qtdeCurtidas+1)
    }).fail(function(erro) {
        console.log(erro)
        alert('Erro ao tentar curtir');
    }).always(function(){
        //always é chamado independetemente de ter havido sucesso ou falha na requisição ajax
        elementoClicado.prop('disabled',false)
    })

}

function descurtirPublicacao(evento) {
    evento.preventDefault();
    //Captura o elemento clicado dentro do evento
    const elementoClicado = $(evento.target);
    
    //Obtém a propriedade publicacao-id, declarada como data-pubicacao-id
    const publicacaoId = elementoClicado.data('publicacao-id');

    //Se o elemento de curtir for chamado muitas vezes em um curto período de tempo, pode gerar
    // problemas de muitas requisições ao mesmo tempo.
    // Para prevenir este problema o elemento será bloqueado e desbloqueado ao final da chamada.
    elementoClicado.prop('disabled',true)
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"        
    }).done(function() {
        const contadorCurtidas = elementoClicado.prev('span');
        const qtdeCurtidas = parseInt(contadorCurtidas.text());
        //Cria e remove dinamicamente classes ao documento
        elementoClicado.addClass("curtir-publicacao")
        elementoClicado.removeClass("text-danger")
        elementoClicado.removeClass("descurtir-publicacao")
        contadorCurtidas.text(qtdeCurtidas-1)
    }).fail(function(erro) {
        console.log(erro)
        alert('Erro ao tentar descurtir');
    }).always(function(){
        //always é chamado independetemente de ter havido sucesso ou falha na requisição ajax
        elementoClicado.prop('disabled',false)
    })

}

function editarPublicacao(evento) {
    botaoPublicacao = $(this)
    botaoPublicacao.prop('disabled',true)
    const publicacaoId = $(this).data('publicacao-id')
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val()
        }
    }).done(function(resultado){
        alert('Publicação atualizada com sucesso')

    }).fail(function(erro){
        alert('Falha na Publicação.')

    }).always(function(){
        botaoPublicacao.prop('disabled',false)
    })
}