(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['result'] = template({"1":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : {}, alias2=helpers.helperMissing, alias3="function", alias4=container.escapeExpression;

  return "        <input id='"
    + alias4(((helper = (helper = helpers.screen_name || (depth0 != null ? depth0.screen_name : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"screen_name","hash":{},"data":data}) : helper)))
    + "' class='input' type='radio' value='"
    + alias4(((helper = (helper = helpers.screen_name || (depth0 != null ? depth0.screen_name : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"screen_name","hash":{},"data":data}) : helper)))
    + "' name='name' checked>\n        <label class='label' for='"
    + alias4(((helper = (helper = helpers.screen_name || (depth0 != null ? depth0.screen_name : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"screen_name","hash":{},"data":data}) : helper)))
    + "'>"
    + alias4(((helper = (helper = helpers.screen_name || (depth0 != null ? depth0.screen_name : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"screen_name","hash":{},"data":data}) : helper)))
    + "</label>\n        <br>\n";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1;

  return "<div class='b-result'>\n    <form class='b-form'>\n      <img class='img-result' src='/images/results.png' alt='result'>\n      <div class='result-text'>(select someone to thank)</div>\n"
    + ((stack1 = helpers.each.call(depth0 != null ? depth0 : {},(depth0 != null ? depth0.followers : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "        <a class='twitter post-tweet-btn' href='http://twitter.com/share'></a>\n    </form>\n    <img class='img-thanks' src='/images/thanks.png' alt='thanks'>\n</div>\n";
},"useData":true});
templates['error-api'] = template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class='b-info-box'>\n    <img class='img-sorry' src='/images/sorry.png' alt='sorry'>\n    An unknown error has occurred, try&nbsp;again later.\n</div>\n";
},"useData":true});
templates['error-followers'] = template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class='b-info-box'>\n    <img class='img-sorry' src='/images/sorry.png' alt='sorry'>\n    Unfortunately you are too cool to use this right now.\n    <br>\n    -\n    <br>\n    You have 75,000+ followers\n</div>\n";
},"useData":true});
templates['landing'] = template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class='b-landing'>\n    <img class='img-landing-text' src='/images/landing-text.png' alt='landing-text'>\n    <a class='btn sign-in-btn' href='/login'>\n        <img src='/images/sign-in.png' alt='sign-in'>\n    </a>\n    <div class='elements-container'>\n        <div class='element-right'>\n            <a class='btn twitter-follow-btn' href=\"https://twitter.com/intent/follow?screen_name=@OdysseusArms\">\n                <img src='/images/follow.png' alt='follow'>\n            </a>\n        </div>\n        <div class='element-left'>\n            <img class='logo' src='/images/logo.png' alt='logo'>\n        </div>\n    </div>\n</div>\n\n<script src=\"https://platform.twitter.com/widgets.js\"></script>\n";
},"useData":true});
templates['loading'] = template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<img src='/images/bird.gif' alt='loading'>\n";
},"useData":true});
templates['thank-you'] = template({"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    return "<div class='b-info-box'>\n    <img class='img-thank-you' src='/images/thank-you.png' alt='thank you'>\n    for using ThankBot.\n    <br>\n    Your Tweet has been sent.\n</div>\n";
},"useData":true});
})();