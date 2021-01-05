# SweetNyanCraft Modpack updater

This is used to update the current modpack for MultiMC modpacks.
This is a simple Modpack updater for MultiMC instances. it also works without MultiMC.

It keeps everything simple without configuration hell. 1 zip file and 1 json file is needed really.

This may not make some mod developers happy that i'm rehosting their mods but downloading from curse is just a nightmare.
(And i'm lazy sorry.)

Currently MultiMC doesn't have a feature to update modpacks sadly so out of frustration I wrote this.
because I do not want to use different launchers such like Techniclauncher, Curse launcher, GDLauncher (So bloated)
MultiMC is not as bloated.

This is written in Go. which is much better than the bloated node.js version of this. You can still get the source code in the branch: `legacy/nodejs```




<b>If you're a mod developer please consider this:</b>

I'm not making money by rehosting your mods nor do I claim that I made the mods. I'm not those Malvertising websites like 9minecraft that is full with malware and ads.

My only goal is only playing and having fun with your mods, and making it easier for me to distribute updates to my

friends without giving zip files to everyone when I change something.

Maybe i'll consider changing the way to still download it from curse. but for now not going to do that sorry.


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

You're free to use this for your own modpack. you do not need to ask permission only what I ask you to comply to GNU version 3 Licence,
and crediting me would be much appreciated. I would love to see people using it or even make it better.

first you need to have a webserver, upload the php file in the `web` directory to your webserver.

This is your version manager, so if you release a new version, change the line where it says `$modpack->version = "0.0.1";` to for example `$modpack->version = "1.0.0";`

now you need to make the downloadable zip, place the `config` folder and `mods` folder in the zip with the `modpack.json` from this project, do not forget to set the right urls etc.

after this, you can test launching it, it should download the zip and then extract it in the same folder where the executable is located.

you can put this as a pre-launch script in multimc. place the binrary in the `.minecraft` folder from the instance from multimc.

`edit intance` > `settings` > `custom commands` > Pre-launch command: `$INST_MC_DIR\modpack-updater64.exe`

Done!