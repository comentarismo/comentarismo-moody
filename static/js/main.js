// Avoid `console` errors in browsers that lack a console.
(function () {
    var method;
    var noop = function () {
    };
    var methods = [
        'assert', 'clear', 'count', 'debug', 'dir', 'dirxml', 'error',
        'exception', 'group', 'groupCollapsed', 'groupEnd', 'info', 'log',
        'markTimeline', 'profile', 'profileEnd', 'table', 'time', 'timeEnd',
        'timeStamp', 'trace', 'warn'
    ];
    var length = methods.length;
    var console = (window.console = window.console || {});

    while (length--) {
        method = methods[length];

        // Only stub undefined methods.
        if (!console[method]) {
            console[method] = noop;
        }
    }
}());

$ = require('jquery'),
    analytics = require('ga-browser')(),
    _ = require('underscore'),
    jdenticon = require("jdenticon"),
    md5 = require('md5'),
    Chart = require("chart.js"),
    moment = require('moment');

var page;
var host = "";


ComentarismoMoody = function (options) {

    analytics('create', 'UA-51773618-1', 'auto');
    this.$el = $(options.selector);

    var vid = getURLParameter('vid');
    var jsonStr = getURLParameter('json');

    if (vid) {
        $("form input").val(vid);
        if (vid) {
            $("#loading").fadeIn(500);
            $.getJSON("/moody?vid=" + vid, function (resp) {
                if (resp.Error) {
                    console.log(resp)
                } else {
                    $("#loading").fadeOut(500, function () {
                        window.json = resp;
                        drawReport(resp);
                    });
                }
            });
        }
    } else if (jsonStr) {
        var jsonObj = JSON.parse(jsonStr);
        if (jsonObj) {
            drawReport(jsonObj);
            $("nav").hide();
        }
    }

};


function getURLParameter(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search) || [, ""])[1].replace(/\+/g, '%20')) || null
}

function dv(data, defaultData) {
    return (data ? data : defaultData);
}

function truncate(str, max) {
    if (str.length > max) {
        str = str.substring(0, max) + "...";
    }

    return str;
}

function userLink(username, type) {
    if (type == 'YouTubeVideo') {
        return "http://youtube.com/user/" + username;
    } else if (type == 'InstagramPic') {
        return "http://instagram.com/" + username;
    }

    return "#";
}

function setText(target, value, json, func, limit) {
    if (func) {
        $(target).text(func(decodeURIComponent(json.Metadata[value]), limit));
    } else if (json) {
        $(target).text(decodeURIComponent(json.Metadata[value]));
    } else {
        $(target).text(value);
    }
}

function setCss(header, target, json) {
    $(header).css({
        "background-image": "url('" + decodeURIComponent(json.Metadata[target]) + "')"
    });
}

function drawReport(json) {
    $("nav input").val(json.URL);
    $("#report").show();
    if (json.Type == 'YouTubeVideo') {

        setText("#video_title", "Title", json, truncate, 50);
        setText("#channel_title", "ChannelTitle", json);
        setText("#network_title", "YouTube");
        setText("#total_comments", "TotalComments", json);
        setCss("#header", "Thumbnail", json);

    } else if (json.Type == 'InstagramPic') {
        $("#video_title").text(truncate(decodeURIComponent(json.Metadata.Caption), 50));
        $("#channel_title").text(decodeURIComponent(json.Metadata.UserName));
        $("#network_title").text("Instagram");
        $("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        $("#video_views").text(decodeURIComponent(json.Metadata.TotalLikes));
        $("#header").css({
            "background-image": "url('" + decodeURIComponent(json.Metadata.Thumbnail) + "')"
        });
    } else if (json.Type == 'FacebookPost') {
        $("#video_title").text(truncate(decodeURIComponent(json.Metadata.Caption), 50));
        $("#channel_title").text(decodeURIComponent(json.Metadata.UserName));
        $("#network_title").text("Facebook");
        $("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        $("#video_views").text(decodeURIComponent(json.Metadata.TotalLikes));
        $("#header").css({
            "background-image": "url('" + decodeURIComponent(json.Metadata.Thumbnail) + "')"
        });
    } else if (json.Type == 'VineVideo') {
        $("#video_title").text(truncate(decodeURIComponent(json.Metadata.Caption), 50));
        $("#channel_title").text(decodeURIComponent(json.Metadata.UserName));
        $("#network_title").text("Vine");
        $("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        $("#video_views").text(decodeURIComponent(json.Metadata.TotalLikes));
        $("#header").css({
            "background-image": "url('" + decodeURIComponent(json.Metadata.Thumbnail) + "')"
        });
    }

    $("#comments_per_day").text(json.CommentAvgPerDay.toFixed(2));
    $(".progress-bar").attr('aria-valuenow', json.CommentCoveragePercent)
        .css({'width': json.CommentCoveragePercent + '%'})
        .text('Comments Analyzed: ' + json.CommentCoveragePercent + '%');

    // Sort the keyword list
    sortable = [];
    for (var k in json.Keywords) {
        sortable.push([k, json.Keywords[k]]);
    }

    sortable.sort(function (a, b) {
        return b[1] - a[1]
    })

    var y = 0;
    for (var s in sortable) {
        y++;
        $("#keywords").append('<tr><td>' + y + '.</td><td>' + sortable[s][0] + ' (' + sortable[s][1] + ')</td></tr>');
    }

    if (json.TopComments) {
        json.TopComments.forEach(function (v, i, a) {
            $("#sampling").append('<div class="comment">' + v.Sentiment + '<h4><a href="' + userLink(v.AuthorName, json.Type) + '">@' + v.AuthorName + '</a> (' + v.Likes + ' likes)</h4><p>' + v.Content + '</p></div>')
        });
    }
    var dp = {};
    var terrible = "Terrible!";
    var sucks = "Sucks";
    var bad = "Bad";
    var notgood = "Not Good";
    var eh = "Eh";
    var neutral = "Neutral";
    var ok = "OK";
    var good = "Good";
    var likeit = "Like It";
    var lovedit = "Loved It";
    var awesome = "Awesome!";
    var unknown = "Unknown";

    var labels = [terrible, sucks, bad, notgood, eh, neutral, ok, good, likeit, lovedit, awesome, unknown];
    var dp = {};
    for (var x in json.Sentiment) {
        var n = json.Sentiment[x]["Name"];
        if (n == "") {
            n = unknown;
        }
        var p = json.Sentiment[x]["Percent"];

        dp[n] = p;
    }

    var ctx = document.getElementById("myChart").getContext("2d");

    // Bar Chart
    var data = {
        labels: labels,
        datasets: [
            {
                label: "Percent",
                data: [dv(dp[terrible], 0), dv(dp[sucks], 0), dv(dp[bad], 0), dv(dp[notgood], 0), dv(dp[eh], 0), dv(dp[neutral], 0), dv(dp[ok], 0), dv(dp[good], 0), dv(dp[likeit], 0), dv(dp[lovedit], 0), dv(dp[awesome], 0), dv(dp[unknown], 0)]
            }
        ]
    };
    new Chart(ctx).Bar(data, {scaleShowGridLines: false, responsive: true, maintainAspectRatio: false});


    if (json.SentimentList) {

        for (var x in json.SentimentList) {
            json.SentimentList[x].forEach(function (v, i, a) {
                $("#all_sampling").append('<div class="comment">' + v.Sentiment + '<h4><a href="' + userLink(v.AuthorName, json.Type) + '">@' + v.AuthorName + '</a> (' + v.Likes + ' likes)</h4><p>' + v.Content + '</p></div>')

            });
        }
    }


}
