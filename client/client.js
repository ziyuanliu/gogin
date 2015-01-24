// var ws = new WebSocket("ws://localhost:5000/subscribe");
// // var ws1 = new WebSocket("ws://127.0.0.1:9080/subscribe");
// ws.onclose = function() { // thing to do on close
//     $('#con_status').val("closed");
// };
//
// ws.onerror = function(error) { // thing to do on error
//     $('#con_status').val("Error "+error);
// };
//
// ws.onmessage = function(e) { // thing to do on message
//     $("#received").prepend("<li>"+e.data+"</li>");
//     var num = parseInt($("#msg_received").text());
//     num+=1;
//     $("#msg_received").text(num.toString());
// };
//
// ws.onopen = function() { // thing to do on open
//     ws.send("hi");
//     $("#con_status").text("open");
// };
function messageReceived(text, id,channel) {
    console.log(text);
    // document.getElementById('messages').innerHTML += id + ': ' + text + '<br>';
};

console.log(window.location.hostname,window.location.port);
var pushstream = new PushStream({
    host:"0.0.0.0",
    port:"5000",
    modes: 'websocket'
});

// pushstream.onmessage=messageReceived;
pushstream.onmessage = function(text,id,channel) {
    $("#received").prepend("<li>"+text+"</li>");
    var num = parseInt($("#msg_received").text());
    num+=1;
    $("#msg_received").text(num.toString());
};

pushstream.onopen = function(){
    $('#con_status').val("open");
};

pushstream.onclose = function(){
    $('#con_status').val("close");
};

pushstream.onstatuschange = function(status){
    if (status==PushStream.OPEN){
        $('#con_status').text("open");
    }else if (status==PushStream.CLOSED){
        $('#con_status').text("closed");
    }
};
pushstream.onerror = function(error){
    console.log("error",error);
};
try {
pushstream.addChannel("1");
pushstream.connect();
} catch(e) {alert(e)};
