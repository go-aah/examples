// JS file

function showFilename() {
  var filePath = document.getElementById('fileUpload').value
  if (filePath) {
    if (filePath.indexOf('\\')) {
      filePath = filePath.substring(filePath.lastIndexOf('\\'))
    } else if (filePath.indexOf('/')) {
      filePath = filePath.substring(filePath.lastIndexOf('/'))
    }

    var chr = filePath.charAt(0);
    if (chr === '\\' || chr === '/') {
      filePath = filePath.substring(1);
    }

    document.getElementById('filenameDisplay').innerHTML = filePath;
    document.getElementById('fileName').value = filePath;
  }
}
