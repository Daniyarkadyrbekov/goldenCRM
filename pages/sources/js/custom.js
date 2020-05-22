$(document).ready(function($) {
    $(".table-row").click(function() {
        window.document.location = $(this).data("href");
    });
});

// initialize with defaults
$("#input-id").fileinput();

// // with plugin options
// $("#input-id").fileinput({'showUpload':false, 'previewFileType':'any'});