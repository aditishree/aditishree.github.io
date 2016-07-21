window.onload = init;

var finalPath = "https://aditishree.github.io/";

function init() {
	$("#projects").click(showproj);
	$("#resume").click(showres);
	$("#cv").click(showcv);
	$("#present").click(showpres);
	$("#about").click(showabt);
	$("#contact").click(showcont);
}

function showproj() {
	window.location = finalPath + "home#projects";
}

function showres() {
	window.location = finalPath + "home#resume";
}

function showcv() {
	window.location = finalPath + "home#cv";
}

function showpres() {
	window.location = finalPath + "home#present";
}

function showabt() {
	window.location = finalPath + "home#about";
}

function showcont() {
	window.location = finalPath + "home#contact";
}