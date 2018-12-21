# SweetNyanCraft Modpack updater

This is used to update the current modpack for multimc modpacks.

Because MultiMC doesnt have a feature to update modpacks sadly so out os fustration i wrote this.



Currently lots of stuff is hardcoded. this is going to chance ofcourse to support more modpacks.

Written in nOdE.jS. compiled into binrary files using `pkg` so the user doesnt need to have the nOdE.jS runtime installed even.


## Features

* Basic SHA-1 Verifying
* Downloads the Zip file
* extracts the zip file in current directory


## planned features
* JSON file format on the server for SHA-1, version checking etc.
* Version manager so it doesnt redownload the already installed version
* Prevent from deleting user-added mods.
* Switch from SHA-1 to something else?
* abillity to launch the game? and interface?
