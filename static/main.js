function sendData() {
  let XHR = new XMLHttpRequest();

  XHR.addEventListener("load", function (event) {
    alert(event.target.responseText);

    let snackbarContainer = document.querySelector("#demo-toast-example");
    let showToastButton = document.querySelector("#demo-show-toast");
    showToastButton.addEventListener("click", function () {
      "use strict";
      let data = { message: event.target.responseText };
      snackbarContainer.MaterialSnackbar.showSnackbar(data);
    });
  });

  XHR.addEventListener("error", function (event) {
    alert("Oops! something went wrong.");
  });

  let start = document.getElementById("start").value;
  let end = document.getElementById("end").value;
  let username = document.getElementById("username").value;
  let password = document.getElementById("password").value;

  XHR.open(
    "POST",
    `http://localhost:3030/form?start=${start}&end=${end}&username=${username}&password=${password}`
  );
  XHR.send();
}
