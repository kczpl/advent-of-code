const INPUT: &str = include_str!("./input1.txt");

fn main() {
    part_1(INPUT);
    part_2(INPUT);
}

fn part_1(input: &str) {
    let mut max = 0;
    let mut buffor = 0;

    for line in input.lines() {
        if line.is_empty() {
            if buffor > max {
                max = buffor;
            }
            buffor = 0;
        } else {
            buffor = buffor + (line.to_string()).parse::<i32>().unwrap();
        }
    }

    println!("{}", max)
}

fn part_2(input: &str) {
    let mut elves = vec![];
    let mut buffor = 0;

    for line in input.lines() {
        if line.is_empty() {
            elves.push(buffor);
            buffor = 0;
        } else {
            buffor = buffor + (line.to_string()).parse::<i32>().unwrap();
        }
    }

    elves.sort();
    let result: i32 = elves.iter().rev().take(3).sum();

    println!("{}", result)
}
