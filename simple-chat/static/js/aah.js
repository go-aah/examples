// JS for Simple Chat

function byId(id) {
  return document.getElementById(id);
}

// http://youmightnotneedjquery.com/#on
function on(id, en, fn) {
  var obj = byId(id);
  if (obj.attachEvent) {
    obj.attachEvent(en, fn);
  } else {
    obj.addEventListener(en, fn);
  }
}

function addChatMsg(msg) {
  var li = document.createElement('li');
  li.appendChild(document.createTextNode(msg));
  byId('chatMessages').appendChild(li);
}

function connectWebSocket(uri) {
  var chatSocket = new WebSocket(uri);

  chatSocket.onopen = function (e) {
    byId('chatStatus').innerHTML = 'Connected: ' + uri;
    byId('chatStatus').style.color = 'green';
  }

  chatSocket.onmessage = function (e) {
    addChatMsg(e.data);
  }

  chatSocket.onclose = function (e) {
    byId('chatStatus').innerHTML = 'Disconnected: ' + uri;
    byId('chatStatus').style.color = 'red';
  }

  chatSocket.onerror = function (e) {
    console.log(e);
  }

  return chatSocket;
}
