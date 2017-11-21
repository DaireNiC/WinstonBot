$("#user-input-form").submit(
    function(event) {
        //render the input inthe chat window
        insertChat("me", $('#user-input').val());

        //AJAX, don't want page to refresh - instead we just change the content
        event.preventDefault();

        $.get('/user-input', {
                value: $('#user-input').val()
            })
            .done(function(data) {
                insertChat("winston", data);
                $('#user-input').val(" ") //clear the  inout field
            })
    });

//get a random greeting only once when the page loads
$(document).ready(function() {
    $.get('/chat-session') //call this handler in the .go file
        .done(function(data) {
            //add the greeting to the HTML Markup
            greeting = '<li style="width:100%;">' +
                '<div class="msj-rta macro">' +
                '<div class="text text-r">' +
                '<p>' + data + " Shall we begin by you giving me a your name? "+ '</p>' +
                '<p><small>' + '</small></p>' +
                '</div>' +
                '<div class="avatar" style="padding:0px 0px 0px 10px !important"><img class="img-circle" style="width:100%;" src="' + winston.avatar + '" /></div>' +
                '</li>';
            ////Append to the chat list
            $("ul").append(greeting);
        })
        
});


//User & Winston avatars
var me = {};
me.avatar = "https://christianlifecoachnow.com/wp-content/uploads/2016/12/Blank-Photo.png";
var winston = {};
winston.avatar = "https://i.pinimg.com/736x/c6/82/05/c68205d4d0fad6fe1c926b0211bccac1--fat-animals-funny-animals.jpg";

//Function to add the date to each of the chat messages
function formatAMPM(date) {
    var hours = date.getHours();
    var minutes = date.getMinutes();
    var ampm = hours >= 12 ? 'PM' : 'AM';
    hours = hours % 12;
    hours = hours ? hours : 12; // the hour '0' should be '12'
    minutes = minutes < 10 ? '0' + minutes : minutes;
    var strTime = hours + ':' + minutes + ' ' + ampm;
    return strTime;
}

//-Insert response & input into the html page
function insertChat(who, data, time = 0) {
    var control = "";
    var date = formatAMPM(new Date()); //new instance of date for each message
    //the response will be sent with winston avatar
    //unser input rendered with user avatar
    if (who == "me") {

        control = '<li style="width:100%">' +
            '<div class="msj macro">' +
            '<div class="avatar"><img class="img-circle" style="width:100%;" src="' + me.avatar + '" /></div>' +
            '<div class="text text-l">' +
            '<p>' + data + '</p>' +
            '<p><small>' + date + '</small></p>' +
            '</div>' +
            '</div>' +
            '</li>';
    } else {
        control = '<li style="width:100%;">' +
            '<div class="msj-rta macro">' +
            '<div class="text text-r">' +
            '<p>' + data + '</p>' +
            '<p><small>' + date + '</small></p>' +
            '</div>' +
            '<div class="avatar" style="padding:0px 0px 0px 10px !important"><img class="img-circle" style="width:100%;" src="' + winston.avatar + '" /></div>' +
            '</li>';
    }
    setTimeout(
        function() {
            //add the new html to the list for display
            $("ul").append(control);
            //make scroll bar come to bottom
           $("ul").scrollTop($("ul")[0].scrollHeight);

        }, time);
}

function resetChat() {
    $("ul").empty();
}