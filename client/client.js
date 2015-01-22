// var ws = new WebSocket("ws://104.236.97.245/subscribe");
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


var pushstream = new PushStream({
    host:"104.236.97.245",
    port:"80",
    modes: 'websocket'
});

pushstream.onmessage = function(e) {
    console.log(e);
    $("#received").prepend("<li>"+e.data+"</li>");
    var num = parseInt($("#msg_received").text());
    num+=1;
    $("#msg_received").text(num.toString());
};
pushstream.onestatuschange = function(status){
    console.log("new status",status);
    if (status==PushStream.OPEN){
        $('#con_status').val("open");
    }else if (status==PushStream.CLOSED){
        $('#con_status').val("closed");
    }
}
pushstream.onerror = function(error){
    console.log("error",error);
}
pushstream.connect();
