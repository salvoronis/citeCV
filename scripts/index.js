window.onload = function () {
    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");

    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }

    document.getElementById("form").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(msg.value);
        msg.value = "";
        return false;
    };

    if (window["WebSocket"]) {
        conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }
};

var shit = Math.ceil((new Date() - new Date("2019/08/30"))/(1000*60*60*24*7))
var week = shit;
function timetable(week){
  $.ajax({
    url: '/gettimetable',
    dataType: 'text',
    data: 'week='+week,
    type: 'get',
    success: function(response){
      $("tbody").html(response);
    }
  });
}
timetable(week);
$("#weekn").html(week)
$("#prew").click(function(){
  week -= 1;
  timetable(week)
  $("#weekn").html(week)
});
$("#next").click(function(){
  week += 1;
  timetable(week)
  $("#weekn").html(week)
});
var offset = 0;
$("#getNews").click(getNews);
function getNews(){
  $.ajax({
    url: '/news',
    dataType: 'text',
    data: {"offset" : offset},
    type: 'post',
    success: function(response){
      $("#news").append(response);
      offset += 10;
    }
  });
}
getNews();
