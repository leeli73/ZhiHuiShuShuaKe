function query() {
    if ($("#popbox_title").length > 0) {
        $(".popboxes_close")[0].click();
    }

    if ($("#chapterList .time_ico.fl").nextAll()[2].children[0].style.width === "100%" || $("video").get(0).ended) {
        var num = -1;
        var text = $("#chapterList .time_ico.fl").parent().nextAll()[++num].id;
        while (text === "" ||
               text.substr(0, 5) != "video" ||
               text.substr(0, 7) === "video-0") {
            text = $("#chapterList .time_ico.fl").parent().nextAll()[++num].id;
        }
        $("#chapterList .time_ico.fl").parent().nextAll()[num].click();
    }

    if ($("video").length > 0 && $("video").get(0).playbackRate != 1.5) {
        $(".speedTab15")[0].click();
    }

    if ($("video").get(0).volume > 0) {
        $(".volumeIcon").click();
	}
}
window.setInterval(query, 1000);
