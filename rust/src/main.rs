use std::fs;

fn main() {
    let file = read_file();
    let lines = filter_lines(&file);
    let blocks = split_into_blocks(lines);
    println!("{}", format_block(&blocks[0]));
    println!("{}", format_block(&blocks[3]));
}

fn format_block(block: &String) -> String {
    let mut words = {
        block
            .split_whitespace()
            .map(|word| word.to_string())
            .collect::<Vec<String>>()
    };
    if !words[0].contains(':') { 
        words[0] = format!("{} {}", words[0], words[1]);
        words.remove(1);
    }
    let mut pairs = Vec::<String>::new();
    let mut pair = String::from("");
    for word in words {
        if word.contains(':') {
            pairs.push(pair);
            pair = String::from("");
            pair += &word;
        }
        else {
            pair += format!(" {}", &word).as_str();
        }
    }
    pairs.join("\n")
}

fn split_into_blocks(file: String) -> Vec<String> {
    let lines = file.split('\n').collect::<Vec<&str>>();
    let mut block = String::from("");
    let mut blocks = Vec::<String>::new();
    for line in lines {
        if line.chars().nth(0) != Some(' ') {
            blocks.push(block);
            block = String::from("");
            block += line;
        }
        else {
            block += line;
        }
    }
    blocks
        .into_iter()
        .filter(|block| block.len() > 10)
        .collect::<Vec<String>>()
}

fn filter_lines(file: &String) -> String {
    file
        .split('\n')
        .into_iter()
        .filter(|&line| &line.chars().nth(0) != &Some('#'))
        .filter(|&line| &line.len() > &10)
        .collect::<Vec<&str>>()
        .join("\n")
}

fn read_file() -> String {
    let path = "../languages-on-github.yml";
    let contents = fs::read_to_string(path);
    match contents {
        Ok(text) => text,
        Err(_) => String::new(),
    }
}
