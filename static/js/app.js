const textarea = document.querySelector('textarea');

function copyToClipboard() {
	textarea.select();
	textarea.setSelectionRange(0, 99999);

	navigator.clipboard.writeText(textarea.value).then(_ => console.log("Copied!"))
}

$(document).ready(function () {
	const copyButton = document.getElementById('copy')
	const saveButton = document.getElementById('save')

	textarea.addEventListener('keyup', function () {
		copyButton.disabled = !this.value;
		saveButton.disabled = !this.value;
	});
});
