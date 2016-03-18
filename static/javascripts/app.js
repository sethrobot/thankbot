(function () {
    function showTemplate(template, data) {
        var rendered = Handlebars.templates[template](data);
        $("#dynamic-content").html(rendered);
    }

    function postTweet() {
        $('.post-tweet-btn').click(function() {
            var radioValue = $('input:radio[name=name]:checked').val();

            var width  = 575,
                height = 400,
                left   = ($(window).width()  - width)  / 2,
                top    = ($(window).height() - height) / 2,
                url    = this.href + '?text=@' + radioValue + '%20Thanks%20for%20being%20my%20first%20follower%20pic.twitter.com/32Prr3FTnm&url=empty',
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
                showTemplate('result', templateData);
                postTweet();
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
