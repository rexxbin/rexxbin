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

	copyButton.onclick = () => copyToClipboard();
	saveButton.onclick = async () => await createPaste(textarea.value);

	function copyToClipboard() {
		textarea.select();
		textarea.setSelectionRange(0, 99999);
		navigator.clipboard
			.writeText(textarea.value)
			.then((_) => console.log("Copied!"));
	}

	async function createPaste(content) {
		await fetch("/api/paste", {
			method: "POST",
			headers: { "Content-Type": "application/json" },
			body: content,
		}).then((res) => {
			if (res.ok) {
				window.location.href = res.url;
			}
		});
	}

	if (window.location.href.indexOf("/v/") > -1) {
		textarea.readOnly = true;

		rawButton.onclick = () => (window.location.href += "/raw");
		saveButton.remove();

		rawButton.style.display = "revert";
		copyButton.disabled = false;
		textarea.remove();
		viewArea.style.display = "block";
		hljs.highlightAll();
	}
});
