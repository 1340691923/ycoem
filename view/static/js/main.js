// 导航设置
$(".dropdown-toggle").click(function(){
	var cssdis = $("#example-navbar-collapse").css("height");
	if (cssdis == "50px") {
		// return false;
		var url = $(this).next().children().eq(0).children().attr("href");
		if (url) {
			window,location.href = url;
		}else{
			return false;
		}
	}else{
		return true;
	}
})

function marqueeStart(ids,text,time){
	// ids    滚动容器
	// text   滚动函数
	// time   执行时间间隔
	var obj = document.getElementById(ids);
	var marqueeVar = window.setInterval(text,time);
	obj.onmouseover = function(){
		window.clearInterval(marqueeVar);
	}
	obj.onmouseout = function(){
		marqueeVar = window.setInterval(text,time);
	}
}

function marquee(ids){
	// ids    滚动容器
	// width  图片宽度
	var obj = document.getElementById(ids);
	var obj1 = document.getElementById(ids+"_1");
	var width = obj1.getElementsByTagName('li')[0].getElementsByTagName('img')[0].width;
	obj.scrollLeft++;

	if(obj.scrollLeft > width+40 ){
        obj1.appendChild(obj1.getElementsByTagName('li')[0]);
        obj.scrollLeft = 22;
    };
}

// 返回顶部按钮
$(window).scroll(function(){
	var sc = $(window).scrollTop();
	if(sc > 0){
		$("#goTopBtn").css("display","block");
	}else{
		$("#goTopBtn").css("display","none");
	}
})
$("#goTopBtn").click(function(){
	var sc=$(window).scrollTop();
	$('body,html').animate({scrollTop:0},500);
})

$(function(){
	// 首页链接
	$(".navbar-nav > .dropdown > a").eq(0).click(function(){
		window.location.href = "/";
	});
	$(".navbar-nav > .dropdown > a:last").click(function(){
		window.location.href = "/concat";
	});
	// 设置二级目录的宽
	var cssdis = $("#example-navbar-collapse").css("height");
	if (cssdis == "50px") {
		var width = $(".navbar-collapse .navbar-nav > li").width()-2;
		$(".index-header .dropdown-menu > li").css("width",width.toString()+"px");

	}
})
