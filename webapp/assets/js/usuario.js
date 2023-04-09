
$('#parar-de-seguir').on('click', pararDeSeguir);
$('#seguir').on('click', seguir);
$('#editar-usuario').on('submit', editar);
$('#atualizar-senha').on('submit', atualizarSenha);
$('#deletar-usuario').on('click', deletarUsuario);

function pararDeSeguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true)

    $.ajax({
        url: `/usuarios/${usuarioId}/parar-de-seguir`,
        method: "POST"
    }).done(function(){
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao parar de seguir o usuário", "error");
        $('#parar-de-seguir').prop('disabled', false)
    });
}

function seguir() {
    const usuarioId = $(this).data('usuario-id');
    $(this).prop('disabled', true)

    $.ajax({
        url: `/usuarios/${usuarioId}/seguir`,
        method: "POST"
    }).done(function(){
        window.location = `/usuarios/${usuarioId}`;
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao seguir o usuário", "error");
        $('#seguir').prop('disabled', false)
    });
}

function editar(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/editar-usuario",
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
        }
    }).done(function(){
        Swal.fire(
            "Sucesso!",
            "Usuario atualizado com sucesso",
            "sucess"
        ).then(function(){
            window.location = "/perfil";
        });
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao atualizar o usuario", "error");
    })
}

function atualizarSenha(evento) {
    evento.preventDefault();

    if ($('#nova-senha').val() != $('#confirmar-senha').val()) {
        Swal.fire("Ops...", "As senhas não coincidem!", "warning");
        return;
    }

    $.ajax({
        url: "/atualizar-senha",
        method: "POST",
        data: {
            atual: $('#senha-atual').val(),
            nova: $('#nova-senha').val()
        }
    }).done(function() {
        Swal.fire("Sucesso!", "A senha foi atualizada com sucesso!", "success")
            .then(function() {
                window.location = "/perfil";
            })
    }).fail(function() {
        Swal.fire("Ops...", "Erro ao atualizar a senha!", "error");
    });
}

function deletarUsuario(evento) {
    Swal.fire({
        title: "Atenção",
        title: "Tem certeza que deseja apagar sua conta? Essa ação é irreversivel",
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confimacao){
        if (confimacao.value){
            $.ajax({
                url: "deletar-usuario",
                method: "Delete",
            }).done(function(){
                Swal.fire("Sucesso!", "Seu usuario foi excluido com sucesso", "sucess")
                .then(function(){
                    window.location = "/logout";
                })
            }).fail(function(){
                Swal.fire("Ops...", "Ocorreu um erro ao excluir o usuario", "error")
            });
        }
    })

  
}