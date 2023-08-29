!function(e){"use strict";var o,t;e(".dropdown-menu a.dropdown-toggle").on("click",function(t){return e(this).next().hasClass("show")||e(this).parents(".dropdown-menu").first().find(".show").removeClass("show"),e(this).next(".dropdown-menu").toggleClass("show"),!1}),[].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]')).map(function(t){return new bootstrap.Tooltip(t)}),[].slice.call(document.querySelectorAll('[data-bs-toggle="popover"]')).map(function(t){return new bootstrap.Popover(t)}),o=document.getElementsByTagName("body")[0],(t=document.querySelectorAll(".light-dark-mode"))&&t.length&&t.forEach(function(t){t.addEventListener("click",function(t){o.hasAttribute("data-bs-theme")&&"dark"==o.getAttribute("data-bs-theme")?document.body.setAttribute("data-bs-theme","light"):document.body.setAttribute("data-bs-theme","dark")})}),Waves.init()}(jQuery);
// 业务代码
!function () {
    $(function () {
        $.toast = function (text) {
            $('#toast').text(text)
            $('#toast').fadeIn(500).delay(5000).fadeOut(500);
        }
        var conn;
        var msg = document.getElementById("msg");
        var log = document.getElementById("log");

        // function appendLog(item) {
        //     var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        //     log.appendChild(item);
        //     if (doScroll) {
        //         log.scrollTop = log.scrollHeight - log.clientHeight;
        //     }
        // }

        // document.getElementById("form").onsubmit = function () {
        //     if (!conn) {
        //         return false;
        //     }
        //     if (!msg.value) {
        //         return false;
        //     }
        //     conn.send(msg.value);
        //     msg.value = "";
        //     return false;
        // };

        if (window["WebSocket"]) {
            conn = new WebSocket("ws://localhost:8080/ws");
            // conn.close()
            conn.onopen = function () {
                console.log("🚀🚀 WebSocket 连接成功")
            }
            conn.onclose = function (evt) {
                $.toast("Connection closed.")
            };
            conn.onmessage = function (evt) {
                var messages = evt.data.split('\n');
                for (var i = 0; i < messages.length; i++) {
                    var item = document.createElement("div");
                    item.innerText = messages[i];
                    // appendLog(item);
                }
            };
            return
        }
        $.toast("Your browser does not support WebSockets.")
    })
}()