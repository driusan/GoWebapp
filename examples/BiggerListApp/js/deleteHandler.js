
function addTag(realxhr, id, data) {
	var headxhr = new XMLHttpRequest();
	headxhr.open("HEAD", "/items/" + id);
	headxhr.onload = function() {
		if(headxhr.status === 200) {
			realxhr.setRequestHeader("If-Match", headxhr.getResponseHeader("ETag"));
		}
		realxhr.send(data);
	}
	headxhr.send()
}
function deleteItem(id) {
	var xhr = new XMLHttpRequest();
	xhr.open("DELETE", "/items/" + id);
	xhr.onload = function() {
		if (xhr.status === 200) {
			window.location = "/"
		}  else {
			console.log("Failed to delete " + id);
		}
	}
	addTag(xhr, id);
}
function updateItem(id, val) {
	var xhr = new XMLHttpRequest();
	xhr.open("PUT", "/items/" + id);
	xhr.onload = function() {
		if (xhr.status === 200) {
			window.location = "/"
		}  else {
			console.log("Failed to update" + id);
		}
	}
	addTag(xhr, id, val);
}
