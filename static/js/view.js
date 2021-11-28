$(document).ready(function () {
	const textarea = document.querySelector("textarea");
	const copyButton = document.getElementById("copy");
	const newButton = document.getElementById("new");
	const rawButton = document.getElementById("raw");

	newButton.onclick = () => (window.location.href = "/");
	rawButton.onclick = () => (window.location.href += "/raw");
	copyButton.onclick = () => navigator.clipboard.writeText(window.location.href).then(() => {
		alert("URL Copied to Clipboard!");
	});

	textarea.remove();

	hljs.highlightAll();
});
