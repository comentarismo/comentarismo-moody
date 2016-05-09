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

