/***********************************************
* CSS3 Analog Clock- by JavaScript Kit (www.javascriptkit.com)
* Visit JavaScript Kit at http://www.javascriptkit.com/ for this script and 100s more
***********************************************/

var hourHand = document.querySelector('#liveclock div.hand.hour');
var minuteHand = document.querySelector('#liveclock div.hand.minute');
var secondHand = document.querySelector('#liveclock div.hand.second');

window.requestAnimationFrame = window.requestAnimationFrame
|| window.mozRequestAnimationFrame
|| window.webkitRequestAnimationFrame
|| window.msRequestAnimationFrame
|| function(f){setTimeout(f, 60)}


function updateclock(){
	var curdate = new Date()
	var hour_as_degree = ( curdate.getHours() + curdate.getMinutes()/60 ) / 12 * 360
	var minute_as_degree = curdate.getMinutes() / 60 * 360
	var second_as_degree = ( curdate.getSeconds() + curdate.getMilliseconds()/1000 ) /60 * 360
	hourHand.style.transform = 'rotate(' + hour_as_degree + 'deg)';
	minuteHand.style.transform = 'rotate(' + minute_as_degree + 'deg)';
	secondHand.style.transform = 'rotate(' + second_as_degree + 'deg)';
	requestAnimationFrame(updateclock)
}

requestAnimationFrame(updateclock)
