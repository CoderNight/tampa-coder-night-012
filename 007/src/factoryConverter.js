var textToConveyor = function(layoutText) {
    var boards = layoutText.split("\n\n");
    var conveyors = [];

    boards.forEach(function(board){
        var conveyor = {};
        var lines = board.split('\n');
        conveyor.startIndex = lines[1].indexOf('X');
        conveyor.northRow = lines[0].split("").map(laserToBoolean);
        conveyor.southRow = lines[2].split("").map(laserToBoolean);
        conveyors.push(conveyor);
    });

    return conveyors;
}

function laserToBoolean(item){
        if(item === '#') return false;
        else if(item === '|') return true;
        else throw Error("Laser row had an invalid character");
}


module.exports = textToConveyor;