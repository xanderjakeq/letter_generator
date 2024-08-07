# letter generator
spend less time editing letter templates in ms word.

## requirements

- [x] user input for required information
- [ ] user can create/save templates
- [x] generate pdf files from user input with some naming convention

- easy to use for non technical user
    - distribute as zip with executables?
    - deploy to cloud, then send pdfs for download?

## usage

### templating
create `.txt` files in the `templates/` folder
`general_ty.txt` for example
```
---
bg: /bg_path/bg.png
---

[full_names] 
[first_names]
[company]
[street_address]
[city_address]
[donation_amount]
[donation_date]
[year]
[donation_total]
[donation-1] //used when there are more than 1 donation in the input block and 
[donation-2] //will result to "$[donation_amount] on [donation_date]" 
             //[donation-2] refers to the second donation and so on

[./image_path.jpg|10]
[/another_image_path.jpg|30]


//text styling
*bold text*
`small text`
```
all between `---` are for setup

all valid variables like `[full_names]` will be replaced with value from the entry
file. others will be treated as a path to an image. the `|30` for the image path
is the custom height (10 by default). everything else will be kept

note:
if using the sample templates in the sample folder, you will have to update the 
paths to the images inside the assets folder.

![image_info](assets/image_info.png)
- navigate to the file in the finder and select it.
- right-click or control-click on it and choose get info.
- in the window that opens, highlight and copy(command + c) the where field.
- paste it between `[]`
- under name and extension replace spaces with underscore `_`
- copy the new name and extension, paste it at the end of the path
- should look like `[/Users/.../file_name.png]`

paragraphs should be written in one line

### app/server
pick a release [release](https://github.com/xanderjakeq/letter_generator/releases).
download a the `.zip` file, unzip and move to location of your choice.

go to the bin folder inside `letter_generator` and double click the server binary. 
copy paste this link on your browser: http://localhost:3000

for convenience, add the server executable file to the dock
![add_server_to_dock](assets/add_server_to_dock.gif)

### input format
format:
```
template_file_name
full_names // can have multiple names separated by `&` like `first last & first last & first last`
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
entries separated by an empty line

to add a nickname to use instead of the firstname, do `firstname|nickname
lastname` on the `[full_names]` line. if the firstname or nickname is more than
one word, use underscore like `John_Paul` or `Mr._J`

click generate and the output folder should open.

### cli
write the input file and save it as `whatever.txt`

the same import format applies

generate pdfs
```
./path_to_letter_generator_executable ./whatever.txt
```
output pdfs are placed in a directory named `pdf_output/date/`

## dev

install: go, make, air
```
go mod download #install deps

make air #run air to watch server files 
make tailwind #run tailwind cli and watch for file changes in server dir

make build #build cli and server binaries
make export_server #build and copy binary and static files to `output/` dir
```

[maroto](https://github.com/johnfercher/maroto)
- generate pdfs

[htmx](https://htmx.org/)
- ui

