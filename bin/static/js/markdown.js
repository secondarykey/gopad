function change() {
    var src = $("#content").val();
    var html = marked(src);
    $('#result').html(html);
    $('pre code').each(function(i, block) {
        hljs.highlightBlock(block);
    });
}

$(document).ready(function() {

    var box = $('#memoForm');
    if ( box.length != 0 ) {
        var boxTop = box.offset().top;
        $(window).scroll(function() {
            if( $(window).scrollTop() >= boxTop - 30 ) {
                box.addClass('memoFix');
            } else {
                box.removeClass('memoFix');
            }
        });
    }

    $('#saveBtn').click(function() {
      $.ajax({
         url: location.href,
         type: 'POST',
         data: { 
             "title" : $("#title").val(),
             "content" : $("#content").val(),
         },
         dataType: 'json'
      }).success(function( data ) {

      }).error(function() {
          alert("Error!");
      });
      return false;
    });

    $('#deleteBtn').click(function() {
      $.ajax({
         url: location.href,
         type: 'DELETE',
         data: { },
         dataType: 'json'
      }).success(function( data ) {
          location.href = "/";
      }).error(function() {
          alert("Error!");
      });
      return false;
    });

    marked.setOptions({
        langPrefix: ''
    });

    $('#content').keyup(function() {
        change();
    });
    change();
});
