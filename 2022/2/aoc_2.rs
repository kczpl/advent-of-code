use std::fs;

const INPUT: &str = "src/input2.txt";

const POINTS: [(char, i32); 3] = [('X', 1), ('Y', 2), ('Z', 3)];
const POINTS_2: [(char, i32); 3] = [('X', 0), ('Y', 3), ('Z', 6)];

const WIN: [(char, char); 3] = [('A', 'Y'), ('B', 'Z'), ('C', 'X')];
const DRAW: [(char, char); 3] = [('A', 'X'), ('B', 'Y'), ('C', 'Z')];
const LOSE: [(char, char); 3] = [('A', 'Z'), ('B', 'X'), ('C', 'Y')];

fn part_1(games: Vec<(char, char)>) {
    let mut score: i32 = 0;

    for game in games {
        for var in POINTS.iter() {
            if var.0 == game.1 {
                score += var.1
            }
        }

        for var in WIN.iter() {
            if var.0 == game.0 && var.1 == game.1 {
                score += 6;
            }
        }

        for var in DRAW.iter() {
            if var.0 == game.0 && var.1 == game.1 {
                score += 3;
            }
        }
    }

    println!("{}", score)
}

fn part_2(games: Vec<(char, char)>) {
    let mut score: i32 = 0;

    for game in games {
        for var in POINTS_2.iter() {
            if var.0 == game.1 {
                score += var.1
            }
        }

        match POINTS_2.iter().find(|&&x| x.0 == game.1) {
            Some((_, 3)) => {
                let choise = DRAW.iter().find(|&&x| x.0 == game.0).unwrap().1;

                score += POINTS.iter().find(|&&x| x.0 == choise).unwrap().1;
            }
            Some((_, 6)) => {
                let choise = WIN.iter().find(|&&x| x.0 == game.0).unwrap().1;

                score += POINTS.iter().find(|&&x| x.0 == choise).unwrap().1;
            }
            Some((_, 0)) => {
                let choise = LOSE.iter().find(|&&x| x.0 == game.0).unwrap().1;
                score += POINTS.iter().find(|&&x| x.0 == choise).unwrap().1;
            }
            _ => (),
        };
    }

    println!("{}", score)
}

fn main() {
    let games: Vec<(char, char)> = fs::read_to_string(INPUT.to_string())
        .unwrap()
        .lines()
        .map(|line| {
            let mut parts = line.split_whitespace();
            (
                parts.next().unwrap().chars().next().unwrap(),
                parts.next_back().unwrap().chars().next_back().unwrap(),
            )
        })
        .collect();

    part_1(games.clone()); // 12794
    part_2(games.clone()); // 14979
}
