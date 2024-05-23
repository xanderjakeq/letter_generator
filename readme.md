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
create `.txt` files in the `templates/` folder
`general_ty.txt` for example
```
---
bg: /bg_path/bg.png
---

[name] // can have multiple names separated by `&` like `first last & first last & first last`
[first_names]
[company]
[street_address]
[city_address]
[donation_amount]
[donation_date]
[year]
[donation_total]
[donation-1] //used when there are more than 1 donation in the input block and will result to "$[donation_amount] on [donation_date]"
[donation-2] //refers to the second donation and so on

[./image_path.jpg|10]
[/another_image_path.jpg|30]

*bold text*

`small text`
```
all between `---` are for setup

all valid variables like `[name]` will be replaced with value from the entry
file. others will be treated as a path to an image. the `|30` for the image path
is the custom height (10 by default). everything else will be kept

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
donation_amount donation_date //there can be more than 1 donation


template_file_name
firstname|nickname lastname & firstname|nickname lastname
company_name (`-` if none)
street_address
city_address
donation_amount donation_date

...
```
each entry separated by a an empty line

to add a nickname, to use instead of the firstname, do `firstname|nickname
lastname` on the name line

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

