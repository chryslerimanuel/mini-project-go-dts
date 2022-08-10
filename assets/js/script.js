$(document).ready(function () {

    $('#modalDelUser-btnDel').click(function () {
        var id = $('#modalDelId').text();
        window.location.href = ('/users/delete?id=' + id)
    });

});

function modalDelUser(id) {
    $('#modalDelUser').modal('show');
    $('#modalDelId').text(id);
}

