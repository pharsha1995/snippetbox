let navLinks = document.querySelectorAll("nav a");

for (let navLink of navLinks) {
  if (navLink.getAttribute("href") === window.location.pathname) {
    navLink.classList.add("live");
    break;
  }
}