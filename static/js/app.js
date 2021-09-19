$(document).ready(function () {
	const textarea = document.querySelector('textarea');
	const copyButton = document.getElementById('copy')
	const saveButton = document.getElementById('save')
	const newButton = document.getElementById('new')
	const rawButton = document.getElementById('raw')

	textarea.addEventListener('keyup', function () {
		copyButton.disabled = !this.value;
		saveButton.disabled = !this.value;
	});

	newButton.onclick = () => window.location.href = '/'
	rawButton.onclick = () => window.location.href = '/raw'

	copyButton.onclick = () => copyToClipboard()
	saveButton.onclick = async () => await createPaste(textarea.value)

	function copyToClipboard() {
		textarea.select();
		textarea.setSelectionRange(0, 99999);
		navigator.clipboard.writeText(textarea.value).then(_ => console.log("Copied!"))
	}

	async function createPaste(content) {
		let response = await fetch("/api/paste", {
			method: "POST",
			headers: {"Content-Type": "application/json"},
			body: content
		})
		console.log(response)
	}
});

