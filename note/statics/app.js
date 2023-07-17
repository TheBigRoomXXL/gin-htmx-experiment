
// Script for notes.html

// Setup masonry grid with js as it is not supported yet by CSS
// I would like to replace that dependy by a self-made one later
function afterSettle(event, masonry) {
  
  if (event.detail.requestConfig.verb == "post"){
    masonry.prepended( main.firstChild );
  }else {
    masonry.reloadItems()
  }
  masonry.layout()
}

window.addEventListener("load", () => {
  var main = document.getElementById("notes")
  var masonry = new Masonry(main, {
    itemSelector: ".grid-item",
    gutter: 10,
    isFitWidth: true,
    transitionDuration: 0,
  });
  main.addEventListener('htmx:afterSettle', (event) => afterSettle(event, masonry));
})


// Script for create.hml

// Bloody hack because text area cannot grow automatically with user input
function resizeTextarea() {
  this.style.height = '24px';
  this.style.height = this.scrollHeight + 12 + 'px';
}

function resetInput() {
  input.value = "";
}

window.addEventListener("load", () => {
  var input = document.getElementById('create_note');
  input.addEventListener('input', resizeTextarea );
  input.addEventListener('htmx:afterRequest', resetInput );
});


