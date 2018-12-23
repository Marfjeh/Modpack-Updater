# SweetNyanCraft Modpack updater

This is used to update the current modpack for multimc modpacks.

Because MultiMC doesnt have a feature to update modpacks sadly so out of fustration i wrote this.

Written in nOdE.jS. compiled into binrary files using `pkg` so the user doesnt need to have the nOdE.jS runtime installed even.


## Features

* Basic SHA-1 Verifying, making sure that the zip file downloaded is correct and not currupted.
* Downloads the Zip file
* extracts the zip file in current directory
* basic version checking, it only downloads if the version is diffrent
* JSON file for the modpack, for name version etc.


## planned features
* Prevent from deleting user-added mods.
* Switch from SHA-1 to something else?
* abillity to launch the game? and interface? (maybe launcher replacement?)




## Compiling binrarys

Compiling is really simple, only you need to do is installing `pkg` by doing: `npm install -g pkg`

after that install the depenencys by doing `npm install`

if you done that, compile it! `npm run build-all`

this creates the `bin` foler and the linux and windows builds will be placed there.

please note, this does not create builds for MacOS. it is possible to do so, but i cant test it if it works.



## using this for your own pack

first you need to have a webserver, upload the php file in the `web` directory to your webserver.

This is your version manager, so if you release a new version, change the line where it says `$modpack->version = "0.0.1";` to for example `$modpack->version = "1.0.0";`

now you need to make the downloadable zip, place the `config` folder and `mods` folder in the zip with the `modpack.json` from this project, do not forget to set the right urls etc.

after this, you can test launching it, it should download the zip and then extract it in the same folder where the executable is located.

you can put this as a pre-launch script in multimc. place the binrary in the `.minecraft` folder from the instance from multimc.

after that, do `edit intance` > `settings` > `custom commands` > Pre-launch command: `$INST_MC_DIR\modpack-updater64.exe`

Done!