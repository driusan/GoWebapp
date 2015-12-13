
function deleteItem(id) {
	var xhr = new XMLHttpRequest();
	xhr.open("DELETE", "/items/" + id);
	//xhr.setRequestHeader("If-Match", "234");
	xhr.onload = function() {
		if (xhr.status === 200) {
			window.location = "/"

			console.log("Deleted " + id);
		}  else {
			console.log("Failed to delete " + id);
		}
	}
	xhr.send();
}
function updateItem(id, val) {
	var xhr = new XMLHttpRequest();
	xhr.open("PUT", "/items/" + id);
	//xhr.setRequestHeader("If-Match", "234");
	xhr.onload = function() {
		if (xhr.status === 200) {
			window.location = "/"
		}  else {
			console.log("Failed to update" + id);
		}
	}
	xhr.send(val);
}
