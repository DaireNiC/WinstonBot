//const form = $("#user-input");
$("#user-input-form").submit(
function(event) {
    event.preventDefault();
    $.get('/user-input', { value: $('#user-input').val() } )
    .done(function (data) {
        $('#output-area').val(data);
    })
});