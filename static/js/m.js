


$(document).ready(function () {
	$(".navbar-nav>li>a").click(function () {
		$(".m-current-item").removeClass("m-current-item");
		$(this).addClass("m-current-item");

		crtid = $(this).attr("id");
		if (crtid == "aboutus") {
			megaDislayFlag = $("#mega").is(":visible");
			megaAboutusFlag = $("#mega-aboutus").is(":visible");

			//alert("mega: " + megaDislayFlag + ", megaAboutusFlag: " + megaAboutusFlag); 
			if (megaDislayFlag && megaAboutusFlag) {
				$("#mega").slideUp("slow");;
			} else {
				$("#mega-service").hide();
				$("#mega-aboutus").show();
				$("#mega").slideDown("slow");
			}
		} else if (crtid == "service") {
			megaDislayFlag = $("#mega").is(":visible");
			megaServiceFlag = $("#mega-service").is(":visible");
			//alert("mega: " + megaDislayFlag + ", megaServiceFlag: " + megaServiceFlag); 
			if (megaDislayFlag && megaServiceFlag) {
				$("#mega").slideUp("slow");
			} else {
				$("#mega-aboutus").hide();
				$("#mega-service").show();
				$("#mega").slideDown("slow");
			}
		};

	}).mouseenter(function () {
		$(this).addClass("m-enter-item");
	}).mouseleave(function () {
		$(this).removeClass("m-enter-item");
	});

	// 鼠标滚动时, 隐藏mega
	$("window").scroll(function () {
		console.log("33333333333333333")
		$("#mega").hide();
	});

});
