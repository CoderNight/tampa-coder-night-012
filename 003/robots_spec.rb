require 'rubygems'
require 'rspec'

require './robots'

describe "RobotGameSet" do
  
  context "parsing the file" do
    before do
      @game_set = RobotGameSet.new('sample-input.txt')
    end

    it "creates an array of games" do
      @game_set.games.should_not be_nil
      @game_set.size.should eql(3)
    end
  end
end


describe "RobotGame" do
  context "initialize" do
    before do
      lines = ["#|#|#|##", "---X----", "###||###"]
      @robot_game = RobotGame.new(lines)
    end

    it "gets the starting position" do
      @robot_game.starting_position.should eql(3)
    end

    it "gets the risk to the west" do
      @robot_game.west_risk.should eql(2)
    end

    it "gets the risk to the east" do
      @robot_game.east_risk.should eql(3)
    end

    it "spits out the correct direction" do
      @robot_game.safest_direction.should eql('GO WEST')
    end

  end
end