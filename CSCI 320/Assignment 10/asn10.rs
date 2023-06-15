use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::env;
use std::error::Error;


// struct to hold movie information
struct Movie {
    year: String,
    title: String,
    actor: String,
    main_character: String,
}

impl Movie {
    fn new(year: String, title: String, actor: String, main_character: String) -> Self {
        Self {
            year,
            title,
            actor,
            main_character,
        }
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    // from the command line arguments
    let fname = env::args()
        .nth(1)
        .ok_or("./asn10 <filename>")?;

    // Open the file
    let location = Path::new(&fname);
    let file = File::open(&location)?;
    let info = io::BufReader::new(file);

    // Parse into Movie structs
    let mut movies = vec![];
    for line in info.lines() {
        let line = line?;
        let word: Vec<&str> = line.split(',').collect();
        let movie = Movie::new(
            word[0].to_string(),
            word[1].to_string(),
            word[2].to_string(),
            word[3].to_string(),
        );
        movies.push(movie);
    }

    // Loop

    loop {
        //input a year
        println!("Enter year (Ctrl-D to exit):");
        let mut input = String::new();

        if let Err(error) = io::stdin().read_line(&mut input) {
            return Err(Box::new(error));
        }
        // Exit if Ctrl D
        if input.is_empty() {
            
            break;
        }

        let year = input.trim();

        // Match the given year
        let mut answer = vec![];

        for movie in &movies {
            if movie.year == year {
                answer.push(movie);
            }
        }

        if answer.is_empty() {
            println!("Movie not found from year {}", year);
        } else {
            println!("Movie information for the year {}: ", year);
            for movie in answer {
                println!(
                    "{} ({}), starring {} as {}",
                    movie.title, movie.year, movie.actor, movie.main_character
                );
            }
        }
    }

    Ok(())
}