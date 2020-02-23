var shit = Math.ceil((new Date() - new Date("2019/08/30"))/(1000*60*60*24*7))
var week = shit;
function timetable(week){
  $.ajax({
    url: '/gettimetable',
    dataType: 'text',
    data: 'week='+week,
    type: 'get',
    success: function(response){
      console.log(response);
      $("#header").html(response);
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
