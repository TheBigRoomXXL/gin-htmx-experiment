
window.addEventListener("load", () => {
  /* Some JS is need to overcode some limitation of HTML and CCS
      1. CSS does not support proper masonry grid with sorting and height variation
      2. Textarea cannot auto-resize to user input (why the fuck not?!)
  */

  var main = document.getElementById("notes")
  
  var masonry = new Masonry(main, {
    itemSelector: ".note",
    gutter: 10,
      isFitWidth: true,
    transitionDuration: 0,
  });
 
  function afterSettle(event) {
    if (event.detail.requestConfig.verb == "post"){
      resizeTextarea.call(main.firstChild)
      masonry.prepended( main.firstChild );
    }else {
      for (textarea of main.children) {
        resizeTextarea.call(textarea)
      }
      masonry.reloadItems();
    }
    masonry.layout()
  }

  main.addEventListener('htmx:afterSettle', (event) => afterSettle(event));
  
  
  // Bloody hack because textarea cannot grow automatically with user input
  function resizeTextarea() {
    this.style.height = '12px';
    this.style.height = this.scrollHeight + 12 + 'px';
  }
  function resizeTextareaAndRefresh() {
    this.style.height = '12px';
    this.style.height = this.scrollHeight + 12 + 'px';
    masonry.layout()
  }

  function resetInput() {
    create_input.value = "";
  }

  var create_input = document.getElementById('create_note');
  create_input.addEventListener('htmx:afterRequest', resetInput );
  
  textareas = document.querySelectorAll("textarea")
  textareas.forEach( textarea => {
    textarea.style.height = textarea.scrollHeight + 12 + 'px';
    textarea.addEventListener('input', resizeTextareaAndRefresh );
  });
  masonry.layout()
});


