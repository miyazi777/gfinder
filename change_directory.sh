#!/bin/bash

echo '
  {
    "command": "/usr/bin/gnome-terminal",
    "command_args": ["--", "zsh", "-c", "cd $HOME; cd ${info}; zsh"],
    "input_resources": [
      {
        "name": "xbackend-api",
        "info": "/home/takeshi-miyajima/workspace_highway/product1/backend-api"
      },
      { 
        "name": "xfrontend-web",
        "info": "/home/takeshi-miyajima/workspace_highway/product1/frontend-web"
      },
      {
        "name": "hiway memo",
        "info": "/home/takeshi-miyajima/workspace_highway/memo"
      },
      {
        "name": "private memo",
        "info": "/home/takeshi-miyajima/private/memo"
      }
    ]
  }
' 
