 var commands = function(conveyors){
    var commands = [];
    conveyors.forEach(function(conveyor){
        commands.push(getCommand(conveyor));
    });

    return commands;
 }

 function getCommand(conveyor){
     var westCount = 0, eastCount =0;
     var tick = 0;
     for (var i = conveyor.startIndex; i >= 0; i--){
         if(conveyor.northRow[i] && tick % 2 == 0
             || conveyor.southRow[i] && tick % 2 == 1 )
             westCount++;
         tick++;
     }

     tick = 0;
     for (i = conveyor.startIndex; i < conveyor.northRow.length; i++){
         if(conveyor.northRow[i] && tick % 2 == 0
             || conveyor.southRow[i] && tick % 2 == 1 )
             eastCount++;
         tick++
     }

     return eastCount >= westCount ? "GO WEST" : "GO EAST";
 }

 module.exports = commands