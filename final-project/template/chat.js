var app = {}
app.ws = undefined
app.container = undefined

app.print = function (message) {
    var el = document.createElement("p")
    el.innerHTML = message
    app.container.append(el)
}

app.doSendMessage = function () {
    var messageRaw = document.querySelector('.input-message').value
    app.ws.send(JSON.stringify({
        Message: messageRaw
    }));

    var message = '<b>me</b>: ' + messageRaw
    app.print(message)

    document.querySelector('.input-message').value = ''
}

app.init = function () {
    if (!(window.WebSocket)) {
        alert('Your browser does not support WebSocket')
        return
    }

    var name = prompt('Enter your name please:') || "No name"
    document.querySelector('.username').innerText = name

    var age = prompt('Enter your age please:') || "17"
    document.querySelector('.age').innerText = age

    app.container = document.querySelector('.container')

    app.ws = new WebSocket("ws://localhost:8080/ws?username=" + name + "&age=" + age)

    app.ws.onopen = function() {
        var message = '<b>me</b>: connected'
        app.print(message)
    }

    app.ws.onmessage = function (event) {
        var res = JSON.parse(event.data)

        var messsage = ''
        if (res.Type === 'New User') {
            message = 'User <b>' + res.From + '</b>: connected'
        } else if (res.Type === 'Leave') {
            message = 'User <b>' + res.From + '</b>: disconnected'
        } else {
            message = '<b>' + res.From + '</b>: ' + res.Message
        }

        app.print(message)
    }

    app.ws.onclose = function () {
        var message = '<b>me</b>: disconnected'
        app.print(message)
    }
}

window.onload = app.init