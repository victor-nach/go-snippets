var express = require('express');
var http = require('http');

// http client
http.request({host: 'localhost', port: 3000}, function(response) {    
    response.on('data', function(data) { process.stdout.write(data);});    
    // response.on('end', function(){server.close();});
}).end();