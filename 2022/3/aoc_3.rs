fn part_1(input: &str) {
    let mut points: u32 = 0;

    for rucksack in input.lines() {
        let chambers: (&str, &str) = rucksack.split_at(rucksack.len() / 2);

        for c in chambers.0.chars() {
            if !chambers.1.contains(c) {
                continue;
            }

            let bytes = c.to_string().into_bytes();
            if bytes[0] < 97 {
                let down_bytes = c.to_ascii_lowercase().to_string().into_bytes();
                points += down_bytes[0] as u32 - 70
            } else {
                points += bytes[0] as u32 - 96
            }

            break;
        }
    }
    println!("{}", points)
}

fn part_2(input: &str) {
    let mut group_points: u32 = 0;

    let groups: Vec<Vec<&str>> = input
        .split('\n')
        .into_iter()
        .collect::<Vec<&str>>()
        .chunks(3)
        .map(|chunk| chunk.to_vec())
        .map(|mut v| {
            v.sort_by_key(|s| s.len());
            v
        })
        .collect();

    for group in groups {
        let chambers: (&str, &str, &str) = (group[0], group[1], group[2]);

        for c in chambers.0.chars() {
            if !chambers.1.contains(c) || !chambers.2.contains(c) {
                continue;
            }

            let bytes = c.to_string().into_bytes();
            if bytes[0] < 97 {
                let down_bytes = c.to_ascii_lowercase().to_string().into_bytes();
                group_points += down_bytes[0] as u32 - 70
            } else {
                group_points += bytes[0] as u32 - 96
            }

            break;
        }
    }
    println!("{}", group_points);
}

fn main() {
    let input = include_str!("input3.txt");

    part_1(input); // 7848
    part_2(input); // 2616
}
