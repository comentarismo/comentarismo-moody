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
    emojione = require("emojione"),
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
    var lang = getURLParameter('lang') || "en";
    var jsonStr = getURLParameter('json');

    if (vid) {
        $("#vid").val(vid);
        if (vid) {
            $("#loading").fadeIn(500);
            $.getJSON("/moody?vid=" + vid+"&lang="+lang, function (resp) {
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
        $(target).text(func(decodeURIComponent(json.metadata[value]), limit));
    } else if (json) {
        $(target).text(decodeURIComponent(json.metadata[value]));
    } else {
        $(target).text(value);
    }
}

function setCss(header, target, json) {
    $(header).css({
        "background-image": "url('" + decodeURIComponent(json.metadata[target]) + "')"
    });
}

function drawReport(json) {
    $("#vid").val(json.url);
    $("#report").show();
    if (json.type === 'YouTubeVideo') {

        setText("#video_title", "title", json, truncate, 50);
        setText("#channel_title", "channeltitle", json);
        setText("#network_title", "YouTube");
        setText("#total_comments", "totalcomments", json);
        setCss("#header", "thumbnail", json);

    } else if (json.type === 'InstagramPic') {
        $("#video_title").text(truncate(decodeURIComponent(json.metadata.caption), 50));
        $("#channel_title").text(decodeURIComponent(json.metadata.username));
        $("#network_title").text("Instagram");
        $("#total_comments").text(decodeURIComponent(json.metadata.totalcomments));
        $("#video_views").text(decodeURIComponent(json.metadata.totallikes));
        $("#header").css({
            "background-image": "url('" + decodeURIComponent(json.metadata.thumbnail) + "')"
        });
    } else if (json.type === 'FacebookPost') {
        $("#video_title").text(truncate(decodeURIComponent(json.metadata.caption), 50));
        $("#channel_title").text(decodeURIComponent(json.metadata.username));
        $("#network_title").text("Facebook");
        $("#total_comments").text(decodeURIComponent(json.metadata.totalcomments));
        $("#video_views").text(decodeURIComponent(json.metadata.totallikes));
        $("#header").css({
            "background-image": "url('" + decodeURIComponent(json.metadata.thumbnail) + "')"
        });
    } else if (json.type === 'VineVideo') {
        $("#video_title").text(truncate(decodeURIComponent(json.metadata.caption), 50));
        $("#channel_title").text(decodeURIComponent(json.metadata.username));
        $("#network_title").text("Vine");
        $("#total_comments").text(decodeURIComponent(json.metadata.totalcomments));
        $("#video_views").text(decodeURIComponent(json.metadata.totallikes));
        $("#header").css({
            "background-image": "url('" + decodeURIComponent(json.metadata.thumbnail) + "')"
        });
    }

    $("#comments_per_day").text(json.commentavgperday.toFixed(2));
    $(".progress-bar").attr('aria-valuenow', json.commentcoveragepercent)
        .css({'width': json.commentcoveragepercent + '%'})
        .text('Comments Analyzed: ' + json.commentcoveragepercent + '%');

    // Sort the keyword list
    sortable = [];
    for (var k in json.keywords) {
        sortable.push([k, json.keywords[k]]);
    }

    sortable.sort(function (a, b) {
        return b[1] - a[1]
    })

    var y = 0;
    for (var s in sortable) {
        y++;
        $("#keywords").append('<tr><td>' + y + '.</td><td>' + sortable[s][0] + ' (' + sortable[s][1] + ')</td></tr>');
    }

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


    var emojis = {
        "Terrible!": terrible + emojione.shortnameToUnicode(":scream:"),
        "Sucks": sucks + emojione.shortnameToUnicode(":angry:"),
        "Bad": bad + emojione.shortnameToUnicode(":worried:"),
        "Not Good": notgood + emojione.shortnameToUnicode(":unamused:"),
        "Eh": eh + emojione.shortnameToUnicode(":confused:"),
        "Neutral": neutral + emojione.shortnameToUnicode(":expressionless:"),
        "OK": ok + emojione.shortnameToUnicode(":neutral_face:"),
        "Good": good + emojione.shortnameToUnicode(":smile:"),
        "Like It": likeit + emojione.shortnameToUnicode(":smiley:"),
        "Loved It": lovedit + emojione.shortnameToUnicode(":yum:"),
        "Awesome!": awesome + emojione.shortnameToUnicode(":grinning:"),
        "Unknown": unknown + emojione.shortnameToUnicode(":no_mouth:")
    };

    var elabels = [
        emojis[terrible], emojis[sucks], emojis[bad], emojis[notgood], emojis[eh], emojis[neutral],
        emojis[ok], emojis[good], emojis[likeit], emojis[lovedit], emojis[awesome], emojis[unknown]
    ];

    var labels = [terrible, sucks, bad, notgood, eh, neutral, ok, good, likeit, lovedit, awesome, unknown];
    var dp = {};
    for (var x in json.sentiment) {
        var n = json.sentiment[x]["name"];
        if (n == "") {
            n = unknown;
        }
        var p = json.sentiment[x]["percent"];

        dp[n] = p;
    }

    var ctx = document.getElementById("myChart").getContext("2d");

    // Bar Chart
    var data = {
        labels: elabels,
        datasets: [
            {
                label: "Percent",
                backgroundColor: [
                    "#FF0000",
                    "#E80957",
                    "#E809BF",
                    "#BB09E8",
                    "#2709E8",
                    "#0957E8",
                    "#0980E8",
                    "#09E8BF",
                    "#09E85B",
                    "#23E809",
                    "#B8E809",
                    "#E8E409",
                ],
                borderColor: "rgba(255,99,132,1)",
                borderWidth: 1,
                hoverBackgroundColor: "rgba(255,99,132,0.4)",
                hoverBorderColor: "rgba(255,99,132,1)",
                data: [dv(dp[terrible], 0), dv(dp[sucks], 0), dv(dp[bad], 0), dv(dp[notgood], 0), dv(dp[eh], 0), dv(dp[neutral], 0), dv(dp[ok], 0), dv(dp[good], 0), dv(dp[likeit], 0), dv(dp[lovedit], 0), dv(dp[awesome], 0), dv(dp[unknown], 0)]
            }
        ]
    };


    new Chart(ctx,{type: 'bar', data:data, options: {showLines: false, responsive: true, maintainAspectRatio: false}});


    var ctx2 = document.getElementById("scaleChart").getContext("2d");

    // Bar Chart
    var data2 = {
        labels: elabels,
        datasets: [
            {
                label: "Percent",
                backgroundColor: [
                    "#FF0000",
                    "#E80957",
                    "#E809BF",
                    "#BB09E8",
                    "#2709E8",
                    "#0957E8",
                    "#0980E8",
                    "#09E8BF",
                    "#09E85B",
                    "#23E809",
                    "#B8E809",
                    "#E8E409",
                ],
                borderColor: "rgba(255,99,132,1)",
                borderWidth: 1,
                hoverBackgroundColor: "rgba(255,99,132,0.4)",
                hoverBorderColor: "rgba(255,99,132,1)",
                data: [dv(dp[terrible], 0), dv(dp[sucks], 0), dv(dp[bad], 0), dv(dp[notgood], 0), dv(dp[eh], 0), dv(dp[neutral], 0), dv(dp[ok], 0), dv(dp[good], 0), dv(dp[likeit], 0), dv(dp[lovedit], 0), dv(dp[awesome], 0), dv(dp[unknown], 0)]
            }
        ]
    };


    new Chart(ctx2,{type: 'pie', data:data2, options: {}});


    //if (json.topcomments) {
    //    json.topcomments.forEach(function (v, i, a) {
    //        $("#sampling").append('<div class="comment">' + emojis[v.sentiment] + '<h4><a href="' + userLink(v.authorname, json.type) + '">@' + v.authorname + '</a> (' + v.likes + ' likes)</h4><p>' + emojione.shortnameToImage(v.content) + '</p></div>')
    //    });
    //}

    //if (json.sentimentlist) {
    //
    //    for (var x in json.sentimentlist) {
    //        json.sentimentlist[x].forEach(function (v, i, a) {
    //            $("#all_sampling").append('<div class="comment">' + emojis[v.sentiment] + '<h4><a href="' + userLink(v.authorname, json.type) + '">@' + v.authorname + '</a> (' + v.likes + ' likes)</h4><p>' +  emojione.shortnameToImage(v.content) + '</p></div>')
    //
    //        });
    //    }
    //}




}
