//when the page loads..
$(window).on('load', function() {
    //check if there is a cookie by calling handler
    $.get('/chat-session', {
            name: $('#name').val() //pass in the value of new inout name if needed
        })
        .done(function(data) {
            //if there is no cookie value
            if (data == "") {
                //let modal appear & ask for user name
                $("#greet").text("Why not enter your name so Winston can get to know you better ");
            } else {
                //if cookie exists & have username, hide the input box for name 
                $("#greet").text("Welcome back " + data + ". Winston is looking forward to chatting to you again!");
                $('#name').hide();
            }
            //show the modal after hmtl ready & cookie checkk complete
            $('#myModal').modal('show');
        });
})

//when the user hits enter after entering chata text
$("#user-input-form").submit(
    function(event) {
        //render the input inthe chat window
        insertChat("me", $('#user-input').val());
        //AJAX, don't want page to refresh - instead we just change the content
        event.preventDefault();
        //get the response from winston & pass in user text
        $.get('/user-input', {
                value: $('#user-input').val()
            })
            .done(function(data) { //return winston's response
                insertChat("winston", data); //render it in the HTML
                $('#user-input').val(" ") //clear the  input field
            })
    });

//get the name the user inputs from the modal
$("#name-form").submit(
    function(event) {
        event.preventDefault(); //AJAX
        $.get('/generate-greeting', {
                name: $("#name").val() //generate a greeting ased upon if cookie stored & name saved or new user
            })
            .done(function(data) { //return greeting
                //attatch the greeting to the html for rendering
                greeting = '<li style="width:100%;">' +
                    '<div class="msj-rta macro">' +
                    '<div class="text text-r">' +
                    '<p>' + data + '</p>' +
                    '<p><small>' + '</small></p>' +
                    '</div>' +
                    '<div class="avatar" style="padding:0px 0px 0px 10px !important"><img class="img-circle" style="width:100%;" src="' + winston.avatar + '" /></div>' +
                    '</li>';
                //Append to the chat list
                $("ul").append(greeting);
                //close down the modal
                $('#myModal').modal('toggle');
            })
    });

//User & Winston avatars
var me = {};
me.avatar = "anon.jpeg";
var winston = {};
winston.avatar = "winston.jpg";

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
        //add html, the user inpput/ winston reponse, avatar image & time
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

        }, 200); //wait before appending too html for smoother UX
}

function resetChat() {
    $("ul").empty();
}