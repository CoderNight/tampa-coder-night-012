require 'rubygems'


class LaserBank

  attr_accessor :lasers

  def initialize(lasers)
    @lasers = lasers
  end

  def threat_at?(location)
    @lasers[location] == '|'
  end


end

class RobotGame
  attr_accessor :starting_position, :lasers, :conveyor_belt, :west_risk, :east_risk

  def initialize(game_lines)
    parse_lines(game_lines)
    get_starting_position
  end

  def safest_direction
    if get_west_risk > get_east_risk
      'GO EAST'
    else
      'GO WEST'
    end
  end

  def risk_at?(location, move_count)
    if move_count.even? 
      @north_lasers.threat_at?(location)
    else
      @south_lasers.threat_at?(location)
    end
  end

  def get_west_risk
    risk = 0
    move_count = 0
    @starting_position.downto(0) do |i|
      risk += 1 if risk_at?(i, move_count)
      move_count += 1
    end
    risk
  end

  def get_east_risk
    risk = 0
    move_count = 0
    @starting_position.upto(@conveyor_belt.size) do |i|
      risk += 1 if risk_at?(i, move_count)
      move_count += 1
    end
    risk
  end

  def get_starting_position
    @starting_position = @conveyor_belt.index('X')
  end

  def parse_lines(game_lines)
    @north_lasers = LaserBank.new(game_lines[0])
    @south_lasers = LaserBank.new(game_lines[2])
    @conveyor_belt = game_lines[1]
  end
end


class RobotGameSet

  attr_accessor :games

  def initialize(filename)
    @games = []
    parse_file(filename)
    @games.map{|g| puts g.safest_direction}
  end

  def size
    games.size
  end

  def parse_file(filename)
    i = 0
    game_lines = []
    File.readlines(filename).each do |line|
      if i < 3
        game_lines << line.strip
        i += 1
      else
        @games << RobotGame.new(game_lines)
        i = 0
        game_lines = []
      end
    end
    @games << RobotGame.new(game_lines)
  end

 
end

RobotGameSet.new(ARGV[0])
