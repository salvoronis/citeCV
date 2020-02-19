var file;
$("#file").change(function(){
  file = this.files[0];
  var formData = new FormData();
  formData.append('file', file);
  $.ajax({
    url: '/profile',
    dataType: 'text',
    cache: false,
    contentType: false,
    processData: false,
    data: formData,
    type: 'post',
    success: function(response){
      console.log(response)
    }
  });
});
