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
            showTemplate('result', templateData);
        },

        error: function () {
            showTemplate('landing');
        }
    });

})();
