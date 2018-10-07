
$(document).ready(function () {
	$(".navbar-nav>li>a").click(function () {
		$(".m-current-item").removeClass("m-current-item");
		$(this).addClass("m-current-item");
	}).mouseenter(function () {
		$(this).addClass("m-enter-item"); 

		crtid = $(this).attr("id");
		if (crtid == "aboutus") {
			$("#mega-service").hide();
			$("#mega-aboutus").show();
			$("#mega").show();
		} else if (crtid == "service") {
			$("#mega-service").show();
			$("#mega-aboutus").hide();
			$("#mega").show();
		} else {
			$("#mega").hide();
		}
	}).mouseleave(function () {
		$(this).removeClass("m-enter-item"); 
	}); 

	// when mouse leave #mega. hide(); 
	$("#mega").mouseleave(function() {
		$("#mega").hide(); 
	});

});
