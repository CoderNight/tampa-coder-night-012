var assert = require('assert'),
    fs = require('fs'),
    path = require('path'),
    factoryConverter = require('../src/factoryConverter.js');


describe("Factory Converter Test", function(){
    before(function(done){
        var rawData = fs.readFileSync(path.join(__dirname, 'data/test.txt'), {encoding:'utf8'});
        this.testData = factoryConverter(rawData);
        done();
    })

    it("correct number of conveyors created", function(done){
        assert.equal(this.testData.length, 3, "Expected 2 conveyor objects");
        done();
    })

    it("first conveyor is built correctly", function(done){
        var conveyor = this.testData[0];
        assert.equal(conveyor.startIndex, 3, "Robot should be at 3");
        assert.deepEqual([false,true,false,true,false,true,false,false],
            conveyor.northRow,
            "Lasers not correct for North row")
        assert.deepEqual([false,false,false,true,true,false,false,false],
            conveyor.southRow,
            "Lasers not correct for South row")
        done();
    })

    it("second conveyor is built correctly", function(done){
        var conveyor = this.testData[1];
        assert.equal(conveyor.startIndex, 4, "Robot should be at 4");
        assert.deepEqual([false,false,true,false,true,false,true,false],
            conveyor.northRow,
            "Lasers not correct for North row")
        assert.deepEqual([false,false,false,true,true,false,false,false],
            conveyor.southRow,
            "Lasers not correct for South row")
        done();
    })

    it("third conveyor is built correctly", function(done){
        var conveyor = this.testData[2];
        assert.equal(conveyor.startIndex, 4, "Robot should be at 4");
        assert.deepEqual([false,false,true,false,true,false,true,false],
            conveyor.northRow,
            "Lasers not correct for North row")
        assert.deepEqual([false,false,false,true,true,true,false,false],
            conveyor.southRow,
            "Lasers not correct for South row")
        done();
    })
})

