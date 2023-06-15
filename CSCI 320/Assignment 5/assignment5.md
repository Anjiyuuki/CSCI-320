# Angie Bui
# Assignment 5- Syntax


# 2a
[a-z][a-z0-9]*

# 2b
\d*\.?\d*

# 3a
cat myfile && cat yourfile
    - => commandline => list
    - => conditional
    - => conditional "&&" pipeline
    - => pipeline "&&" pipeline
    - => command myfile && pipeline
    - => cat myfile && pipline
    - => cat myfile && command
    - => cat myfile && command yourfile
    - => cat myfile && cat yourfile

# 3b 
ls | grep something > somefile
    - => commandline => list
    - => conditional => pipeline
    - => pipeline "|" command
    - => command ls | pipeline
    - => command ls | grep something pipeline
    - => command => redirection
    - => redirectionop => ">"
    - => command ls | grep something > command word
    - => ls | grep something > somefile


    
# 3c
somecmd someparm || othercmd; gcc myfile > output
    - => commandline -> list ";"
    - => list ";" conditional 
    - => somecmd someparm "||" othercmd ";" gcc myfile ">" output 
    - => command ">" filename 
    - => word myfile ">" filename 
    - => redirection ">" filename 
    - => ">" filename myfile 
    - => redirectionop ">" filename filename 
    - => redirection

# 3d 
cmd1; cmd2 && cmd3; less output
    - =>commandline -> list ";" 
    - => list ";" conditional -> 
    - => list ";" command 
    - => list ";" word ";" cmd2 "&&" cmd3 ";" less output 
    - => conditional ";" less output 
    - => pipeline ";" less output
    - => command ";" less output 
    - => word ";" less output  
    - => less output ";" less output 
    - => less output "&&" pipeline ";" less output 
    - => less output "&&" command ";" less output 
    - => less output "&&" word ";" less output 
    - => less output "||" pipeline ";" less output
    - => less output "||" command ";" less output 
    - => less output "||" word ";" less output 
    - => less output redirection ";" less output
    - => redirection ">" filename ";" less output
    - => redirectionop ">" filename filename ";" less output 
    - => ">" filename output ";" less output  
    - => redirection ">" filename ";" command 
    - => redirection