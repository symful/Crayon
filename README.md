# Crayon
An interpreted programming language written in Go for idiots.

# What the Fuck?
There are only tags, values, constant values, commands (statements), variables and scope. Well, for now.

# How to Use
```sh
./cor.sh <File>
```
This bash script will build the engine and run it.

# What the Fuck is a Tag?
It works like a function.

## Kinds of Tag
### 1. MAIN
```
[MAIN]
```
It will execute when the engine ready.
<br />
<br />

### 2. FRAME
```
[FRAME XX]
```
It will execute after XX seconds.
<br />
<br />

### 3. LOOP
```
[LOOP]
```
It will execute every time until you exit the program.

### 4. Custom
```
[@AnyName]
```
Your own named tags! This type of tag can be called by your code.

<br />
<br />
ProTip: If you want to end a tag, just use
```
~ <Value?>
```

# Kinds of Value
## 1. String
```
"This is a string"
```
<br />
<br />

## 2. Number (Float 64)
```
10.10
```
<br />
<br />

## 3. Bool
```
#yes
#no
```
<br />
<br />

## 4. PI
```
#PI
```
<br />
<br />

## 5. Object (Untested)
```
#{
    key1 = <Any>,
    key2 = <Any>
}
```
<br />
<br />

## 6. Array
```
#[<Any>, <Any>]
```
<br />
<br />

## 7. None (Untested)
```
#none
```
<br />
<br />


# Commands
## 1. Exit the program
```
K THX BYE;
```
<br />
<br />

## 2. Console
```
write <Any> to console;
```
<br />
<br />

## 3. Define
```
define variable $[variable];
```
<br />
<br />

## 4. Assign
```
assign value <Any> to $[variable];
```
<br />
<br />

## 5. Delete 
```
delete variable $[variable];
```
<br />
<br />

## 6. Global 
```
make variable $[variable] global;
```
<br />
<br />

## 7. Math Operators
```
add <Number> by <Number>;
sub <Number> by <Number>;
divide <Number> by <Number>;
multiply <Number> by <Number>;
```
<br />
<br />

## 8. Logical Operators (All the results are `<Bool>`)
```
is <Any> equal with <Any>;
is <Any> not equal with <Any>;

not <Bool>;

is <Number> less than <Number>;
is <Number> less than or equal with <Number>;

is <Number> more than <Number>;
is <Number> more than or equal with <Number>;

<Bool> and <Bool>;
<Bool> or <Bool>;
```
<br />
<br />

Protip: You can use commands within commands like
```
command1 @(command2);
```
<br />
<br />

## 9. If Else
```
if <Bool> {
    this is scope
}

if <Bool> {
    this is scope
} else {
    this is scope
}
```
<br />
<br />

## 10. Get & Call Custom Tag
```
call tag @[CustomTagName] with args <Array>;
```
To get the arguments in the custom tag, the custom tag just need to get the variable `$args`.
<br />
<br />

## 11. Get Property
```
get property <Number> from <Array>;
get property <String> from <Object>;
```
<br />
<br />

## 12. Set Property
```
set <Any> as property <Number> from <Array>;
set <Any> as property <String> from <Object>;
```
<br />
<br />

## 13. Get Type Of (Untested)
```
get type of <Any>;
```
<br />
<br />