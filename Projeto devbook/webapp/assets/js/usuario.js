$('#parar-seguir').on('click',pararSeguir)
$('#seguir').on('click',seguir)


function pararSeguir() {   
    //Obtém a propriedade usuario-id
    const usuarioId = $(this).data('usuario-id');

    $(this).prop('disabled',true);
    $.ajax({
        url: `/publicacoes/${publicacaoId}/parar-de-seguir`,
        method: "POST"
    }).done(function() {
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops..","Erro ao parar de seguir usuário","error");

    }).always(function() {
        $(this).prop('disabled',false);
    })
}

function seguir() {
    
}