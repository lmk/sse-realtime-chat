package main

import "html/template"

var html = template.Must(template.New("chat_room").Parse(`
<html lang="kr"> 
<head> 
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width">
    <title>{{.roomid}}</title>
    <script src="https://code.jquery.com/jquery-3.5.0.js"></script>
    <script> 
        $(window).on('load', function(){
            
            $('#message_form').val('');
            $('#message_form').focus();

            if (!!window.EventSource) {
                var source = new EventSource('/stream/{{.roomid}}');
                source.addEventListener('message', function(e) {
                    $('#messages').append(e.data + "</br>");
                    $('html, body').animate({scrollTop:$(document).height()}, 'slow');

                }, false);
            } else {
                alert("NOT SUPPORTED");
            }

            $('#myForm').submit(function(event) {
              event.preventDefault();
              var msg = $('#message_form').val()
      
              if (msg.trim() != '') {

                $.post( "/room/{{.roomid}}", { user: $('#user_form').val(), message: msg })
                .fail(function( xhr, status, error ){
                  console.log( "Error: ", xhr, status, error );
                });
                
                $('#user_form').attr('readonly', true);
                $('#message_form').val('');
                $('#message_form').focus();
              }
            })
            
            
        });
    </script> 
    </head>
    <body>
    <h1>Welcome to {{.roomid}} room</h1>
    <div id="messages"></div>
    <form id="myForm" action="/room/{{.roomid}}" method="post" onsubmit=""> 
    User: <input id="user_form" name="user" value="{{.userid}}" placeholder="{{.userid}}">
    Message: <input id="message_form" name="message" placeholder="input message">
    <input type="submit" value="Submit"> 
    </form>
</body>
</html>
`))
