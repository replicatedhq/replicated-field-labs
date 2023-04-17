console.log("ahoy")

var input = document.getElementById("textbox");
var con = document.getElementById("console");
var url

function newline(message,c) {
        if (typeof message.target !== 'undefined'){
            let hr = document.createElement("hr")
            let host = document.createElement("p")
            host.innerText = message.target;
            host.classList.add("host")
            con.appendChild(hr);
            con.appendChild(host);
        }
        let t = document.createElement("p");
        t.innerText = message.text;
        t.classList.add(c);
        con.appendChild(t);
        con.scrollTop = con.scrollHeight;
}

function connect(path) {
    let host = "ntfy.sh"
    let wspath = "wss://"+host+"/"+path+"/ws"
    url = "https://"+host+"/"+path
    console.log("connect to: "+wspath)
    socket = new WebSocket(wspath)
    socket.onopen = function(e) {
        newline({text:"connected"},"green")
    }

    socket.onclose = function(event){
        newline({text:"disconnected"},"red")
    }

    socket.onmessage = function(event){
        let data = JSON.parse(event.data);
        if (typeof data.message !== 'undefined'){
            let message = JSON.parse(data.message);
            switch (message.kind) {
                case "response":
                    newline(message,"default")
            }
        }
    }
}

input.addEventListener("keypress", function(event) {
  if (event.key === "Enter") {
    event.preventDefault();
    if (input.value != "") {
        //console.log(input.value);
        if (input.value.startsWith("/")) {
            let string = input.value.split(" ");
            let command = string[0];
            let arg = string [1];
            console.log(command)
            switch (command) {
                case "/c":
                    connect(arg);
            }
        } else {
            let host = document.querySelector('input[name="host"]:checked').value;
            let payload = {"kind":"shell","target":host,"text":input.value}
            let body = JSON.stringify(payload)
            console.log(body)
            var xhr = new XMLHttpRequest();
            console.log(url)
            xhr.open("POST",url);
            xhr.send(body)

        }
        input.value = "";
    }
  }
});
