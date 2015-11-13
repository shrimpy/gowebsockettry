# gowebsockettry
Golang WebSocket sample on Azure App Service

[![Deploy to Azure](http://azuredeploy.net/deploybutton.png)](https://azuredeploy.net/)


- Trouble Shooting

If you are seeing below error from client, make sure you enable "WebSocket" for your site. By default is off.

````
WebSocket connection to 'wss://{your host name}.azurewebsites.net/echo' failed: Error during WebSocket handshake: Unexpected response code: 503
````

