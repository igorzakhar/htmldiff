function readTextFile(fileInputId) {

  const file = document.getElementById(fileInputId).files[0];
  if (file) {
    getText(file, fileInputId);
  }
}

function getText(readFile, fileInputId) {

  let reader = new FileReader();

  reader.readAsText(readFile, "UTF-8");

  reader.onload = function (event) {
    let textareaId = "textarea" + fileInputId[fileInputId.length - 1];
    document.getElementById(textareaId).value = event.target.result;
  };

  reader.onerror = function (event) {
    console.error("File could not be read! Code " + event.target.error.code);
  };
}


function ReadFiles(inputIdList) {
  inputIdList.forEach(function (inputId) {
    document.getElementById(inputId).addEventListener("change", function (event) {
      readTextFile(inputId);
    }, false);
  });
}

document.addEventListener('DOMContentLoaded',ReadFiles(["file1", "file2"]));