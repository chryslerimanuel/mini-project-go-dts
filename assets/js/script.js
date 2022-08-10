$(document).ready(function () {

    $('#modalDelUser-btnDel').click(function () {
        var id = $('#modalDelId').text();
        window.location.href = ('/users/delete?id=' + id)
    });

    // Set up the Select2 control
    $('#sendToName').select2({
        placeholder: "Pilih user tujuan",
        ajax: {
            url: '/users/dropdown',
            type: 'get',
            dataType: 'json',
            delay: 1200,
            data: function (params) {
                return {
                    searchTerm: params.term // search term
                };
            },
            processResults: function (response) {
                return {
                    results: response
                };
            },
            cache: true
        }
    });

    var sendToId = $("#sendToId").val();
    if (sendToId != undefined && sendToId != "") {
        var selectUser = $('#sendToName');
        $.ajax({
            type: 'GET',
            url: '/users/view?id=' + sendToId
        }).then(function (data) {
            var option = new Option(data.text, data.id, true, true);
            selectUser.append(option).trigger('change');

            // supaya jadi selected
            selectUser.trigger({
                type: 'select2:select',
                params: {
                    data: data
                }
            });
        });
    }
});

function modalDelUser(id) {
    $('#modalDelUser').modal('show');
    $('#modalDelId').text(id);
}

