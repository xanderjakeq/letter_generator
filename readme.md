# letter generator

spend less time editing letter templates in ms word.

## requirements

- user input for required information
- user can create/save templates
- generate pdf files from user input with some naming convention

- easy to use for non technical user
    - i want this be an standalone executable bypass complexity of
      deploying/managing a website
    - explore later

## usage

### templating
```
---
bg: /bg_path/bg.png
---

[name]
[first_names]
[company]
[street_address]
[city_address]
[donation_amount]
[donation_date]

[./image_path.jpg]
[/another_image_path.jpg]

*bold text*

`small text`
```
all between `---` are for setup

all valid variables like `[name]` will be replaced with value from the entry
file. others will be treated as a path to an image. everything else will be kept

paragraphs should be written in one line

### cli
write the input file and save it as `whatever.txt`

format:
```
template_file_name
name
company_name (`-` if none)
street_address
city_address
donation_amount donation_date

...
```
each entry separated by a an empty line

generate pdfs
```
./letter_generator ./whatever.txt
```

output pdfs are placed in a directory named `output_today's_date/`

## dev

[maroto](https://github.com/johnfercher/maroto)
- generate pdfs

[htmx](https://htmx.org/)
- ui

