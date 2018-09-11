$(".m-nav-ul>li").click(function() {
    $(".m-current-item").removeClass("m-current-item");
    $(this).addClass("m-current-item");
}).mouseenter(function() {
    $(this).addClass("m-enter-item");
}).mouseleave(function() {
    $(this).removeClass("m-enter-item");
});

