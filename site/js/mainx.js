const genButton = document.getElementById("regen-button");
genButton.addEventListener("click", () => {
  fetch("/gen")
  .then(() => location.reload());
});

const compileButton = document.getElementById("compile-button");
compileButton.addEventListener("click", () => {
  fetch("/compile")
  .then(() => location.reload());
});
