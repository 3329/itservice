$(".m-collapse>ul>li").click(function() {
    $(this).addClass("m-current-item");
}).mouseenter(function() {
    $(this).addClass("m-enter-item");
}).mouseleave(function() {
    $(this).removeClass("m-enter-item");
});

$(".")