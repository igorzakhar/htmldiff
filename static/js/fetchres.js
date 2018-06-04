let legendHTML = `
      <div class="my-legend">
        <div class="legend-title pb-1">Legend:</div>
        <div class="legend-scale">
        <ul class="legend-labels">
          <li><span class="added"></span>Added</li>
          <li><span class="deleted"></span>Deleted</li>
        </ul>
        </div>
      </div>
`;

function escapeHtml (string) {
  let entityMap = {
    '&': '&amp;',
    '<': '&lt;',
    '>': '&gt;',
    '"': '&quot;',
    "'": '&#39;',
    '/': '&#x2F;',
    '`': '&#x60;',
    '=': '&#x3D;'
  };

  return String(string).replace(/[&<>"'`=\/]/g, function (s) {
    return entityMap[s];
  });
}

let compareButton = document.getElementById("compare-button");

compareButton.addEventListener("click", function () {
  let original = escapeHtml(document.getElementById("textarea1").value);
  let changed = escapeHtml(document.getElementById("textarea2").value);

  let resultDiff = document.getElementById("diff-result") 
  
  fetch('/api/v1/htmldiff', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      text1: original, 
      text2: changed,
    })
  }).then(function(response) {  

    response.json().then(function(data) {  
      resultDiff.innerHTML = "<h3 class=\"mt-2\">Result:</h3>" + data.result + legendHTML
    }).catch(function(error) {
      resultDiff.innerHTML = "<h3 class=\"mt-2\">ERROR:</h3>" + error
    });
  });

});
