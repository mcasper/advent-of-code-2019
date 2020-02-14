# frozen_string_literal: true

abort('Need filename') if ARGV[0].nil?

class Ingrediant
  attr_reader :quantity
  attr_reader :name

  def initialize(quantity:, name:)
    @quantity = quantity
    @name = name
  end

  def self.from_string(string)
    quantity, name = string.split(" ")
    new(quantity: quantity.to_i, name: name)
  end

  def to_s
    "#{quantity} #{name}"
  end
end

class Reaction
  attr_reader :inputs
  attr_reader :output

  def initialize(inputs:, output:)
    @inputs = inputs
    @output = output
  end

  def self.from_string(line)
    inputs_string, output_string = line.split(' => ')
    inputs = inputs_string.split(",").map { |s| Ingrediant.from_string(s) }
    output = Ingrediant.from_string(output_string)
    new(inputs: inputs, output: output)
  end

  def to_s
    "Reaction: #{inputs.map(&:to_s).join(", ")} => #{output}"
  end
end

input = File.read(ARGV[0])
reactions = input.strip.split("\n").map do |line|
  Reaction.from_string(line)
end

need = {}
fuel_reaction = reactions.detect { |r| r.output.name == "FUEL" }
inputs_to_resolve = fuel_reaction.inputs
leftovers = {}

while inputs_to_resolve.any?
  input = inputs_to_resolve.pop

  needed = input.quantity
  need[input.name] ||= 0
  need[input.name] += needed

  if leftovers[input.name]
    if leftovers[input.name] >= needed
      leftovers[input.name] -= needed
      needed = 0
    else
      needed -= leftovers[input.name]
      leftovers[input.name] = 0
    end
  end

  next if needed == 0

  reaction = reactions.detect { |r| r.output.name == input.name }
  times_to_run = (needed.to_f / reaction.output.quantity).ceil
  to_be_outputted = times_to_run * reaction.output.quantity
  leftovers[input.name] ||= 0
  leftovers[input.name] += (to_be_outputted - needed)
  times_to_run.times do
    reaction.inputs.each do |iinput|
      if iinput.name != "ORE"
        inputs_to_resolve << iinput
      end
    end
  end
end

puts need

ore_needed = 0

need.each do |k, v|
  from_ore = reactions.detect { |r| r.output.name == k && r.inputs.first.name == "ORE" }
  if from_ore
    times_to_run = (v.to_f / from_ore.output.quantity).ceil
    ore_needed += (from_ore.inputs.first.quantity * times_to_run)
  end
end

puts ore_needed
