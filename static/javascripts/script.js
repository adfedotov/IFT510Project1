const main = function() {
  const host = 'http://' + window.location.host;

  $('.menu a').on('click', function(event) {
    const scrollTarget = $(this).attr('to');
    $('html, body').animate({
      scrollTop: $(scrollTarget).offset().top
    }, 2000);
    return false;
  });

  $('.admin a.login').on('click', function(event) {
  	const code = $('.admin input').val();
  	$.post(host + '/login', code, function(data) {
      window.location.href = host + '/admin/edit';
    });
  	return false;
  });

  $('.edit a.logout').on('click', function(event) {
    $.post(host + '/logout', function(data) {
      window.location.href = host + '/admin';
    });
    return false;
  });
}

$(document).ready(main);
