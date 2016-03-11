var source = $("#landing").html();
var template = Handlebars.compile(source);

$("#dynamic-content").html(template);
