$(document).ready(function () {
	const textarea = document.querySelector("textarea");
	const copyButton = document.getElementById("copy");
	const saveButton = document.getElementById("save");
	const newButton = document.getElementById("new");
	const rawButton = document.getElementById("raw");
	const viewArea = document.getElementById("view-code");

	textarea.addEventListener("keyup", function () {
		copyButton.disabled = !this.value;
		saveButton.disabled = !this.value;
	});

	newButton.onclick = () => (window.location.href = "/");
	rawButton.style.display = "none";
	viewArea.style.display = "none";

	copyButton.style.display = "none";
	saveButton.onclick = async () => await createPaste(textarea.value);

	copyButton.onclick = () => navigator.clipboard.writeText(window.location.href).then(() => {
		alert("URL Copied to Clipboard!");
	});

	if (window.location.href.indexOf("/v/") > -1) {
		textarea.readOnly = true;

		rawButton.onclick = () => (window.location.href += "/raw");
		saveButton.remove();

		rawButton.style.display = "revert";
		copyButton.style.display = "revert";
		copyButton.disabled = false;

		textarea.remove();
		viewArea.style.display = "block";
		hljs.highlightAll();
	}

	async function createPaste(content) {
		await fetch("/api/paste", {
			method: "POST",
			headers: {"Content-Type": "application/json"},
			body: content,
		}).then((res) => {
			if (res.ok) {
				window.location.href = res.url;
			}
		});
	}
});
