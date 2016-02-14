window.onload = function() {
    var vid = getURLParameter('vid');
    var jsonStr = getURLParameter('json');
    if(vid) {
        jQuery("form input").val(vid);
        if(vid) {
            jQuery("#loading").fadeIn(500);
            jQuery.getJSON("/api?vid="+vid, function(resp) {
                if(resp.Error){
                    console.log(resp)
                }else {
                    jQuery("#loading").fadeOut(500, function() {
                        window.json = resp;
                        drawReport(resp);
                    });
                }
            });
        }
    } else if(jsonStr) {
        var jsonObj = JSON.parse(jsonStr);
        if(jsonObj) {
            drawReport(jsonObj);
            jQuery("nav").hide();
        }
    }
}

function getURLParameter(name) {
    return decodeURIComponent((new RegExp('[?|&]' + name + '=' + '([^&;]+?)(&|#|;|$)').exec(location.search)||[,""])[1].replace(/\+/g, '%20'))||null
}

function dv(data, defaultData) {
    return (data ? data : defaultData);
}

function truncate(str, max) {
    if(str.length > max) {
        str = str.substring(0, max) + "...";
    }

    return str;
}

function userLink(username, type) {
    if(type == 'YouTubeVideo') {
        return "http://youtube.com/user/" + username;
    } else if(type == 'InstagramPic') {
        return "http://instagram.com/" + username;
    }

    return "#";
}

function drawReport(json) {
    jQuery("nav input").val(json.URL);
    jQuery("#report").show();
    if (json.Type == 'YouTubeVideo') {
        jQuery("#video_title").text(truncate(decodeURIComponent(json.Metadata.Title), 50));
        jQuery("#channel_title").text(decodeURIComponent(json.Metadata.ChannelTitle));
        jQuery("#network_title").text("YouTube");
        jQuery("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        jQuery("#header").css({
            "background-image": "url('"+decodeURIComponent(json.Metadata.Thumbnail)+"')"
        });
    } else if (json.Type == 'InstagramPic') {
        jQuery("#video_title").text(truncate(decodeURIComponent(json.Metadata.Caption), 50));
        jQuery("#channel_title").text(decodeURIComponent(json.Metadata.UserName));
        jQuery("#network_title").text("Instagram");
        jQuery("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        jQuery("#video_views").text(decodeURIComponent(json.Metadata.TotalLikes));
        jQuery("#header").css({
            "background-image": "url('"+decodeURIComponent(json.Metadata.Thumbnail)+"')"
        });
    } else if (json.Type == 'FacebookPost') {
        jQuery("#video_title").text(truncate(decodeURIComponent(json.Metadata.Caption), 50));
        jQuery("#channel_title").text(decodeURIComponent(json.Metadata.UserName));
        jQuery("#network_title").text("Facebook");
        jQuery("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        jQuery("#video_views").text(decodeURIComponent(json.Metadata.TotalLikes));
        jQuery("#header").css({
            "background-image": "url('"+decodeURIComponent(json.Metadata.Thumbnail)+"')"
        });
    } else if (json.Type == 'VineVideo') {
        jQuery("#video_title").text(truncate(decodeURIComponent(json.Metadata.Caption), 50));
        jQuery("#channel_title").text(decodeURIComponent(json.Metadata.UserName));
        jQuery("#network_title").text("Vine");
        jQuery("#total_comments").text(decodeURIComponent(json.Metadata.TotalComments));
        jQuery("#video_views").text(decodeURIComponent(json.Metadata.TotalLikes));
        jQuery("#header").css({
            "background-image": "url('"+decodeURIComponent(json.Metadata.Thumbnail)+"')"
        });
    }

    jQuery("#comments_per_day").text(json.CommentAvgPerDay.toFixed(2));
    jQuery(".progress-bar").attr('aria-valuenow', json.CommentCoveragePercent)
        .css({'width': json.CommentCoveragePercent+'%'})
        .text('Comments Analyzed: '+json.CommentCoveragePercent+'%');

    // Sort the keyword list
    sortable = []
    for (var x in json.Keywords) {
        sortable.push([x, json.Keywords[x]]);
    }

    sortable.sort(function(a, b) {return b[1] - a[1]})

    y = 0
    for (var x in sortable) {
        y++;
        jQuery("#keywords").append('<tr><td>'+y+'.</td><td>'+sortable[x][0]+' ('+sortable[x][1]+')</td></tr>');
    }

    if (json.TopComments) {
        jQuery("#sampling_header").text("Top Comments")
        json.TopComments.forEach(function(v, i, a) {
            jQuery("#sampling").append('<div class="comment"><h4><a href="'+userLink(v.AuthorName, json.Type)+'">@'+v.AuthorName+'</a> ('+v.Likes+' likes)</h4><p>'+v.Content+'</p></div>')
        });
    } else if (json.SampleComments) {
        json.SampleComments.forEach(function(v, i, a) {
            jQuery("#sampling").append('<div class="comment"><h4><a href="'+userLink(v.AuthorName, json.Type)+'">@'+v.AuthorName+'</a></h4><p>'+v.Content+'</p></div>')
        });
    }

    var dp = {};
    for (var x in json.Sentiment) {
        var n = json.Sentiment[x]["Name"];
        var p = json.Sentiment[x]["Percent"];

        dp[n] = p;
    }

    var ctx = document.getElementById("myChart").getContext("2d");

    // Bar Chart
    var data = {
        labels: ["Terrible!", "Sucks", "Bad", "Not Good", "Eh", "Neutral", "OK", "Good", "Like It", "Loved It", "Awesome!"],
        datasets: [
            {
                label: "Percent",
                data: [dv(dp["-5"], 0), dv(dp["-4"], 0), dv(dp["-3"], 0), dv(dp["-2"], 0), dv(dp["-1"], 0), dv(dp["0"], 0), dv(dp["1"], 0), dv(dp["2"], 0), dv(dp["3"], 0), dv(dp["4"], 0), dv(dp["5"], 0)]
            }
        ]
    }
    var myBarChart = new Chart(ctx).Bar(data, {scaleShowGridLines: false, responsive: true, maintainAspectRatio: false});
}