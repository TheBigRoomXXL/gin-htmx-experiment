
window.addEventListener("load", () => {
  // Setup masonry grid with js as it is not supported yet by CSS
  // I would like to replace that dependy by a self-made one later
  var main = document.getElementById("notes")
  
  var masonry = new Masonry(main, {
    itemSelector: ".note",
    gutter: 10,
      isFitWidth: true,
    transitionDuration: 0,
  });
 
  function afterSettle(event) {
    
    if (event.detail.requestConfig.verb == "post"){
      masonry.prepended( main.firstChild );
    }else {
      masonry.reloadItems()
    }
    masonry.layout()
  }

  function enterEditMode (event) {
    console.log(event)
    note = event.target
    // note.setAttribute("contenteditable", true)

  }

  main.addEventListener("click", enterEditMode)
  main.addEventListener('htmx:afterSettle', (event) => afterSettle(event));



  // Bloody hack because textarea cannot grow automatically with user input
  function resizeTextarea() {
    this.style.height = '24px';
    this.style.height = this.scrollHeight + 12 + 'px';
  }

  function resetInput() {
    input.value = "";
  }

  var input = document.getElementById('create_note');
  input.addEventListener('input', resizeTextarea );
  input.addEventListener('htmx:afterRequest', resetInput );
});


