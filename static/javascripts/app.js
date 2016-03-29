(function () {
    function showTemplate(template, data) {
        var rendered = Handlebars.templates[template](data);
        $("#dynamic-content").html(rendered);
    }

    function postTweet() {
        $('.post-tweet-btn').click(function() {
            var radioValue = $('input:radio[name=name]:checked').val();

            var text_shout_param = encodeURIComponent('Shout out to @');
            var for_what_you_thank_part = encodeURIComponent(' for being my #FoundingFollower on twitter! (Thank yours at www.thankbot.co) pic.twitter.com/32Prr3FTnm');

            var width  = 575,
                height = 400,
                left   = ($(window).width()  - width)  / 2,
                top    = ($(window).height() - height) / 2,
                url    = this.href + '?text=' + text_shout_param + radioValue + for_what_you_thank_part + '&url=empty',
                opts   = 'status=1' +
                    ',width='  + width  +
                    ',height=' + height +
                    ',top='    + top    +
                    ',left='   + left;

            var twitterWindow = window.open(url, 'twitter', opts);

            var interval = window.setInterval(function() {
                try {
                    if (twitterWindow == null || twitterWindow.closed) {
                        window.clearInterval(interval);
                        showTemplate('thank-you');
                    }
                }
                catch (e) {
                    console.warn('Error has occured.');
                    console.trace();
                }
            }, 1000);

            return false;
        });
    }

    showTemplate('loading');

    $.ajax({
        url: '/followers',

        success: function (data) {
            var jsonData = JSON.parse(data);
            var templateData = {followers: jsonData.followers};
            var followersCount = jsonData.followers.length;

            Handlebars.registerHelper('indexNum', function(value) {
                return parseInt(value) + 1;
            });

            if (followersCount === 0) {
                showTemplate('error-api');
            } else if (followersCount > 75000) {
                showTemplate('error-followers');
            } else {
                function showResult() {
                    showTemplate('result', templateData);
                    var radioFirst = $('input:radio[name=name]')[0];
                    radioFirst.checked = true;
                    postTweet();
                }
                setTimeout(showResult, 8000);
            }
        },

        error: function (xhr) {
            if (xhr.status === 401) {
                showTemplate('landing');
            } else {
                showTemplate('error-api');
            }
        }
    });

})();
