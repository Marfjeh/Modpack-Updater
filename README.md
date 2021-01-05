# SweetNyanCraft Modpack updater

This is used to update the current modpack for MultiMC modpacks.

Because, MultiMC doesn't have a feature to update modpacks sadly so out of frustration I wrote this.

Written in Go. which is much better than the node.js version of this. You can still get the source code in the branch: `legacy/nodejs```


## Features

* Basic SHA-1 Verifying, making sure that the zip file downloaded is correct and not corrupted.
* Downloads the Zip file
* extracts the zip file in current directory
* basic version checking, it only downloads if the version is different
* JSON file for the modpack, for name version etc.


## planned features
* Prevent from deleting user-added mods.
* Switch from SHA-1 to something else?
* ability to launch the game and interface? (maybe launcher replacement?) with microsoft login support




## Compiling binrarys

Compiling this is really simple, as everything is in main.go, you should have go installed on your system.

build it with `go build main.go`

please note this will only compile it for your system. to compile it for different platforms refer to the golang doc.


## using this for your own pack

first you need to have a webserver, upload the php file in the `web` directory to your webserver.

This is your version manager, so if you release a new version, change the line where it says `$modpack->version = "0.0.1";` to for example `$modpack->version = "1.0.0";`

now you need to make the downloadable zip, place the `config` folder and `mods` folder in the zip with the `modpack.json` from this project, do not forget to set the right urls etc.

after this, you can test launching it, it should download the zip and then extract it in the same folder where the executable is located.

you can put this as a pre-launch script in multimc. place the binrary in the `.minecraft` folder from the instance from multimc.

`edit intance` > `settings` > `custom commands` > Pre-launch command: `$INST_MC_DIR\modpack-updater64.exe`

Done!