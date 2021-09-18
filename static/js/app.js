const textarea = document.querySelector('textarea');

function copyToClipboard() {
	textarea.select();
	textarea.setSelectionRange(0, 99999);

	navigator.clipboard.writeText(textarea.value).then(_ => console.log("Copied!"))
}

$(document).ready(function () {
	const button = document.getElementById('copy')

	textarea.addEventListener('keyup', function () {
		button.disabled = !this.value;
	});
});
