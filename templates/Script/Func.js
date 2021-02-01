function googleCon() {
  const numberOfReloads = window.location.hash.substr(1) | 0;
  window.location.hash = numberOfReloads + 1;
  console.log("555");
  if (confirm("Do you want to connect your google account?")) {
    var win = open(
      "https://f92a609a8d2e.ngrok.io/googleregister",
      "_self",
      "_blank"
    );
  } else {
  }
}

function lineCon() {
  const numberOfReloads = window.location.hash.substr(1) | 0;
  window.location.hash = numberOfReloads + 1;
  console.log("555");
  if (confirm("Do you want to connect your google account?")) {
    var win = open(
      "https://f92a609a8d2e.ngrok.io/lineregister",
      "_self",
      "_blank"
    );
  } else {
  }
}

function facebookCon() {
  const numberOfReloads = window.location.hash.substr(1) | 0;
  window.location.hash = numberOfReloads + 1;
  console.log("555");
  if (confirm("Do you want to connect your google account?")) {
    var win = open(
      "https://f92a609a8d2e.ngrok.io/facebookregister",
      "_self",
      "_blank"
    );
  } else {
  }
}

function alertLogOut() {
  if (confirm("Are you OK that you want to log out??")) {
    var win = open("https://f92a609a8d2e.ngrok.io/logout", "_self", "_blank");
  } else {
  }
}
