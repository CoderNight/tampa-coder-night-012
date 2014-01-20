var fs = require('fs'),
    converter = require('./factoryConverter');
    robot = require('./robot');



var rawContent = fs.readFileSync(process.argv[2], {encoding : 'utf8'});

var factorList = converter(rawContent);

var commandList  = robot(factorList);

var output = fs.createWriteStream("./output.txt");
//do something with command, for now we log it to console
commandList.forEach(function(command){
    output.write(command + "\n");
})
