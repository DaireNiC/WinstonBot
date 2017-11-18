const form = $("#user-input-form");
const form = $("#output-area");
const input = $("#user-input");
const list = $("#conversation_list");
console.log("hello");
$("#user-input-form").submit(
function(event) {
    console.log("hello");
    event.preventDefault();
    $.get('/user-input', { value: $('#user-input').val() } )
    .done(function (data) {
        $('#output-area').val(data);
    })
});