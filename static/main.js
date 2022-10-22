function sendData() {
  let XHR = new XMLHttpRequest();

  XHR.addEventListener("load", function (event) {
    alert(event.target.responseText);
  });

  XHR.addEventListener("error", function (event) {
    alert("Oops! something went wrong.");
  });

  let start = document.getElementById("start").value;
  let end = document.getElementById("end").value;

  let username, password;
  let token = document.getElementById("token");

  // as specified in Redmine documentation, token can be passed as username with random password
  if (token) {
    username = token.value;
    password = "password";
  } else {
    username = document.getElementById("username").value;
    password = document.getElementById("password").value;
  }

  console.log(username);

  XHR.open(
    "POST",
    `http://localhost:3333/form?start=${start}&end=${end}&username=${username}&password=${password}`
  );
  XHR.send();
}
