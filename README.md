# Next 5 exercise

## Scenario
Build a "Next 5" Web UI that consumes an API.

## Overview
To make it easy for our customers to place bets on upcoming races that are close to starting, we
display a module that lists the Next 5 upcoming races, sorted by time.
Please create, in a public repository, a functional website that loads a list of Next 5 races.
You are free to use any frameworks (or none) you wish to complete this task.

### UI Requirements
1. The index page will always contain a list of the next 5 races that are open for betting
2. The races will be sorted by their closing time, ascendingly
3. The races will be a mix of the 3 types
a. Thoroughbred
b. Greyhounds
c. Harness
4. Each race will countdown to when betting is suspended
5. When betting is suspended on a race, it will disappear from the list.
6. Clicking on a race will take you to the race page.
7. The race page will contain a list of all the competitors in the race, as well as their position.

### Assumptions
1. A Meeting is a location at which Races are run
2. A Meeting has at least 1 Race
3. Meetings have a type; Thoroughbred, Greyhound or Harness
4. Each Race has at least 4 Competitors
5. A Race has a close time, after which betting is not allowed
6. A Competitor has a Position Number, which is their location at the start line

