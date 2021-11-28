$(document).ready(function () {
	const textarea = document.querySelector("textarea");
	const saveButton = document.getElementById("save");
	const newButton = document.getElementById("new");

	textarea.addEventListener("keyup", function () {
		saveButton.disabled = !this.value;
	});

	newButton.onclick = () => (window.location.href = "/");
	saveButton.onclick = async () => await createPaste(textarea.value);

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
