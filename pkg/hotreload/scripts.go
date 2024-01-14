package hotreload

import (
	"bytes"
	"fmt"
)

var socketScripts []byte

func generate(port string) func(b []byte) []byte {
	socketScripts = []byte(fmt.Sprintf(script, port))
	return func(b []byte) []byte {
		newBody := append([]byte("<body>"), socketScripts...)
		b = bytes.Replace(b, []byte("<body>"), newBody, 1)
		return b
	}
}

var script = `
<script>
    const socket = new WebSocket("ws://localhost:%s/ws");

    socket.addEventListener("open", (event) => {
        console.log("hot reload: Connected to hot reload server");

        const newDiv = document.createElement("div");
        newDiv.textContent = "Connected to hot reload server";
        newDiv.style.cssText = "color:white;background-color:black;padding:15px;position:fixed;right:0;bottom:0";

        document.body.appendChild(newDiv);
    
        setTimeout(() => {
          newDiv.remove();
        }, 2000)
    });

    socket.addEventListener("message", (event) => {
      if(event && event.data === "refresh") {
        console.log("hot reload: reloading page...")
        document.location.reload()
      }
    });

    socket.addEventListener("close", (event) => {
        console.log("hot reload: Disconnected from hot reload server");
        const newDiv = document.createElement("div");
        newDiv.textContent = "Disonnected from hot reload server";
        newDiv.style.cssText = "color:white;background-color:black;padding:15px;position:fixed;right:0;bottom:0";

        document.body.appendChild(newDiv);
    
        setTimeout(() => {
          newDiv.remove();
        }, 2000)
    });
</script>
`
