const downloadURL = "https://fileserver.marfprojects.nl/modpack/download/FabulousCraft_patch.zip";
const http = require('https');
const fs = require('fs');
const unzip = require("unzip");
const sha1File = require('sha1-file');
const rimraf = require('rimraf');
const request = require("request");
let sha1expected = "";
console.log("SweetNyanCraft Modpack Updater Version 0.1");
console.log("------------------------------------------");
console.log("WARNING! This will overwrite your mods and config folders!");
console.log("This means it will remove custom mods you've added.");

function getsha1() {
    console.log("Getting Checksum from Server...");
    request({
        uri: "https://fileserver.marfprojects.nl/modpack/update.txt",
        method: "GET",
        timeout: 1000,
        followRedirect: true,
        maxRedirects: 10
        }, function(error, response, body) {
            if (error) {
                console.error("Error getting checksum. do you have internet?");
                console.error(error);
                process.exit(1);
            }
            sha1expected = body.replace(/\n$/, '');
            console.log("got: " + sha1expected);
            console.log("Downloading update...");
            download(downloadURL, "./test.zip", null);
        });
}

function clearFolders() {
    rimraf('./mods/*', function () { 
        rimraf('./config/*', function () { 
            unzipjob();
        });
    });
}

function verify() {    
    console.log("Checking SHA-1 sum...");
    let sha1zip = sha1File("./test.zip");
    console.log("Expecting SHA-1: " + sha1expected);
    console.log("Got SHA-1: " + sha1zip);
    if (sha1zip === sha1expected) {
        console.log("SHA-1 Verify success! Unzipping...");
        clearFolders();
    }
    else {
        console.error("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@");
        console.error("@  WARNING: SHA-1 IDENTIFICATION MATCH IS INVAILID!  @");
        console.error("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@");
        console.error("IT IS POSSIBLE THAT SOMEONE IS DOING SOMETHING NASTY!");
        console.error("Someone could be eavesdropping on you right now (man-in-the-middle attack)!");
        console.error("Or, Marf derped and he didnt update the correct SHA-1 on the server.");
        console.error("Cancelling update...");
        process.exit(1);
    }
    
}

function unzipjob() {
    fs.createReadStream('./test.zip').pipe(unzip.Extract({ path: '.' }));
    fs.unlinkSync("./test.zip");
    console.log("Finishing in a bit, still extracting...");
}

var download = function(url, dest, cb) {
  var file = fs.createWriteStream(dest);
  var request = http.get(url, function(response) {
    response.pipe(file);
    file.on('finish', function() {
      file.close(cb);
      verify();
    });
  }).on('error', function(err) { // Handle errors
    fs.unlink(dest);
    if (cb) cb(err.message);
  });
};
//clearFolders();
getsha1();
