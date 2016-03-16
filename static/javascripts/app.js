(function () {
    function showTemplate(template, data) {
        var rendered = Handlebars.templates[template](data);
        $("#dynamic-content").html(rendered);
    }

    showTemplate('loading');

    $.ajax({
        url: '/followers',

        success: function (data) {
            var jsonData = JSON.parse(data);
            var templateData = {followers: jsonData.followers};
            var followersCount = jsonData.followers.length;

            if (followersCount === 0) {
                showTemplate('error-api');
            } else if (followersCount > 75000) {
                showTemplate('error-followers');
            } else {
                showTemplate('result', templateData);
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
