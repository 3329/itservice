

$(".m-nav-ul>li").click(function() {
    $(".m-current-item").removeClass("m-current-item");
    $(this).addClass("m-current-item"); 
    
    /*
    console.log($(this).attr("id"));
    if ($(this).attr("id") == "service") {
        $(".m-sub-nav").show();
    } else if ($(this).attr("id") == "ooo") {
        $(".m-sub-nav").show();
    } else {
        $(".m-sub-nav").hide();
    }
    */
}).mouseenter(function() {
    $(this).addClass("m-enter-item");
}).mouseleave(function() {
    $(this).removeClass("m-enter-item");
});

$(document).ready(function () {
    $(".m-nav-ul>li").click(function() {
        $(".m-current-item").removeClass("m-current-item");
        $(this).addClass("m-current-item"); 
    }).mouseenter(function() {
        $(this).addClass("m-enter-item");
    }).mouseleave(function() {
        $(this).removeClass("m-enter-item");
    });
});