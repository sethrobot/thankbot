var source = $("#result").html();
var template = Handlebars.compile(source);

$("#dynamic-content").html(template);
