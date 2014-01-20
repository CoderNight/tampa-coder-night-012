var assert = require('assert'),
    robot = require('../src/robot.js');

describe("robot output test", function(){
    it("Test 1 conveyor", function(){
        var conveyors = [{
            startIndex: 3,
            northRow: [false, true, false ,true, false, true, false, false],
            southRow: [false, false, false, true, true, false, false, false]
        }]

        assert.deepEqual(["GO WEST"], robot(conveyors))
    })

    it("Test 2 conveyors", function(){
        var conveyors = [{
                startIndex: 4,
                northRow: [false, false, true, false, true, false, true, false],
                southRow: [false, false, false, true, true, false, false, false]
            },
            {
                startIndex: 4,
                northRow: [false, false, true, false, true, false, true, false],
                southRow: [false, false, false, true, true, true, false, false]
            }
        ]
        assert.deepEqual(["GO EAST","GO WEST"], robot(conveyors))
    })
})