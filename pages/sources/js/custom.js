function processForm(e) {
    if (e.preventDefault) e.preventDefault();

    /* do what you want with the form */

    // You must return false to prevent the default form behavior
    return false;
}

var form = document.getElementById('my-form');
if (form.attachEvent) {
    form.attachEvent("submit", processForm);
} else {
    form.addEventListener("submit", processForm);
}